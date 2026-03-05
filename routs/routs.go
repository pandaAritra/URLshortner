package routs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pandaAritra/URLshortner/db"
	"github.com/pandaAritra/URLshortner/tools"
)

// handler struct
// --------------
type Handlers struct {
	Store db.Store //stores the Urls : map
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
func (h *Handlers) Shortner(w http.ResponseWriter, r *http.Request) {
	var req BigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON body"})
		return
	}
	if req.URL == "" {
		tools.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Must contain URL"})
		return
	}

	code := tools.GenerateCode()
	h.Store.Save(code, req.URL)

	tools.WriteJSON(w, http.StatusCreated, ShortenResponse{
		Code:     code,
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", code),
	})
}
func (h *Handlers) FetchUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		tools.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Must contain code"})
	}
	url, ok := h.Store.Fetch(code)
	if !ok {
		// code not found — 404
		tools.WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: "short URL not found"})
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
