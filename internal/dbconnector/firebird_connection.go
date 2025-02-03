package dbconnector

import (
	"database/sql"
	"fmt"

	// This needs to be blank imported as not directly referenced, but required
	_ "github.com/nakagami/firebirdsql"
)

const driverFirebirdSQL = "firebirdsql"

// newFirebirdStore initializes a Firebird SQL connection.
func newFirebirdStore(
	database Database,
	host string,
	port int,
	user, password, dbName string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@%s:%d%s", user, password, host, port, dbName)

	db, err := database.Open(driverFirebirdSQL, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open FirebirdSQL connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping FirebirdSQL database: %w", err)
	}

	return db, nil
}
