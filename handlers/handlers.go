package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pandaAritra/URLshortner/db"
	"github.com/pandaAritra/URLshortner/models"
	"github.com/pandaAritra/URLshortner/tools"
)

// handler struct
// --------------
type Handlers struct {
	Store db.Store //stores the Urls : map
}

// request handler shortner
// ------------------------
func (h *Handlers) Shortner(w http.ResponseWriter, r *http.Request) {
	var req models.BigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid JSON body"})
		return
	}
	if req.URL == "" {
		tools.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "Must contain URL"})
		return
	}

	code, ok := h.Store.FindByURL(req.URL)
	if !ok {
		code = tools.GenerateCode()
		for h.Store.Exists(code) {
			code = tools.GenerateCode()
		}
		h.Store.Save(code, req.URL)
	}

	tools.WriteJSON(w, http.StatusCreated, models.ShortenResponse{
		Code:     code,
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", code),
	})
}

// request handler fetch
// ---------------------
func (h *Handlers) FetchUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		tools.WriteJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "Must contain code"})
	}
	url, ok := h.Store.Fetch(code)
	if !ok {
		// code not found — 404
		tools.WriteJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "short URL not found"})
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
