package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type pgsSslMode struct {
	Disable    string
	Require    string
	VerifyCa   string
	VerifyFull string
	Prefer     string
	Allow      string
}

var PgsSslMode = &pgsSslMode{
	Disable:    "disable",
	Require:    "require",
	VerifyCa:   "verify-ca",
	VerifyFull: "verify-full",
	Prefer:     "prefer",
	Allow:      "allow",
}

func NewPostgresStore(
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
