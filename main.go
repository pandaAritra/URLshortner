package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// data base
// --------------
type Store interface {
	Save(code, uri string)
	Get(code string) (string, bool)
}
type Uristore struct {
	mu  sync.RWMutex
	Box map[string]string
}

// Save and Get
// -------------
func (store *Uristore) Save(code, uri string) {
	store.mu.Lock()
	store.Box[code] = uri
	store.mu.Unlock()
}
func (store *Uristore) Get(code string) (string, bool) {
	store.mu.Lock()
	uri, ok := store.Box[code]
	store.mu.Unlock()
	if ok {
		return uri, ok
	}

	return "", ok

}

// request handlers
// ------------------
func shortner(w http.ResponseWriter, r *http.Request) {

}
func getUrl(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shortner", shortner)

	mux.HandleFunc("GET /{code}", getUrl)

	//server config

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  // max time to read the full request
		WriteTimeout: 10 * time.Second, // max time to write the full response
		IdleTimeout:  60 * time.Second, // max time to keep idle connections alive
	}

	println("server starting")

	log.Fatal(server.ListenAndServe())

}
