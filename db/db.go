package db

import "sync"

// data base
// --------------
type Store interface {
	Save(code, uri string)
	Fetch(code string) (string, bool)
	FindByURL(url string) (string, bool)
	Exists(code string) bool
}
type UriStore struct {
	mu   sync.RWMutex
	urls map[string]string
}

// New data base instance
// -----------------------
func NewUriStore() *UriStore {
	return &UriStore{
		urls: make(map[string]string),
	}
}
