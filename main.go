package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	migrator "github.com/olbrichattila/godbmigrator"
)

const defaultMigrationPath = "./migrations"

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	if err := routeCommandLineParameters(); err != nil {
		displayUsage()
	}
}

func routeCommandLineParameters() error {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migrate()
		case "rollback":
			rollback()
		case "add":
			add()
		default:
			return fmt.Errorf("Cannot find command")
		}
	} else {
		return fmt.Errorf("Invalid parameter count")
	}

	return nil
}

func migrate() {
	fmt.Println("Migrating")
	conn, provider, count, err := migrationInit()
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	err = migrator.Migrate(conn, provider, migrationPath, count)
	if err != nil {
		fmt.Println(err)
	}
}

func rollback() {
	fmt.Println("Rolling back")
	conn, provider, count, err := migrationInit()
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	err = migrator.Rollback(conn, provider, migrationPath, count)
	if err != nil {
		fmt.Println(err)
	}
}

func add() {
	fmt.Println("Adding new migration")
	customText := ""
	if len(os.Args) > 2 {
		customText = "-" + os.Args[2]
	}

	migrationPath := migrationPath()
	migrator.AddNewMigrationFiles(migrationPath, customText)
}

func displayUsage() {
	fmt.Printf(`
Usage:
	migrator migrate
	migrator rollback
	migrator migrate 2
	migrator rollback 2

The number of rollbacks and migrates are not mandatory.
If it is set, for rollbacks it only apply for the last rollback batch

`)
}

func migrationCount() (int, error) {
	if len(os.Args) > 2 {
		return strconv.Atoi(os.Args[2])
	}

	return 0, nil
}

func migrationInit() (*sql.DB, migrator.MigrationProvider, int, error) {
	conn, err := connection()
	if err != nil {
		return nil, nil, 0, err
	}

	provider, err := provider(conn)
	if err != nil {
		return nil, nil, 0, err
	}

	count, err := migrationCount()
	if err != nil {
		return nil, nil, 0, err
	}

	return conn, provider, count, err
}

func connection() (*sql.DB, error) {
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

func provider(db *sql.DB) (migrator.MigrationProvider, error) {
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

func migrationPath() string {
	migrationPath := os.Getenv("MIGRATOR_MIGRATION_PATH")
	if migrationPath == "" {
		return defaultMigrationPath
	}

	return removeLastSlashOrBackslash(migrationPath)
}

func removeLastSlashOrBackslash(inputString string) string {
	if len(inputString) <= 1 {
		return inputString
	}

	lastChar := inputString[len(inputString)-1:]
	if lastChar == "/" || lastChar == "\\" {
		return inputString[:len(inputString)-1]
	}

	return inputString
}
