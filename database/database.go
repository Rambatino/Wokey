package database

import (
	"fmt"
	reflect "reflect"
	"time"
)

type comment struct {
	ID      string
	comment string
}

type pullRequest struct {
	ID            string `json:"id"`
	Number        int    `json:"number"`
	ApprovalState string `json:"approvalState"`
	CIStatus      string `json:"cIStatus"`
	Title         string `json:"title"`
	Link          string `json:"link"`
	Branch        string `json:"branch"`
	Repo          string `json:"repo"`

	Comments []comment
}

type issue struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Title string `json:"title"`
	Link  string `json:"link"`
	State string `json:"state"`

	Comments []string
}

type state struct {
	githubQuery  GithubQuerier
	jiraQuery    JiraQuerier
	Issues       []issue
	PullRequests []pullRequest
	Change       []stateChange
}

type stateChange struct {
	Type          string
	Read          bool
	Message       string
	PullRequestID string
	IssueID       string
	CreatedAt     time.Time
	SeenAt        time.Time
}

var (
	// NEW_ISSUE_ASSIGNED   = "NEW_ISSUE_ASSIGNED"
	// newIssueAssignedText = func(issueName, issueId string) string {
	// 	return fmt.Sprintf("%s [%s] has been assigned to you", issueName, issueId)
	// }

	NEW_ISSUE_COMMENT   = "NEW_ISSUE_COMMENT"
	newIssueCommentText = func(issueName, issueKey string, count int) string {
		return fmt.Sprintf("'%s' [%s] has %v new comment(s)", issueName, issueKey, count)
	}

	NEW_PULL_REQUEST_COMMENT  = "NEW_PULL_REQUEST_COMMENT"
	newPullRequestCommentText = func(prName string, prNumber int, repo string, count int) string {
		return fmt.Sprintf("#%v '%s' [%s] has %v new comment(s)", prNumber, prName, repo, count)
	}

	COLUMN_CHANGE    = "COLUMN_CHANGE"
	columnChangeText = func(issueName, issueKey, from, to string) string {
		return fmt.Sprintf("'%s' [%s] has been moved from '%s' to '%s'", issueName, issueKey, from, to)
	}

	APPROVAL_CHANGE    = "APPROVAL_CHANGE"
	approvalChangeText = func(prName string, prNumber int, repo string) string {
		return fmt.Sprintf("#%v '%s' [%s] has been approved", prNumber, prName, repo)
	}

	CHANGES_REQUESTED    = "CHANGES_REQUESTED"
	changesRequestedText = func(prName string, prNumber int, repo string) string {
		return fmt.Sprintf("#%v '%s' [%s] - changes have been requested", prNumber, prName, repo)
	}

	CI_FAILED    = "CI_FAILED"
	ciFailedText = func(prName string, prNumber int, repo string) string {
		return fmt.Sprintf("#%v '%s' [%s] has failed", prNumber, prName, repo)
	}

	CI_SUCCEEDED    = "CI_SUCCEEDED"
	ciSucceededText = func(prName string, prNumber int, repo string) string {
		return fmt.Sprintf("#%v '%s' [%s] has succeeded", prNumber, prName, repo)
	}
)

func FindOrNew(bucketID string) *state {
	return &state{githubQuery: &githubQuery{}, jiraQuery: &jiraQuery{}}
}

func (s *state) Store(bucketID string) bool {
	return false
}

func CheckForStateChange(state state, bucketID string) (newState state, hasChanged bool) {
	// get previous state, or return empty state
	// state := state.(*state)
	// get github pull requests and jira stories
	// check if state is different to previous, if new, no state change, otherwise
	// calculate state change

	var issuesChan = make(chan []issue)
	var prsChan = make(chan []pullRequest)
	newStateChangeStore := []stateChange{}

	go func() {
		issuesChan <- state.jiraQuery.getJiraIssues()
	}()

	go func() {
		prsChan <- state.githubQuery.getGithubPullRequests()
	}()
	issues := <-issuesChan
	prs := <-prsChan
	for _, newIssue := range issues {
		for _, initialIssue := range state.Issues {
			if newIssue.ID == initialIssue.ID {
				if len(newIssue.Comments) > len(initialIssue.Comments) {
					newStateChange := stateChange{
						CreatedAt: time.Now(),
						Type:      NEW_ISSUE_COMMENT,
						Message:   newIssueCommentText(newIssue.Title, newIssue.Key, len(newIssue.Comments)-len(initialIssue.Comments)),
					}
					newStateChangeStore = append(newStateChangeStore, newStateChange)
				}
			}
		}
	}

	if !reflect.DeepEqual(state.PullRequests, prs) {

	}

	state.Change = append(state.Change, newStateChangeStore...)
	state.Issues = issues
	state.PullRequests = prs

	// store state
	return state, len(newStateChangeStore) > 0
}
