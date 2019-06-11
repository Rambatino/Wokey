package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubResp struct {
	ID           string `json:"id"`
	Desc         string `json:"desc"`
	DescHtml     string `json:"descHtml"`
	DescMarkdown string `json:"descMarkdown"`
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle"`
	Url          string `json:"url"`
	State        string `json:"state"` // set by users
	Branch       string `json:"branch"`
	Status       string `json:"status"` // set by CI
	Repo         string `json:"repo"`
}

func AllCurrentPullRequests(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	prs, _, err := client.Search.Issues(ctx, "author:"+os.Getenv("GITHUB_USER")+" type:pr state:open", nil) // &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}}
	if err != nil {
		log.Println(err.Error())
		return
	}

	store := []GithubResp{}

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

			store = append(store, GithubResp{
				ID:           strconv.Itoa(int(i.GetID())),
				DescMarkdown: i.GetBody(),
				Title:        i.GetTitle(),
				SubTitle:     repo,
				State:        state,
				Url:          pr.GetHTMLURL(),
				Branch:       pr.GetHead().GetRef(),
				Status:       statuses[0].GetState(),
				Repo:         repo,
			})
		}(issue, &wg)
	}
	wg.Wait()
	json.NewEncoder(w).Encode(store)
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
