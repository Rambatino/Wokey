package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
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
	State        string `json:"state"`
	Branch       string `json:"branch"`
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
		pp.Println(err.Error())
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

			reviews, _, err := client.PullRequests.ListReviews(ctx, split[0], split[1], i.GetNumber(), nil)
			pr, _, err := client.PullRequests.Get(ctx, split[0], split[1], i.GetNumber())

			if err != nil {
				pp.Println(err.Error())
				return
			}
			state := ""
			for _, rev := range reviews {
				state = rev.GetState()
			}

			store = append(store, GithubResp{
				ID:           strconv.Itoa(int(i.GetID())),
				DescMarkdown: i.GetBody(),
				Title:        i.GetTitle(),
				SubTitle:     repo,
				State:        state,
				Url:          i.GetURL(),
				Branch:       pr.GetHead().GetRef(),
			})
		}(issue, &wg)
	}
	wg.Wait()
	json.NewEncoder(w).Encode(store)
}
