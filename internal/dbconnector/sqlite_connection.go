package dbconnector

import (
	"database/sql"
	"fmt"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

const (
	driverTypeSqLite = "sqlite3"
)

func newSqliteStore(
	database Database,
	fileName string,
) (*sql.DB, error) {
	db, err := database.Open(driverTypeSqLite, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite connection: %w", err)
	}

	return db, nil
}
