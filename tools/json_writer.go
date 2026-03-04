package tools

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json") // tell the client what format we're sending
	w.WriteHeader(status)                              // set HTTP status code (200, 400, 404 etc)
	json.NewEncoder(w).Encode(v)                       // encode our struct as JSON and write it to the response
}
