package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	migrator "github.com/olbrichattila/godbmigrator"
)

type migrationInitInterface interface {
	migrationInit(args []string) (*sql.DB, migrator.MigrationProvider, int, error)
}

type migrationInit struct {
}

func newMigrationInit() *migrationInit {
	return &migrationInit{}
}

func (m *migrationInit) migrationInit(args []string) (*sql.DB, migrator.MigrationProvider, int, error) {
	conn, err := m.connection()
	if err != nil {
		return nil, nil, 0, err
	}

	provider, err := m.provider(conn)
	if err != nil {
		return nil, nil, 0, err
	}

	count, err := m.migrationCount(args)
	if err != nil {
		return nil, nil, 0, err
	}

	return conn, provider, count, err
}

func (m *migrationInit) connection() (*sql.DB, error) {
	dbConnection := os.Getenv("DB_CONNECTION")

	switch dbConnection {
	case "sqlite":
		db, err := NewSqliteStore(os.Getenv("DB_DATABASE"))
		return db, err
	case "pgsql":
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		db, err := NewPostgresStore(
			os.Getenv("DB_HOST"),
			port,
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			PgsSslMode.Disable,
		)
		return db, err
	case "mysql":
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		db, err := NewMysqlStore(
			os.Getenv("DB_HOST"),
			port,
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
		return db, err
	default:
		return nil, fmt.Errorf("Invalid DB_CONNECTION %s", dbConnection)
	}
}

func (m *migrationInit) provider(db *sql.DB) (migrator.MigrationProvider, error) {
	migrationProvider := os.Getenv("MIGRATOR_MIGRATION_PROVIDER")

	switch migrationProvider {
	case "db", "":
		return migrator.NewMigrationProvider("db", db)
	case "json":
		return migrator.NewMigrationProvider("json", nil)
	default:
		return nil, fmt.Errorf("Migration provider for type %s does not exists", migrationProvider)
	}
}

func (m *migrationInit) migrationCount(args []string) (int, error) {
	if len(args) > 2 {
		return strconv.Atoi(args[2])
	}

	return 0, nil
}
