package dbconnector

import (
	"database/sql"
	"fmt"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/lib/pq"
)

const (
	driverPostgresQL = "postgres"
)

func newPostgresStore(
	host string,
	port int,
	user,
	password,
	dbName,
	sslMode string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbName,
		sslMode,
	)

	db, err := sql.Open(driverPostgresQL, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgresQL connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping PostgresQL database: %w", err)
	}

	return db, nil
}
