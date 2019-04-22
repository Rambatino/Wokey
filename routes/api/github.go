package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
	"golang.org/x/oauth2"
)

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

	store := []Resp{}
	for _, issue := range prs.Issues {
		repo := strings.Replace(issue.GetRepositoryURL(), "https://api.github.com/repos/", "", -1)
		split := strings.Split(repo, "/")
		reviews, _, err := client.PullRequests.ListReviews(ctx, split[0], split[1], issue.GetNumber(), nil)
		if err != nil {
			pp.Println(err.Error())
			return
		}
		state := ""
		for _, rev := range reviews {
			state = rev.GetState()
		}
		store = append(store, Resp{
			ID:       strconv.Itoa(int(issue.GetID())),
			Desc:     issue.GetBody(),
			Title:    issue.GetTitle(),
			SubTitle: repo,
			State:    state,
			Url:      issue.GetURL(),
		})
	}
	json.NewEncoder(w).Encode(store)
}
