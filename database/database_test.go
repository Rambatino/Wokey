package database

import (
	reflect "reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	gomock "github.com/golang/mock/gomock"
)

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
		hasChanged          bool
	}{
		{"New Issue Comment", []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1"}}, []pullRequest{}, []issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", Comments: []string{"hello"}}}, []stateChange{{Type: "NEW_ISSUE_COMMENT", Message: "'Title 1' [EPIC/EPIC-23] has 1 new comment(s)", CreatedAt: now}}, true},
		// {"Issue Comment Deleted", []*pullRequest{}, []*issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1"}}, []*pullRequest{}, []*issue{{ID: "123", Key: "EPIC/EPIC-23", Title: "Title 1", Comments: []string{"hello"}}}, []*stateChange{{Type: "NEW_ISSUE_COMMENT", Message: "'Title 1' [EPIC/EPIC-23] has 1 new comment(s)", CreatedAt: now}}, true},
		// {"New PR Comment", 0, -2, -2},
		// {"New Column Change for Issue", 0, 0, 0},
		// {"PR has been approved", 0, 0, 0},
		// {"Changes requested for PR", 0, 0, 0},
		// {"CI failed for PR", 0, 0, 0},
		// {"CI Succeeded for PR", 0, 0, 0},
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
