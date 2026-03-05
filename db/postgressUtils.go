package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// ADD and Fetch in database
// --------------------------
func (store *PostgresStore) Save(code, uri string) {
	_, err := store.conn.Exec(
		context.Background(),
		"INSERT INTO urls (code, url) VALUES ($1, $2)",
		code, uri,
	)
	if err != nil {
		fmt.Printf("error saving url: %v\n", err)
	}
}
func (store *PostgresStore) Fetch(code string) (string, bool) {
	var url string

	err := store.conn.QueryRow(
		context.Background(),
		"SELECT url FROM urls WHERE code = $1",
		code,
	).Scan(&url) // Scan reads the result into our url variable

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", false // code not found — not an error, just missing
		}
		fmt.Printf("error fetching url: %v\n", err)
		return "", false
	}

	return url, true
}

// checks if this URL already exists
// -------------------------------------
func (store *PostgresStore) FindByURL(url string) (string, bool) {
	var code string

	err := store.conn.QueryRow(
		context.Background(),
		"SELECT code FROM urls WHERE url = $1",
		url,
	).Scan(&code)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", false // url not found — not an error
		}
		fmt.Printf("error finding url: %v\n", err)
		return "", false
	}

	return code, true
}

// checks if Code Exists
// ---------------------
func (store *PostgresStore) Exists(code string) bool {
	var exists bool

	err := store.conn.QueryRow(
		context.Background(),
		"SELECT EXISTS(SELECT 1 FROM urls WHERE code = $1)",
		code,
	).Scan(&exists)

	if err != nil {
		fmt.Printf("error checking code existence: %v\n", err)
		return false
	}

	return exists
}

func (store *PostgresStore) Close() {
	store.conn.Close(context.Background())
}
