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
	getGithubPullRequests() []pullRequest
}

type githubQuery struct {
}

func NewGithubQuerier() GithubQuerier {
	return &githubQuery{}
}

func (g *githubQuery) getGithubPullRequests() []pullRequest {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	prs, _, err := client.Search.Issues(ctx, "author:"+os.Getenv("GITHUB_USER")+" type:pr state:open", nil) // &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}}
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	store := []pullRequest{}

	var wg sync.WaitGroup
	for _, issue := range prs.Issues {
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
			state := ""
			go func() {
				pr, _, _ := client.PullRequests.Get(ctx, split[0], split[1], i.GetNumber())
				if pr.GetMergeableState() == "dirty" {
					state = formatStatus("CONFLICTED", state)
				}

				prChan <- pr
			}()
			pr := <-prChan
			statuses, _, err := client.Repositories.ListStatuses(ctx, split[0], split[1], pr.GetHead().GetSHA(), nil)

			reviews := <-reviewsChan

			if err != nil {
				log.Println(err.Error())
				return
			}

			for _, rev := range reviews {
				state = formatStatus(rev.GetState(), state)
			}

			store = append(store, pullRequest{
				ID:            strconv.Itoa(int(i.GetID())),
				Title:         i.GetTitle(),
				ApprovalState: state,
				Link:          pr.GetHTMLURL(),
				Branch:        pr.GetHead().GetRef(),
				CIStatus:      statuses[0].GetState(),
				Repo:          repo,
				Number:        i.GetNumber(),
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

	statusMap := map[string]string{
		"CONFLICTED":      "conflicted",
		"COMMENTED":       "commented",
		"REQUEST_CHANGES": "requestChanges",
		"APPROVED":        "approved",
	}

	weightMap := map[string]int{
		"conflicted":     0,
		"commented":      1,
		"requestChanges": 2,
		"approved":       3,
	}

	mappedStatus := statusMap[newStatus]
	if previousStatus == "" || mappedStatus == previousStatus {
		return mappedStatus
	}

	if weightMap[mappedStatus] > weightMap[previousStatus] {
		return mappedStatus
	}

	return previousStatus
}
