package database

import (
	reflect "reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	gomock "github.com/golang/mock/gomock"
)

// if approved but conflicted then unconflicted, does it go back to approve? Could it go
// back to just commented?

func TestCheckForState(t *testing.T) {
	now := time.Now()
	patch := monkey.Patch(time.Now, func() time.Time { return now })
	defer patch.Unpatch()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		name                string
		initialPullRequests []pullRequest
		initialIssues       []issue
		newPullRequests     []pullRequest
		newIssues           []issue
		expected            []stateChange
		hasChanged          int
	}{
		{"New Issue Comment", []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1"}}, []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", Comments: []comment{{Comment: "hello"}}}}, []stateChange{{Type: "NEW_ISSUE_COMMENT", Message: "'Title 1' [EPIC/EPIC-23] has 1 new comment(s)", CreatedAt: now}}, 1},
		{"New PR Comment", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Comments: []comment{{Comment: "new comment"}}, Number: 299}}, []issue{}, []stateChange{{Type: "NEW_PULL_REQUEST_COMMENT", Message: "#299 'My PR' [Rambatino/CHAID] has 1 new comment(s)", CreatedAt: now}}, 1},
		{"New Column Change", []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", State: "Ready to Merge"}}, []pullRequest{}, []issue{{State: "Code Review", ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", Comments: []comment{{Comment: "hello"}}}}, []stateChange{{Type: "NEW_ISSUE_COMMENT", Message: "'Title 1' [EPIC/EPIC-23] has 1 new comment(s)", CreatedAt: now}, {Type: "COLUMN_CHANGE", Message: "'Title 1' [EPIC/EPIC-23] has been moved from 'Ready to Merge' to 'Code Review'", CreatedAt: now}}, 2},
		{"New PR Approval", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: "CHANGES_REQUESTED"}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: "APPROVED"}}, []issue{}, []stateChange{{Type: "APPROVAL_CHANGE", Message: "#299 'My PR' [Rambatino/CHAID] has been approved", CreatedAt: now}}, 1},
		{"New Changes Requested", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: ""}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: "CHANGES_REQUESTED"}}, []issue{}, []stateChange{{Type: "CHANGES_REQUESTED", Message: "#299 'My PR' [Rambatino/CHAID] has had changes requested", CreatedAt: now}}, 1},
		{"CI failed for PR", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: "RUNNING"}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: "FAILED"}}, []issue{}, []stateChange{{Type: "CI_FAILED", Message: "#299 'My PR' [Rambatino/CHAID] has failed", CreatedAt: now}}, 1},
		{"CI Succeeded for PR", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: "FAILED"}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: "SUCCEEDED"}}, []issue{}, []stateChange{{Type: "CI_SUCCEEDED", Message: "#299 'My PR' [Rambatino/CHAID] has succeeded", CreatedAt: now}}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockJira := NewMockJiraQuerier(ctrl)
			mockJira.EXPECT().getJiraIssues().Return(tc.newIssues)
			mockGithub := NewMockGithubQuerier(ctrl)
			mockGithub.EXPECT().getGithubPullRequests().Return(tc.newPullRequests)
			initialState := state{
				PullRequests: tc.initialPullRequests,
				Issues:       tc.initialIssues,
				jiraQuery:    mockJira,
				githubQuery:  mockGithub,
				Change:       []stateChange{},
			}
			newState, hasChanged := CheckForStateChange(initialState, "")

			if !reflect.DeepEqual(newState.Change, tc.expected) {
				t.Error("Change is not equal to expected", newState.Change, tc.expected)
			}

			if hasChanged != tc.hasChanged {
				t.Error("Has changed is not equal to expected", hasChanged, tc.hasChanged)
			}
		})
	}
}
