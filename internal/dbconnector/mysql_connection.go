package dbconnector

import (
	"database/sql"
	"fmt"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverMySQL = "mysql"
)

func newMysqlStore(
	database Database,
	host string,
	port int,
	user,
	password,
	dbName string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		user,
		password,
		host,
		port,
		dbName,
	)

	db, err := database.Open(driverMySQL, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping MySQL database: %w", err)
	}

	return db, nil
}
