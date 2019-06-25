package database

type Manager struct {
	key   string
	state state
}

func FindOrNew(bucketID string) state {
	return state{githubQuery: &githubQuery{}, jiraQuery: &jiraQuery{}}
}

func (s *state) Store(bucketID string) bool {
	return false
}

func NewManager(id string) Manager {
	return Manager{key: id}
}

func (o *Manager) Observe() (state, int) {
	return CheckForStateChange(FindOrNew(""), "hello")
}
