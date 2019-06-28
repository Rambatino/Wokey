package database

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type comment struct {
	Count           int    `json:"count"`
	LastCommentLink string `json:"lastCommentLink"`
}

type ciStatus struct {
	Status string `json:"status"`
	Link   string `json:"link"`
}

type pullRequest struct {
	ID            string   `json:"id"`
	Number        int      `json:"number"`
	ApprovalState string   `json:"approvalState"`
	CIStatus      ciStatus `json:"ciStatus"`
	Title         string   `json:"title"`
	Link          string   `json:"link"`
	Branch        string   `json:"branch"`
	Repo          string   `json:"repo"`
	state         string
	author        string
	Comments      comment `json:"comments"`
}

func (p *pullRequest) label() string {
	return fmt.Sprintf("#%v '%s' [%s]", p.Number, p.Title, p.Repo)
}

type issue struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Title string `json:"title"`
	Link  string `json:"link"`
	State string `json:"state"`

	PullRequests []pullRequest `json:"pullRequests"`
	Comments     comment       `json:"comments"`
}

func (i *issue) label() string {
	return fmt.Sprintf("'%s' [%s]", i.Title, i.Key)
}

type state struct {
	githubQuery  GithubQuerier
	jiraQuery    JiraQuerier
	Issues       []issue       `json:"issues"`
	PullRequests []pullRequest `json:"pullRequests"`
	Changes      []stateChange `json:"changes"`
}

type stateChange struct {
	Type          string    `json:"type"`
	Read          bool      `json:"read"`
	Message       string    `json:"message"`
	PullRequestID string    `json:"pullRequestID"`
	IssueID       string    `json:"issueID"`
	CreatedAt     time.Time `json:"createdAt"`
	SeenAt        time.Time `json:"seenAt"`
}

var (
	NEW_ISSUE_ASSIGNED   = "NEW_ISSUE_ASSIGNED"
	newIssueAssignedText = func(issue issue) string {
		return fmt.Sprintf("%s has been assigned to you", issue.label())
	}

	NEW_ISSUE_COMMENT   = "NEW_ISSUE_COMMENT"
	newIssueCommentText = func(issue issue, count int) string {
		return fmt.Sprintf("%s has %v new comment(s)", issue.label(), count)
	}

	NEW_PULL_REQUEST_COMMENT  = "NEW_PULL_REQUEST_COMMENT"
	newPullRequestCommentText = func(pr pullRequest, count int) string {
		return fmt.Sprintf("%s has %v new comment(s)", pr.label(), count)
	}

	COLUMN_CHANGE    = "COLUMN_CHANGE"
	columnChangeText = func(issue issue, from, to string) string {
		return fmt.Sprintf("%s has been moved from '%s' to '%s'", issue.label(), from, to)
	}

	APPROVAL_CHANGE    = "APPROVAL_CHANGE"
	approvalChangeText = func(pr pullRequest) string {
		return fmt.Sprintf("%s has been approved", pr.label())
	}

	CHANGES_REQUESTED    = "CHANGES_REQUESTED"
	changesRequestedText = func(pr pullRequest) string {
		return fmt.Sprintf("%s has had changes requested", pr.label())
	}

	CI_FAILED    = "CI_FAILED"
	ciFailedText = func(pr pullRequest) string {
		return fmt.Sprintf("%s has failed", pr.label())
	}

	CI_SUCCEEDED    = "CI_SUCCEEDED"
	ciSucceededText = func(pr pullRequest) string {
		return fmt.Sprintf("%s has succeeded", pr.label())
	}

	NEW_PULL_REQUEST_ASSIGNED_TO_ISSUE = "NEW_PULL_REQUEST_ASSIGNED_TO_ISSUE"
	newPullRequestAssignedToIssueText  = func(pr pullRequest, issue issue) string {
		return fmt.Sprintf("%s has been assigned to %s", pr.label(), issue.label())
	}

	NEW_PULL_REQUEST_OPENED  = "NEW_PULL_REQUEST_OPENED"
	newPullRequestOpenedText = func(pr pullRequest) string {
		return fmt.Sprintf("%s has been opened", pr.label())
	}
)

func CheckForStateChange(state state, bucketID string) (newState state, changeCount int) {
	// get previous state, or return empty state
	// state := state.(*state)
	// get github pull requests and jira stories
	// check if state is different to previous, if new, no state change, otherwise
	// calculate state change

	newStateChangeStore := []stateChange{}

	issues := state.jiraQuery.getJiraIssues()

	keys := map[string]bool{}
	keyStore := []string{}
	for _, issue := range issues {
		if !keys[issue.Key] {
			keyStore = append(keyStore, issue.Key)
			keys[issue.Key] = true
		}
	}
	prs := state.githubQuery.getGithubPullRequests(keyStore)

	for _, newIssue := range issues {
		found := false
		for _, initialIssue := range state.Issues {
			if newIssue.ID == initialIssue.ID {
				found = true
				// check if any added comments
				if newIssue.Comments.Count > initialIssue.Comments.Count {
					newStateChangeStore = append(newStateChangeStore, stateChange{
						CreatedAt: time.Now(),
						Type:      NEW_ISSUE_COMMENT,
						Message:   newIssueCommentText(newIssue, newIssue.Comments.Count-initialIssue.Comments.Count),
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
		if !found {
			newStateChangeStore = append(newStateChangeStore, stateChange{
				CreatedAt: time.Now(),
				Type:      NEW_ISSUE_ASSIGNED,
				Message:   newIssueAssignedText(newIssue),
			})
		}
	}

	for _, newPullRequest := range prs {
		found := false
		for _, initialIssue := range state.Issues {
			for _, initialPullRequest := range initialIssue.PullRequests {
				if newPullRequest.ID == initialPullRequest.ID {
					newStateChangeStore = append(newStateChangeStore, checkPullRequests(initialPullRequest, newPullRequest)...)
					found = true
				}
			}
		}
		for _, initialPullRequest := range state.PullRequests {
			if newPullRequest.ID == initialPullRequest.ID {
				newStateChangeStore = append(newStateChangeStore, checkPullRequests(initialPullRequest, newPullRequest)...)
				found = true
			}
		}
		if !found {
			newStateChangeStore = append(newStateChangeStore, stateChange{
				CreatedAt: time.Now(),
				Type:      NEW_PULL_REQUEST_OPENED,
				Message:   newPullRequestOpenedText(newPullRequest),
			})
		}
	}

	return formatState(issues, prs, append(state.Changes, newStateChangeStore...)), len(newStateChangeStore)
}

// need to combine new comments together if there are mutliple new comments coming in
func formatState(issues []issue, prs []pullRequest, stateChange []stateChange) state {
	lonePullRequests := []pullRequest{}
	issuesCopy := append(issues[:0:0], issues...)

	for _, pr := range prs {
		found := false
		for i, issueVar := range issues {
			if strings.Contains(pr.Branch, issueVar.Key) || strings.Contains(pr.Title, issueVar.Key) {
				var is issue
				is = issuesCopy[i]
				pulls := append(is.PullRequests, pr)
				sort.Slice(pulls, func(i, j int) bool {
					return pulls[i].Number < pulls[j].Number
				})
				is.PullRequests = pulls
				issuesCopy[i] = is
				found = true
			}
		}
		if !found && pr.state == "open" && pr.author == os.Getenv("GITHUB_USER") {
			lonePullRequests = append(lonePullRequests, pr)
		}
	}

	sort.Slice(lonePullRequests, func(i, j int) bool {
		return lonePullRequests[i].Number < lonePullRequests[j].Number
	})

	return state{
		Changes:      stateChange,
		Issues:       issuesCopy,
		PullRequests: lonePullRequests,
	}
}

func checkPullRequests(initialPullRequest, newPullRequest pullRequest) []stateChange {
	newStateChangeStore := []stateChange{}

	// check if comment number changed
	if newPullRequest.Comments.Count > initialPullRequest.Comments.Count {
		newStateChange := stateChange{
			CreatedAt: time.Now(),
			Type:      NEW_PULL_REQUEST_COMMENT,
			Message:   newPullRequestCommentText(newPullRequest, newPullRequest.Comments.Count-initialPullRequest.Comments.Count),
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
		if newPullRequest.CIStatus.Status == CI_FAILED_STATE {
			newStateChangeStore = append(newStateChangeStore, stateChange{
				CreatedAt: time.Now(),
				Type:      CI_FAILED,
				Message:   ciFailedText(newPullRequest),
			})
		}
		if newPullRequest.CIStatus.Status == CI_SUCCEEDED_STATE {
			newStateChangeStore = append(newStateChangeStore, stateChange{
				CreatedAt: time.Now(),
				Type:      CI_SUCCEEDED,
				Message:   ciSucceededText(newPullRequest),
			})
		}
	}
	return newStateChangeStore
}
