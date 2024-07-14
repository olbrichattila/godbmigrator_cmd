package migrator

import (
	"database/sql"
	"fmt"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/go-sql-driver/mysql"
)

func newMysqlStore(
	host string,
	port int,
	user,
	password,
	dbname string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		user,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
