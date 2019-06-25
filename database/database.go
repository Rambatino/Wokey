package database

import (
	"fmt"
	"time"
)

type comment struct {
	ID      string
	Comment string
}

type pullRequest struct {
	ID            string `json:"id"`
	Number        int    `json:"number"`
	ApprovalState string `json:"approvalState"`
	CIStatus      string `json:"ciStatus"`
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

	Comments []comment
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
	columnChangeText = func(issue issue, from, to string) string {
		return fmt.Sprintf("'%s' [%s] has been moved from '%s' to '%s'", issue.Title, issue.Key, from, to)
	}

	APPROVAL_CHANGE    = "APPROVAL_CHANGE"
	approvalChangeText = func(pr pullRequest) string {
		return fmt.Sprintf("#%v '%s' [%s] has been approved", pr.Number, pr.Title, pr.Repo)
	}

	CHANGES_REQUESTED    = "CHANGES_REQUESTED"
	changesRequestedText = func(pr pullRequest) string {
		return fmt.Sprintf("#%v '%s' [%s] has had changes requested", pr.Number, pr.Title, pr.Repo)
	}

	CI_FAILED    = "CI_FAILED"
	ciFailedText = func(pr pullRequest) string {
		return fmt.Sprintf("#%v '%s' [%s] has failed", pr.Number, pr.Title, pr.Repo)
	}

	CI_SUCCEEDED    = "CI_SUCCEEDED"
	ciSucceededText = func(pr pullRequest) string {
		return fmt.Sprintf("#%v '%s' [%s] has succeeded", pr.Number, pr.Title, pr.Repo)
	}
)

func FindOrNew(bucketID string) *state {
	return &state{githubQuery: &githubQuery{}, jiraQuery: &jiraQuery{}}
}

func (s *state) Store(bucketID string) bool {
	return false
}

func CheckForStateChange(state state, bucketID string) (newState state, changeCount int) {
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

				// check if any added comments
				if len(newIssue.Comments) > len(initialIssue.Comments) {
					newStateChangeStore = append(newStateChangeStore, stateChange{
						CreatedAt: time.Now(),
						Type:      NEW_ISSUE_COMMENT,
						Message:   newIssueCommentText(newIssue.Title, newIssue.Key, len(newIssue.Comments)-len(initialIssue.Comments)),
					})
				}

				// check if state has changed
				if newIssue.State != initialIssue.State {
					newStateChangeStore = append(newStateChangeStore, stateChange{
						CreatedAt: time.Now(),
						Type:      COLUMN_CHANGE,
						Message:   columnChangeText(newIssue, initialIssue.State, newIssue.State),
					})
				}

			}
		}
	}

	for _, newPullRequest := range prs {
		for _, initialPullRequest := range state.PullRequests {
			if newPullRequest.ID == initialPullRequest.ID {

				// check if comment number changed
				if len(newPullRequest.Comments) > len(initialPullRequest.Comments) {
					newStateChange := stateChange{
						CreatedAt: time.Now(),
						Type:      NEW_PULL_REQUEST_COMMENT,
						Message:   newPullRequestCommentText(newPullRequest.Title, newPullRequest.Number, newPullRequest.Repo, len(newPullRequest.Comments)-len(initialPullRequest.Comments)),
					}
					newStateChangeStore = append(newStateChangeStore, newStateChange)
				}

				// check if state has changed
				if newPullRequest.ApprovalState != initialPullRequest.ApprovalState {
					if newPullRequest.ApprovalState == APPROVED_STATE {
						newStateChangeStore = append(newStateChangeStore, stateChange{
							CreatedAt: time.Now(),
							Type:      APPROVAL_CHANGE,
							Message:   approvalChangeText(newPullRequest),
						})
					}
					if newPullRequest.ApprovalState == CHANGES_REQUESTED_STATE {
						newStateChangeStore = append(newStateChangeStore, stateChange{
							CreatedAt: time.Now(),
							Type:      CHANGES_REQUESTED,
							Message:   changesRequestedText(newPullRequest),
						})
					}
				}
				if newPullRequest.CIStatus != initialPullRequest.CIStatus {
					if newPullRequest.CIStatus == CI_FAILED_STATE {
						newStateChangeStore = append(newStateChangeStore, stateChange{
							CreatedAt: time.Now(),
							Type:      CI_FAILED,
							Message:   ciFailedText(newPullRequest),
						})
					}
					if newPullRequest.CIStatus == CI_SUCCEEDED_STATE {
						newStateChangeStore = append(newStateChangeStore, stateChange{
							CreatedAt: time.Now(),
							Type:      CI_SUCCEEDED,
							Message:   ciSucceededText(newPullRequest),
						})
					}
				}
			}
		}
	}

	state.Change = append(state.Change, newStateChangeStore...)
	state.Issues = issues
	state.PullRequests = prs

	// store state
	return state, len(newStateChangeStore)
}
