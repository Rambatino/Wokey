package database

import (
	"os"

	"github.com/andygrunwald/go-jira"
)

type JiraQuerier interface {
	getJiraIssues() []issue
}

type jiraQuery struct {
}

func NewJiraQuery() JiraQuerier {
	return &jiraQuery{}
}

func (j *jiraQuery) getJiraIssues() []issue {
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_EMAIL"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	baseURL := "https://zigroup.atlassian.net"

	client, err := jira.NewClient(tp.Client(), baseURL)

	if err != nil {
		return nil
	}
	u, _, err := client.Issue.Search("assignee = currentUser() AND resolution = Unresolved order by updated DESC", &jira.SearchOptions{Expand: "renderedFields"})

	if err != nil {
		return nil
	}

	store := []issue{}

	for _, jiraIssue := range u {
		store = append(store, issue{
			ID:    jiraIssue.ID,
			Title: jiraIssue.Fields.Summary,
			Key:   jiraIssue.Key,
			State: jiraIssue.Fields.Status.Name,
			Link:  baseURL + "/browse/" + jiraIssue.Key,
		})
	}
	return store
}
