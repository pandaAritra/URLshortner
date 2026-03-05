package db

import "context"

// migrate creates the tables if they don't exist.
// Runs every time the server starts — safe because of IF NOT EXISTS.
func (s *PostgresStore) migrate() error {
	_, err := s.conn.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS urls (
            id         SERIAL PRIMARY KEY,
            code       VARCHAR(10)  UNIQUE NOT NULL,
            url        TEXT         NOT NULL,
            created_at TIMESTAMP    DEFAULT NOW()
        );

        CREATE INDEX IF NOT EXISTS idx_url ON urls(url);
    `)
	return err
}
