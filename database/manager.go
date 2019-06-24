package database

import (
	"time"
)

const MAX_QUERIERS = 100

type Manager interface {
	AddObserver(id string)
	RemoveObserver(id string)
}

type manager struct {
	observers map[string]*observer
}

func NewManager() Manager {
	return manager{map[string]*observer{}}
}

func (m manager) AddObserver(id string) {
	o := NewObserver(id)
	m.observers[id] = &o
}

func (m manager) RemoveObserver(id string) {

}

type observer struct {
	key          string
	shouldCancel <-chan bool
}

func NewObserver(id string) observer {
	o := observer{id, make(chan bool)}
	go o.observe()
	return o
}

func (o *observer) observe() {
	for {
		select {
		case <-time.After(60 * time.Second):
			// trigger check
		case <-o.shouldCancel:
		}
	}
}
