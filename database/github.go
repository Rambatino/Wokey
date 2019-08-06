package database

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
	"golang.org/x/oauth2"
)

type GithubQuerier interface {
	getGithubPullRequests([]string) []pullRequest
}

type githubQuery struct {
	client *github.Client
	ctx    context.Context
}

func NewGithubQuerier() GithubQuerier {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &githubQuery{github.NewClient(tc), ctx}
}

var (
	CHANGES_REQUESTED_STATE = "CHANGES_REQUESTED"
	APPROVED_STATE          = "APPROVED"
	COMMENTED_STATE         = "COMMENTED"
	CONFLICTED_STATE        = "CONFLICTED"
	IS_MERGED               = "IS_MERGED"

	CI_FAILED_STATE    = "FAILED"
	CI_SUCCEEDED_STATE = "SUCCEEDED"
	CI_RUNNING_STATE   = "RUNNING"
)

func (g *githubQuery) getGithubPullRequests(jiraKeys []string) []pullRequest {
	// search open,
	// search by branch name matching partial (doesn't have to be just me as author),
	// search by title match
	// head:feature/ADPULSE
	var wg sync.WaitGroup
	prs := make(chan []github.Issue)
	chanErr := make(chan error)
	pullRequestChan := make(chan pullRequest)

	wg.Add(1 + len(jiraKeys)*2)
	go g.queryPrs("author:"+os.Getenv("GITHUB_USER")+" type:pr state:open", prs, chanErr)
	for _, key := range jiraKeys {
		go g.queryPrs(key+" in:title,body,comments,branch type:pr", prs, chanErr)
		go g.queryPrs("type:pr head:feature/"+key+" head:hotfix/"+key, prs, chanErr)
	}

	go func(wg *sync.WaitGroup) {
		idMap := map[int64]bool{}
		for {
			select {
			case newPrs := <-prs:
				for _, p := range newPrs {
					if !idMap[p.GetID()] {
						idMap[p.GetID()] = true
						wg.Add(1)
						go g.fillExtraInformation(p, pullRequestChan)
					}
				}
				wg.Done()
			case err := <-chanErr:
				pp.Println("An error occured fetching PRs: " + err.Error())
				wg.Done()
			}
		}
	}(&wg)

	store := []pullRequest{}

	go func(wg *sync.WaitGroup) {
		for {
			select {
			case filledPr := <-pullRequestChan:
				store = append(store, filledPr)
				wg.Done()
			}
		}
	}(&wg)

	wg.Wait()
	// only explore open and PRs and PRs assigned to issues in more detail
	// get the comments
	// get the statuses

	return store
}

func formatStatus(newStatus, previousStatus string) string {
	if newStatus == "" && previousStatus == "" {
		return ""
	}

	weightMap := map[string]int{
		COMMENTED_STATE:         1,
		CHANGES_REQUESTED_STATE: 2,
		APPROVED_STATE:          3,
		CONFLICTED_STATE:        4,
		IS_MERGED:               5,
	}

	if previousStatus == "" || newStatus == previousStatus {
		return newStatus
	}

	if weightMap[newStatus] > weightMap[previousStatus] {
		return newStatus
	}

	return previousStatus
}

func getCIStatus(statuses []*github.RepoStatus) ciStatus {
	s := ciStatus{}
	if len(statuses) == 0 {
		return s
	}
	contextStore := map[string]bool{}
	weightMap := map[string]int{
		"":        -1,
		"success": 0,
		"pending": 1,
		"failure": 2,
	}
	for _, status := range statuses {
		if !contextStore[status.GetContext()] {
			if weightMap[status.GetState()] > weightMap[s.Status] {
				s = ciStatus{status.GetState(), status.GetTargetURL()}
			}
			contextStore[status.GetContext()] = true
		}
	}
	return s
}

func (g *githubQuery) queryPrs(queryString string, ch chan<- []github.Issue, chanErr chan<- error) {
	prs, _, err := g.client.Search.Issues(g.ctx, queryString, &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}})
	if err != nil {
		chanErr <- err
		return
	}
	ch <- prs.Issues
}

func (g *githubQuery) fillExtraInformation(issue github.Issue, res chan<- pullRequest) {
	repo := strings.Replace(issue.GetRepositoryURL(), "https://api.github.com/repos/", "", -1)
	split := strings.Split(repo, "/")

	reviewsChan := make(chan []*github.PullRequestReview)
	prChan := make(chan *github.PullRequest)

	go func() {
		rvs, _, _ := g.client.PullRequests.ListReviews(g.ctx, split[0], split[1], issue.GetNumber(), nil)
		reviewsChan <- rvs
	}()
	go func() {
		pr, _, _ := g.client.PullRequests.Get(g.ctx, split[0], split[1], issue.GetNumber())
		prChan <- pr
	}()
	pr := <-prChan

	state := ""

	if pr.GetMergeableState() == "dirty" || pr.GetMergeableState() == "behind" {
		state = formatStatus(CONFLICTED_STATE, state)
	}

	statuses, _, err := g.client.Repositories.ListStatuses(g.ctx, split[0], split[1], pr.GetHead().GetSHA(), nil)

	if err != nil {
		log.Println(err.Error())
		return
	}

	lastCommentURL := ""
	commentCount := 0
	for _, rev := range <-reviewsChan {
		if rev.GetHTMLURL() != "" {
			lastCommentURL = rev.GetHTMLURL()
		}
		commentCount++
		state = formatStatus(rev.GetState(), state)
	}

	if pr.GetMerged() {
		state = IS_MERGED
	}

	res <- pullRequest{
		ID:            strconv.Itoa(int(issue.GetID())),
		Title:         issue.GetTitle(),
		ApprovalState: state,
		Link:          pr.GetHTMLURL(),
		Branch:        pr.GetHead().GetRef(),
		CIStatus:      getCIStatus(statuses),
		Repo:          repo,
		Number:        issue.GetNumber(),
		state:         pr.GetState(),
		author:        pr.GetUser().GetLogin(),
		Comments:      comment{LastCommentLink: lastCommentURL, Count: commentCount},
	}
}
