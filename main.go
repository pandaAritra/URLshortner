package main

import (
	"log"
	"net/http"
	"time"
)

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
