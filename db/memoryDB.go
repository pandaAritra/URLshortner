package db

import "sync"

// data base
// ----------
type InMemoryStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

// New data base instance
// -----------------------
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		urls: make(map[string]string),
	}
}
