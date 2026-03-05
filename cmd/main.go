package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pandaAritra/URLshortner/db"
	"github.com/pandaAritra/URLshortner/handlers"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	connString := os.Getenv("POSTGRESSURL")
	if connString == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	store, err := db.NewPostgresStore(connString)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	h := &handlers.Handlers{Store: store}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shortner", h.Shortner)

	mux.HandleFunc("GET /{code}", h.FetchUrl)

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
