package database

import (
	"bytes"
	"encoding/gob"

	"github.com/dgraph-io/badger"
	"github.com/k0kubun/pp"
)

type Manager struct {
	key   string
	state state
	db    *badger.DB
}

func (o *Manager) FindOrNew(bucketID string) state {
	var valCopy []byte
	err := o.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(bucketID))

		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return state{githubQuery: &githubQuery{}, jiraQuery: &jiraQuery{}}
	}
	dec := gob.NewDecoder(bytes.NewBuffer(valCopy))
	var s state
	err = dec.Decode(&s)

	if err != nil {
		pp.Println("Decode error", err)
		return state{githubQuery: &githubQuery{}, jiraQuery: &jiraQuery{}}
	}

	s.githubQuery = &githubQuery{}
	s.jiraQuery = &jiraQuery{}
	return s
}

func (s *state) Store(bucketID string) bool {
	return false
}

func NewManager(id string, db *badger.DB) Manager {
	return Manager{key: id, db: db}
}

func (o *Manager) Observe() (state, int) {
	s, i := CheckForStateChange(o.FindOrNew(o.key))
	if i > 0 {
		o.db.Update(func(txn *badger.Txn) error {
			var b bytes.Buffer
			dec := gob.NewEncoder(&b)
			err := dec.Encode(s)
			if err != nil {
				return err
			}
			return txn.Set([]byte(o.key), b.Bytes())
		})
	}
	return s, i
}
