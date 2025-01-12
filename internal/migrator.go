package migrator

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	migrator "github.com/olbrichattila/godbmigrator"
)

const (
	envFileName          = ".env.migrator"
	defaultMigrationPath = "./migrations"
	messageRollingBack   = "Rolling back"
)

type migratorInterface interface {
	Migrate(*sql.DB, migrator.MigrationProvider, string, int) error
	Rollback(*sql.DB, migrator.MigrationProvider, string, int) error
	Refresh(*sql.DB, migrator.MigrationProvider, string) error
	Report(*sql.DB, migrator.MigrationProvider, string) (string, error)
	ChecksumValidation(*sql.DB, migrator.MigrationProvider, string) []string
	AddNewMigrationFiles(string, string) error
}

// Init starts migration command, reads .env.migrator and command line arguments, and execute what was requested
func Init() {
	if err := loadEnv(); err != nil {
		fmt.Printf("Error loading %s file:%s\n", err.Error(), envFileName)
		return
	}

	migrationAdapter := newMigrationAdapter()
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
		case "report":
			report(args, migrationAdapter, migrationInit)
		case "add":
			add(args, migrationAdapter)
		case "validate":
			validate(args, migrationAdapter, migrationInit)
		case "help":
			displayFullHelp()
		default:
			return fmt.Errorf("cannot find command")
		}
	} else {
		return fmt.Errorf("invalid parameter count")
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
	fmt.Println(messageRollingBack)
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
	fmt.Println(messageRollingBack)
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

func report(args []string, migrationAdapter migratorInterface, migrationInit migrationInitInterface) {
	fmt.Println("Migration report")
	conn, provider, _, err := migrationInit.migrationInit(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	report, err := migrationAdapter.Report(conn, provider, migrationPath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(report)
}

func add(args []string, migrationAdapter migratorInterface) {
	fmt.Println("Adding new migration")
	customText := ""
	if len(args) > 2 {
		customText = "-" + args[2]
	}

	migrationPath := migrationPath()
	err := migrationAdapter.AddNewMigrationFiles(migrationPath, customText)
	if err != nil {
		fmt.Println(err)
	}
}

func validate(args []string, migrationAdapter migratorInterface, migrationInit migrationInitInterface) {
	fmt.Println("Validating")
	conn, provider, _, err := migrationInit.migrationInit(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	migrationPath := migrationPath()
	errors := migrationAdapter.ChecksumValidation(conn, provider, migrationPath)

	for _, errorString := range errors {
		fmt.Println(" - " + errorString)
	}
	
	fmt.Println("Done.")
}

func displayUsage() {
	fmt.Printf(`
Usage:
	migrator migrate
	migrator rollback
	migrator migrate 2
	migrator rollback 2
	migrator report
	migrator validate
	migrator add <optional suffix>

For help how to set up:
	migrator help

The number of rollbacks and migrates are not mandatory.
If it is set, for rollbacks it only apply for the last rollback batch
validate verifies if any migration file changed since applied
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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func loadEnv() error {
	if fileExists(envFileName) {
		if err := godotenv.Load(envFileName); err != nil {
			return err
		}
	}

	return nil
}
