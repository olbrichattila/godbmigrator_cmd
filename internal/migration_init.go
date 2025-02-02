package migrator

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

const (
	defaultPrefix = "olb"
	envDBHost     = "DB_HOST"
	envDBUserName = "DB_USERNAME"
	envDBPassword = "DB_PASSWORD"
	envDBDatabase = "DB_DATABASE"
	envDBPort     = "DB_PORT"
	envDBPrefix   = "TABLE_PREFIX"
)

type migrationInitInterface interface {
	migrationInit(args []string, initMigrationTables bool) (*sql.DB, string, int, error)
}

type migrationInit struct {
}

func newMigrationInit() *migrationInit {
	return &migrationInit{}
}

func (m *migrationInit) migrationInit(args []string, initMigrationTables bool) (*sql.DB, string, int, error) {
	conn, err := m.connection()
	if err != nil {
		return nil, "", 0, err
	}

	count, err := m.migrationCount(args)
	if err != nil {
		return nil, "", 0, err
	}

	prefix := m.getDBPrefix()

	return conn, prefix, count, err
}

func (m *migrationInit) getDBPrefix() string {
	prefix := os.Getenv(envDBPrefix)
	if prefix == "" {
		return defaultPrefix
	}

	return prefix
}

func (m *migrationInit) connection() (*sql.DB, error) {
	dbConnection := os.Getenv("DB_CONNECTION")

	switch dbConnection {
	case providerSqLiteDataType:
		db, err := newSqliteStore(os.Getenv(envDBDatabase))
		return db, err
	case providerPgSQLDataType:
		sslMode, err := m.getPostgresSSLMode()
		if err != nil {
			return nil, err
		}
		port, err := strconv.Atoi(os.Getenv(envDBPort))
		if err != nil {
			return nil, err
		}
		db, err := newPostgresStore(
			os.Getenv(envDBHost),
			port,
			os.Getenv(envDBUserName),
			os.Getenv(envDBPassword),
			os.Getenv(envDBDatabase),
			sslMode,
		)
		return db, err
	case providerMySQLDataType:
		port, err := strconv.Atoi(os.Getenv(envDBPort))
		if err != nil {
			return nil, err
		}
		db, err := newMysqlStore(
			os.Getenv(envDBHost),
			port,
			os.Getenv(envDBUserName),
			os.Getenv(envDBPassword),
			os.Getenv(envDBDatabase),
		)
		return db, err
	case providerFirebirdDataType:
		port, err := strconv.Atoi(os.Getenv(envDBPort))
		if err != nil {
			return nil, err
		}
		db, err := newFirebirdStore(
			os.Getenv(envDBHost),
			port,
			os.Getenv(envDBUserName),
			os.Getenv(envDBPassword),
			os.Getenv(envDBDatabase),
		)
		return db, err
	default:
		return nil, fmt.Errorf("invalid DB_CONNECTION %s", dbConnection)
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
		return pgsSSLMode.Disable, nil
	case "require":
		return pgsSSLMode.Require, nil
	case "verify-ca":
		return pgsSSLMode.VerifyCa, nil
	case "verify-full":
		return pgsSSLMode.VerifyFull, nil
	case "prefer":
		return pgsSSLMode.Prefer, nil
	case "allow":
		return pgsSSLMode.Allow, nil

	default:
		return "", fmt.Errorf(`the provided DB_SSLMODE environment variable is invalid '%s'.,
Should be one of: disable, require, verify-ca, verify-full, prefer, allow:
If not set it will default to disable`,
			envSSLMOde,
		)
	}
}
