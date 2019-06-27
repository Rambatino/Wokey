package database

import (
	"os"
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
		{"New Issue Assigned", []pullRequest{}, []issue{}, []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1"}}, []stateChange{{Type: "NEW_ISSUE_ASSIGNED", Message: "'Title 1' [EPIC/EPIC-23] has been assigned to you", CreatedAt: now}}, 1},
		{"New Issue Comment", []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1"}}, []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", Comments: []comment{{Comment: "hello"}}}}, []stateChange{{Type: "NEW_ISSUE_COMMENT", Message: "'Title 1' [EPIC/EPIC-23] has 1 new comment(s)", CreatedAt: now}}, 1},
		{"New PR Comment", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Comments: []comment{{Comment: "new comment"}}, Number: 299}}, []issue{}, []stateChange{{Type: "NEW_PULL_REQUEST_COMMENT", Message: "#299 'My PR' [Rambatino/CHAID] has 1 new comment(s)", CreatedAt: now}}, 1},
		{"New PR Comment inside an Issue", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299}}, []issue{{PullRequests: []pullRequest{{ID: "1233", Repo: "Rambatino/CHAID", Title: "My Other PR", Number: 300}}, Title: "hello", Key: "EPIC-2"}}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Comments: []comment{{Comment: "new comment"}}, Number: 299}, {ID: "1233", Repo: "Rambatino/CHAID", Title: "My Other PR", Number: 300, Comments: []comment{{Comment: "new comment"}}}}, []issue{{PullRequests: []pullRequest{{ID: "1233", Repo: "Rambatino/CHAID", Title: "My Other PR", Comments: []comment{{Comment: "new comment"}}, Number: 300}}, Title: "hello", Key: "EPIC-2"}}, []stateChange{{Type: "NEW_PULL_REQUEST_COMMENT", Message: "#299 'My PR' [Rambatino/CHAID] has 1 new comment(s)", CreatedAt: now}, {Type: "NEW_PULL_REQUEST_COMMENT", Message: "#300 'My Other PR' [Rambatino/CHAID] has 1 new comment(s)", CreatedAt: now}}, 2},
		{"New Column Change", []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", State: "Ready to Merge"}}, []pullRequest{}, []issue{{State: "Code Review", ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", Comments: []comment{{Comment: "hello"}}}}, []stateChange{{Type: "NEW_ISSUE_COMMENT", Message: "'Title 1' [EPIC/EPIC-23] has 1 new comment(s)", CreatedAt: now}, {Type: "COLUMN_CHANGE", Message: "'Title 1' [EPIC/EPIC-23] has been moved from 'Ready to Merge' to 'Code Review'", CreatedAt: now}}, 2},
		{"New PR Approval", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: "CHANGES_REQUESTED"}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: "APPROVED"}}, []issue{}, []stateChange{{Type: "APPROVAL_CHANGE", Message: "#299 'My PR' [Rambatino/CHAID] has been approved", CreatedAt: now}}, 1},
		{"New Changes Requested", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: ""}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, ApprovalState: "CHANGES_REQUESTED"}}, []issue{}, []stateChange{{Type: "CHANGES_REQUESTED", Message: "#299 'My PR' [Rambatino/CHAID] has had changes requested", CreatedAt: now}}, 1},
		{"CI failed for PR", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: ciStatus{Status: "RUNNING"}}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: ciStatus{Status: "FAILED"}}}, []issue{}, []stateChange{{Type: "CI_FAILED", Message: "#299 'My PR' [Rambatino/CHAID] has failed", CreatedAt: now}}, 1},
		{"CI Succeeded for PR", []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: ciStatus{Status: "FAILED"}}}, []issue{}, []pullRequest{{ID: "123", Repo: "Rambatino/CHAID", Title: "My PR", Number: 299, CIStatus: ciStatus{Status: "SUCCEEDED"}}}, []issue{}, []stateChange{{Type: "CI_SUCCEEDED", Message: "#299 'My PR' [Rambatino/CHAID] has succeeded", CreatedAt: now}}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockJira := NewMockJiraQuerier(ctrl)
			mockJira.EXPECT().getJiraIssues().Return(tc.newIssues)
			mockGithub := NewMockGithubQuerier(ctrl)
			keys := map[string]bool{}
			keyStore := []string{}
			for _, issue := range tc.newIssues {
				if !keys[issue.Key] {
					keyStore = append(keyStore, issue.Key)
					keys[issue.Key] = true
				}
			}
			mockGithub.EXPECT().getGithubPullRequests(keyStore).Return(tc.newPullRequests)
			initialState := state{
				PullRequests: tc.initialPullRequests,
				Issues:       tc.initialIssues,
				jiraQuery:    mockJira,
				githubQuery:  mockGithub,
				Changes:      []stateChange{},
			}
			newState, hasChanged := CheckForStateChange(initialState, "")
			if !reflect.DeepEqual(newState.Changes, tc.expected) {
				t.Errorf("Change is not equal to expected\n%+v\n%+v	", newState.Changes, tc.expected)
			}

			if hasChanged != tc.hasChanged {
				t.Error("Has changed is not equal to expected", hasChanged, tc.hasChanged)
			}
		})
	}
}

func TestFormatState(t *testing.T) {
	os.Setenv("GITHUB_USER", "Rambatino")
	cases := []struct {
		name            string
		newPullRequests []pullRequest
		newIssues       []issue
		newState        state
	}{
		{"New Pull Request Assigned to Issue", []pullRequest{{Number: 21, ID: "123", Branch: "feature/EPIC-1", Repo: "Rambatino/CHAID"}, {Number: 22, ID: "1", Branch: "feature/EPIC-1", Repo: "Rambatino/OTHER-REPO"}}, []issue{{ID: "123", Key: "EPIC-1", Title: "Title 1"}}, state{PullRequests: []pullRequest{}, Changes: []stateChange{}, Issues: []issue{{ID: "123", Key: "EPIC-1", Title: "Title 1", PullRequests: []pullRequest{{Number: 21, ID: "123", Branch: "feature/EPIC-1", Repo: "Rambatino/CHAID"}, {Number: 22, ID: "1", Branch: "feature/EPIC-1", Repo: "Rambatino/OTHER-REPO"}}}}}},
		{"New Pull Request Assigned to Issue and other one not", []pullRequest{{ID: "123", Branch: "feature/EPIC-2", Repo: "Rambatino/CHAID", author: "Rambatino", state: "open"}, {ID: "1", Branch: "feature/EPIC-1", Repo: "Rambatino/OTHER-REPO"}}, []issue{{ID: "123", Key: "EPIC-1", Title: "Title 1"}}, state{PullRequests: []pullRequest{{ID: "123", Branch: "feature/EPIC-2", Repo: "Rambatino/CHAID", author: "Rambatino", state: "open"}}, Changes: []stateChange{}, Issues: []issue{{ID: "123", Key: "EPIC-1", Title: "Title 1", PullRequests: []pullRequest{{ID: "1", Branch: "feature/EPIC-1", Repo: "Rambatino/OTHER-REPO"}}}}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			newState := formatState(tc.newIssues, tc.newPullRequests, []stateChange{})
			if !reflect.DeepEqual(newState, tc.newState) {
				t.Errorf("Change is not equal to expected\n%+v\n%+v	", newState, tc.newState)
			}
		})
	}
}
