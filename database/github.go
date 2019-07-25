package database

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubQuerier interface {
	getGithubPullRequests([]string) []pullRequest
}

type githubQuery struct {
}

func NewGithubQuerier() GithubQuerier {
	return &githubQuery{}
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
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// search open,
	// search by branch name matching partial (doesn't have to be just me as author),
	// search by title match
	// head:feature/ADPULSE
	var wg sync.WaitGroup

	prs := make(chan []github.Issue)
	uniqPrs := []github.Issue{}
	wg.Add(1)
	go queryPrs(client, ctx, "author:"+os.Getenv("GITHUB_USER")+" type:pr state:open", prs, &wg)
	for _, key := range jiraKeys {
		wg.Add(1)
		go queryPrs(client, ctx, key+" in:title,body,comments,branch type:pr", prs, &wg)
		wg.Add(1)
		go queryPrs(client, ctx, "type:pr head:feature/"+key+" head:hotfix/"+key, prs, &wg)
	}

	wg.Wait()
	idMap := map[int64]bool{}
	for _, p := range <-prs {
		if !idMap[p.GetID()] {
			uniqPrs = append(uniqPrs, p)
			idMap[p.GetID()] = true
		}
	}

	// only explore open and PRs and PRs assigned to issues in more detail
	// get the comments
	// get the statuses

	store := []pullRequest{}

	for _, issue := range uniqPrs {
		wg.Add(1)
		go func(i github.Issue, w *sync.WaitGroup) {
			defer w.Done()
			repo := strings.Replace(i.GetRepositoryURL(), "https://api.github.com/repos/", "", -1)
			split := strings.Split(repo, "/")

			reviewsChan := make(chan []*github.PullRequestReview)
			prChan := make(chan *github.PullRequest)

			go func() {
				rvs, _, _ := client.PullRequests.ListReviews(ctx, split[0], split[1], i.GetNumber(), nil)
				reviewsChan <- rvs
			}()
			go func() {
				pr, _, _ := client.PullRequests.Get(ctx, split[0], split[1], i.GetNumber())
				prChan <- pr
			}()
			pr := <-prChan

			state := ""

			if pr.GetMergeableState() == "dirty" || pr.GetMergeableState() == "behind" {
				state = formatStatus(CONFLICTED_STATE, state)
			}

			statuses, _, err := client.Repositories.ListStatuses(ctx, split[0], split[1], pr.GetHead().GetSHA(), nil)

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

			store = append(store, pullRequest{
				ID:            strconv.Itoa(int(i.GetID())),
				Title:         i.GetTitle(),
				ApprovalState: state,
				Link:          pr.GetHTMLURL(),
				Branch:        pr.GetHead().GetRef(),
				CIStatus:      getCIStatus(statuses),
				Repo:          repo,
				Number:        i.GetNumber(),
				state:         pr.GetState(),
				author:        pr.GetUser().GetLogin(),
				Comments:      comment{LastCommentLink: lastCommentURL, Count: commentCount},
			})
		}(issue, &wg)
	}
	wg.Wait()
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

func queryPrs(client *github.Client, ctx context.Context, queryString string, ch chan<- []github.Issue, wg *sync.WaitGroup) {
	defer wg.Done()
	prs, _, err := client.Search.Issues(ctx, queryString, &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}})
	if err != nil {
		log.Println(err.Error())
		close(ch)
	}
	ch <- prs.Issues
}
