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
		sslMode, err := m.getPostgresSSLMode()
		if err != nil {
			return nil, err
		}
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
			sslMode,
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
	case "firebird":
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		db, err := NewFirebirdStore(
			os.Getenv("DB_HOST"),
			port,
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
		return db, err
	default:
		return nil, fmt.Errorf("invalid DB_CONNECTION %s", dbConnection)
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
		return nil, fmt.Errorf("migration provider for type %s does not exists", migrationProvider)
	}
}

func (m *migrationInit) migrationCount(args []string) (int, error) {
	if len(args) > 2 {
		return strconv.Atoi(args[2])
	}

	return 0, nil
}

func (m *migrationInit) getPostgresSSLMode() (string, error) {
	envSSLMOde := os.Getenv("DB_SSLMODE")

	switch envSSLMOde {
	case "disable", "":
		return PgsSslMode.Disable, nil
	case "require":
		return PgsSslMode.Require, nil
	case "verify-ca":
		return PgsSslMode.VerifyCa, nil
	case "verify-full":
		return PgsSslMode.VerifyFull, nil
	case "prefer":
		return PgsSslMode.Prefer, nil
	case "allow":
		return PgsSslMode.Allow, nil

	default:
		return "", fmt.Errorf(`the provided DB_SSLMODE environment variable is invalid '%s'.,
Should be one of: disable, require, verify-ca, verify-full, prefer, allow:
If not set it will default to disable`,
			envSSLMOde,
		)
	}
}
