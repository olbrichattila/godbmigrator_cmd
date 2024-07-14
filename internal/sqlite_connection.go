package migrator

import (
	"database/sql"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func newSqliteStore(fileName string) (*sql.DB, error) {
	db, err := sql.Open(driverTypeSqLite, fileName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
