package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	envFileName = ".env.migrator"

	// Defaults
	defaultPrefix        = "olb"
	defaultMigrationPath = "./migrations"
	defaultDBPort        = 3306
	defaultSSLMode       = "disable"

	// Environment Variables
	envDBConnection  = "DB_CONNECTION"
	envMigrationPath = "MIGRATOR_MIGRATION_PATH"
	envDBHost        = "DB_HOST"
	envDBUserName    = "DB_USERNAME"
	envDBPassword    = "DB_PASSWORD"
	envDBDatabase    = "DB_DATABASE"
	envDBPort        = "DB_PORT"
	envDBPrefix      = "TABLE_PREFIX"
	envDBSSLMode     = "DB_SSLMODE"
)

// EnvironmentManager provides access to environment variables.
type EnvironmentManager interface {
	GetMigrationPath() string
	GetDBConnection() string
	GetDBHost() string
	GetDBUserName() string
	GetDBPassword() string
	GetDBDatabase() string
	GetDBPort() int
	GetDBPrefix() string
	GetDBPostgresSSLMode() string
}

// environment holds environment variables.
type environment struct {
	migrationPath string
	dbConnection  string
	dbHost        string
	dbUserName    string
	dbPassword    string
	dbDatabase    string
	dbPort        int
	dbPrefix      string
	dbSSLMode     string
}

// New creates a new environment manager.
func New() (EnvironmentManager, error) {
	e := &environment{}
	if err := e.loadEnv(); err != nil {
		return nil, err
	}
	e.loadValues()
	return e, nil
}

// loadEnv loads environment variables from the .env file.
func (e *environment) loadEnv() error {
	if _, err := os.Stat(envFileName); os.IsNotExist(err) {
		return nil // No .env file, use system env
	}
	return godotenv.Load(envFileName)
}

// loadValues initializes struct fields from environment variables.
func (e *environment) loadValues() {
	e.migrationPath = getEnvOrDefault(envMigrationPath, defaultMigrationPath)
	e.dbConnection = os.Getenv(envDBConnection)
	e.dbHost = os.Getenv(envDBHost)
	e.dbUserName = os.Getenv(envDBUserName)
	e.dbPassword = os.Getenv(envDBPassword)
	e.dbDatabase = os.Getenv(envDBDatabase)
	e.dbPort = getEnvAsInt(envDBPort, defaultDBPort)
	e.dbPrefix = getEnvOrDefault(envDBPrefix, defaultPrefix)
	e.dbSSLMode = getValidSSLMode(os.Getenv(envDBSSLMode))
}

// getEnvOrDefault retrieves an environment variable or returns a default value.
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt retrieves an environment variable as an int or returns a default.
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getValidSSLMode validates the SSL mode, returning a default if invalid.
func getValidSSLMode(mode string) string {
	validModes := map[string]bool{
		"disable": true, "require": true, "verify-ca": true,
		"verify-full": true, "prefer": true, "allow": true,
	}
	if validModes[mode] {
		return mode
	}
	return defaultSSLMode
}

// EnvironmentManager interface methods.
func (e *environment) GetMigrationPath() string     { return e.migrationPath }
func (e *environment) GetDBConnection() string      { return e.dbConnection }
func (e *environment) GetDBHost() string            { return e.dbHost }
func (e *environment) GetDBUserName() string        { return e.dbUserName }
func (e *environment) GetDBPassword() string        { return e.dbPassword }
func (e *environment) GetDBDatabase() string        { return e.dbDatabase }
func (e *environment) GetDBPort() int               { return e.dbPort }
func (e *environment) GetDBPrefix() string          { return e.dbPrefix }
func (e *environment) GetDBPostgresSSLMode() string { return e.dbSSLMode }
