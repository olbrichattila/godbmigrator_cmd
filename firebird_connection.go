package main

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

func NewFirebirdStore(
	host string,
	port int,
	user,
	password,
	dbname string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@%s:%d%s",
		user,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("firebirdsql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
