package db

import "context"

// ADD and Fetch in database
// --------------------------
func (store *PostgresStore) Save(code, uri string) {

}
func (store *PostgresStore) Fetch(code string) (string, bool) {

}

// checks if this URL already exists
// -------------------------------------
func (store *PostgresStore) FindByURL(url string) (string, bool) {

}

// checks if Code Exists
// ---------------------
func (store *PostgresStore) Exists(code string) bool {

}

func (store *PostgresStore) Close() {
	store.conn.Close(context.Background())
}
