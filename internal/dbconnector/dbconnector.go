// Package dbconnector returns the *sql.DB according to environment variables
package dbconnector

import (
	"database/sql"
	"fmt"

	"github.com/olbrichattila/godbmigrator_cmd/internal/env"
)

// Database providers
const (
	providerSqLite   = "sqlite"
	providerPgSQL    = "pgsql"
	providerMySQL    = "mysql"
	providerFirebird = "firebird"
)

// New returns a new DBConnector instance.
func New(env env.EnvironmentManager) DBConnector {
	return &dbc{env: env}
}

// DBConnector abstracts database connection handling.
type DBConnector interface {
	GetConnection() (*sql.DB, error)
}

type dbc struct {
	env env.EnvironmentManager
}

// GetConnection returns the appropriate database connection.
func (d *dbc) GetConnection() (*sql.DB, error) {
	dbConnection := d.env.GetDBConnection()
	connectionHandlers := map[string]func() (*sql.DB, error){
		providerSqLite: func() (*sql.DB, error) {
			return newSqliteStore(d.env.GetDBDatabase())
		},
		providerPgSQL: func() (*sql.DB, error) {
			return newPostgresStore(
				d.env.GetDBHost(),
				d.env.GetDBPort(),
				d.env.GetDBUserName(),
				d.env.GetDBPassword(),
				d.env.GetDBDatabase(),
				d.env.GetDBPostgresSSLMode(),
			)
		},
		providerMySQL: func() (*sql.DB, error) {
			return newMysqlStore(
				d.env.GetDBHost(),
				d.env.GetDBPort(),
				d.env.GetDBUserName(),
				d.env.GetDBPassword(),
				d.env.GetDBDatabase(),
			)
		},
		providerFirebird: func() (*sql.DB, error) {
			return newFirebirdStore(
				d.env.GetDBHost(),
				d.env.GetDBPort(),
				d.env.GetDBUserName(),
				d.env.GetDBPassword(),
				d.env.GetDBDatabase(),
			)
		},
	}

	if handler, exists := connectionHandlers[dbConnection]; exists {
		return handler()
	}

	return nil, fmt.Errorf("invalid DB_CONNECTION: %s", dbConnection)
}
