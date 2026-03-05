package db

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

// checks if this URL already exists
// -------------------------------------
func (store *UriStore) FindByURL(url string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	for code, storedURL := range store.urls {
		if storedURL == url {
			return code, true // already exists, return the existing code
		}
	}
	return "", false
}

// checks if Code Exists
// ---------------------
func (s *UriStore) Exists(code string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.urls[code]
	return ok
}
