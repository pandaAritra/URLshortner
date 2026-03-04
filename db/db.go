package db

import "sync"

// data base
// --------------
type Store interface {
	Save(code, uri string)
	Fetch(code string) (string, bool)
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

// ADD and Fetch in database
// --------------------------
func (store *UriStore) Save(code, uri string) {
	store.mu.Lock()
	store.urls[code] = uri
	store.mu.Unlock()
}
func (store *UriStore) Fetch(code string) (string, bool) {
	store.mu.RLock()
	uri, ok := store.urls[code]
	store.mu.RUnlock()
	return uri, ok

}
