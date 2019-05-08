package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Rambatino/go-jira"
	"github.com/k0kubun/pp"
)

type JiraResp struct {
	ID       string `json:"id"`
	Desc     string `json:"desc"`
	DescHtml string `json:"descHtml"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Url      string `json:"url"`
	State    string `json:"state"`
}

func AllIssuesHandler(w http.ResponseWriter, r *http.Request) {
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_EMAIL"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	client, err := jira.NewClient(tp.Client(), "https://zigroup.atlassian.net")

	if err != nil {
		pp.Println(err)
		return
	}
	u, _, err := client.Issue.Search("assignee = currentUser() AND resolution = Unresolved order by updated DESC", &jira.SearchOptions{Expand: "renderedFields"})

	if err != nil {
		pp.Println(err)
		return
	}

	store := []JiraResp{}

	for _, issue := range u {
		if issue.Fields.Status.Name != "Done" {
			store = append(store, JiraResp{
				ID:       issue.ID,
				DescHtml: issue.RenderedFields.Description,
				Desc:     issue.Fields.Description,
				Title:    issue.Fields.Summary,
				SubTitle: issue.Key,
				State:    issue.Fields.Status.Name,
				Url:      "https://zigroup.atlassian.net/browse/" + issue.Key,
			})
		}
	}

	json.NewEncoder(w).Encode(store)
}
