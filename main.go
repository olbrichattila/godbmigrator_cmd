package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	migrator "github.com/olbrichattila/godbmigrator"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migrate()
		case "rollback":
			rollback()
		case "add":
			add()
		default:
			displayUsage()
		}
	} else {
		displayUsage()
	}
}

func migrate() {
	fmt.Println("Migrating")
	conn, provider, count, err := migrationInit()
	if err != nil {
		fmt.Println(err)
		return
	}

	migrator.Migrate(conn, provider, count)
}

func rollback() {
	fmt.Println("Rolling back")
	conn, provider, count, err := migrationInit()
	if err != nil {
		fmt.Println(err)
		return
	}

	migrator.Rollback(conn, provider, count)
}

func add() {
	fmt.Println("Adding new migration")
	customText := ""
	if len(os.Args) > 2 {
		customText = "-" + os.Args[2]
	}

	migrator.AddNewMigrationFiles(customText)
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
