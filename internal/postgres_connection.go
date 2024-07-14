package migrator

import (
	"database/sql"
	"fmt"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/lib/pq"
)

type pgsSSLModeTypes struct {
	Disable    string
	Require    string
	VerifyCa   string
	VerifyFull string
	Prefer     string
	Allow      string
}

var pgsSSLMode = &pgsSSLModeTypes{
	Disable:    "disable",
	Require:    "require",
	VerifyCa:   "verify-ca",
	VerifyFull: "verify-full",
	Prefer:     "prefer",
	Allow:      "allow",
}

func newPostgresStore(
	host string,
	port int,
	user,
	password,
	dbname,
	sslmode string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		sslmode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
