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
	Urls map[string]string
}

// New data base instance
// -----------------------
func NewUriStore() *UriStore {
	return &UriStore{
		Urls: make(map[string]string),
	}
}

// ADD and Fetch in database
// --------------------------
func (store *UriStore) Save(code, uri string) {
	store.mu.Lock()
	store.Urls[code] = uri
	store.mu.Unlock()
}
func (store *UriStore) Fetch(code string) (string, bool) {
	store.mu.RLock()
	uri, ok := store.Urls[code]
	store.mu.RUnlock()
	return uri, ok

}
