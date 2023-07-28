package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	migrator "github.com/olbrichattila/godbmigrator"
)

const defaultMigrationPath = "./migrations"

type providerInterface interface {
}

type migratorInterface interface {
	Migrate(*sql.DB, migrator.MigrationProvider, string, int) error
	Rollback(*sql.DB, migrator.MigrationProvider, string, int) error
	Refresh(*sql.DB, migrator.MigrationProvider, string) error
	AddNewMigrationFiles(string, string)
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	migrationAdapter := NewMigrationAdapter()
	migrationInit := newMigrationInit()
	if err := routeCommandLineParameters(os.Args, migrationAdapter, migrationInit); err != nil {
		displayUsage()
	}
}

func routeCommandLineParameters(args []string, migrationAdapter migratorInterface, migrationInit migrationInitInterface) error {
	if len(args) > 1 {
		switch args[1] {
		case "migrate":
			migrate(args, migrationAdapter, migrationInit)
		case "rollback":
			rollback(args, migrationAdapter, migrationInit)
		case "refresh":
			refresh(args, migrationAdapter, migrationInit)
		case "add":
			add(args, migrationAdapter)
		default:
			return fmt.Errorf("Cannot find command")
		}
	} else {
		return fmt.Errorf("Invalid parameter count")
	}

	return nil
}

func migrate(args []string, migrationAdapter migratorInterface, migrationInit migrationInitInterface) {
	fmt.Println("Migrating")
	conn, provider, count, err := migrationInit.migrationInit(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	err = migrationAdapter.Migrate(conn, provider, migrationPath, count)
	if err != nil {
		fmt.Println(err)
	}
}

func rollback(args []string, migrationAdapter migratorInterface, migrationInit migrationInitInterface) {
	fmt.Println("Rolling back")
	conn, provider, count, err := migrationInit.migrationInit(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	err = migrationAdapter.Rollback(conn, provider, migrationPath, count)
	if err != nil {
		fmt.Println(err)
	}
}

func refresh(args []string, migrationAdapter migratorInterface, migrationInit migrationInitInterface) {
	fmt.Println("Rolling back")
	conn, provider, _, err := migrationInit.migrationInit(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	err = migrationAdapter.Refresh(conn, provider, migrationPath)
	if err != nil {
		fmt.Println(err)
	}
}

func add(args []string, migrationAdapter migratorInterface) {
	fmt.Println("Adding new migration")
	customText := ""
	if len(args) > 2 {
		customText = "-" + args[2]
	}

	migrationPath := migrationPath()
	migrationAdapter.AddNewMigrationFiles(migrationPath, customText)
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
