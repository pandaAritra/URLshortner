package main

import (
	"log"
	"net/http"
	"time"

	"github.com/pandaAritra/URLshortner/db"
)

// handler struct
// --------------
type Handlers struct {
	store db.Store
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// request handlers
// ------------------
func (h *Handlers) shortner(w http.ResponseWriter, r *http.Request) {

}
func (h *Handlers) fetchUrl(w http.ResponseWriter, r *http.Request) {

}

func main() {
	store := db.NewUriStore()
	h := &Handlers{store: store}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shortner", h.shortner)

	mux.HandleFunc("fetch /{code}", h.fetchUrl)

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
