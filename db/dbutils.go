package db

// ADD and Fetch in database
// --------------------------
func (store *UriStore) Save(code, uri string) {
	store.mu.Lock()
	defer store.mu.RUnlock()
	store.urls[code] = uri
}
func (store *UriStore) Fetch(code string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	uri, ok := store.urls[code]
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
func (store *UriStore) Exists(code string) bool {
	store.mu.RLock()
	defer store.mu.RUnlock()
	_, ok := store.urls[code]
	return ok
}
