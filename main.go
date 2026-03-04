package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pandaAritra/URLshortner/db"
	"github.com/pandaAritra/URLshortner/tools"
)

// handler struct
// --------------
type Handlers struct {
	store db.Store //stores the Urls : map
}

type BigRequest struct {
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
	var req BigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON body"})
	}
	if req.URL == "" {
		tools.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Must contain URL"})
	}

	code := tools.GenerateCode()
	h.store.Save(code, req.URL)

	tools.WriteJSON(w, http.StatusCreated, ShortenResponse{
		Code:     code,
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", code),
	})
}
func (h *Handlers) fetchUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		tools.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Must contain code"})
	}
	url, ok := h.store.Fetch(code)
	if !ok {
		// code not found — 404
		tools.WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: "short URL not found"})
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	store := db.NewUriStore()
	h := &Handlers{store: store}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shortner", h.shortner)

	mux.HandleFunc("GET /{code}", h.fetchUrl)

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
