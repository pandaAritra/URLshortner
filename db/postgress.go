package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// data base
// ---------
type PostgresStore struct {
	conn *pgx.Conn // a single connection to the database
}

// data base instance
// ------------------
func NewPostgresStore(connString string) (*PostgresStore, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	store := &PostgresStore{conn: conn}

	// run migrations on every startup
	if err := store.migrate(); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return store, nil
}
