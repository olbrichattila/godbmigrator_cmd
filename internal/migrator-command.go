package migratorcommand

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	migrator "github.com/olbrichattila/godbmigrator"
	"github.com/olbrichattila/godbmigrator/messager"
	"github.com/olbrichattila/godbmigrator_cmd/internal/dbconnector"
	"github.com/olbrichattila/godbmigrator_cmd/internal/env"
)

const (
	reset              = "\033[0m"
	red                = "\033[31m"
	genericMessageType = -1
)

// Init will initiate the command line utility
func Init(messageCallback messager.CallbackFunc) {
	migrator.SubscribeToMessages(messageCallback)

	// Load environment configuration
	environment, err := env.New()
	if err != nil {
		messageCallback(genericMessageType, err.Error())
		return
	}

	// Establish database connection
	db, err := dbconnector.New(environment).GetConnection()
	if err != nil {
		messageCallback(genericMessageType, err.Error())
		return
	}

	// Parse command-line arguments
	args := os.Args
	if len(args) < 2 {
		messageCallback(genericMessageType, "invalid parameter count")
		return
	}

	dbPrefix := environment.GetDBPrefix()
	migrationFilePath := environment.GetMigrationPath()
	command := args[1]
	commandArgs := args[2:]

	runCommand(command, commandArgs, db, dbPrefix, migrationFilePath, messageCallback)
}

func runCommand(command string, args []string, db *sql.DB, dbPrefix, migrationPath string, messageCallback messager.CallbackFunc) {
	commands := map[string]func(){
		"migrate": func() {
			count := parseMigrationCount(args)
			handleError(migrator.Migrate(db, dbPrefix, migrationPath, count), messageCallback)
		},
		"rollback": func() {
			count := parseMigrationCount(args)
			handleError(migrator.Rollback(db, dbPrefix, migrationPath, count), messageCallback)
		},
		"refresh": func() {
			handleError(migrator.Refresh(db, dbPrefix, migrationPath), messageCallback)
		},
		"report": func() {
			report, err := migrator.Report(db, dbPrefix, migrationPath)
			if err != nil {
				messageCallback(genericMessageType, err.Error())
			} else {
				messageCallback(genericMessageType, report)
			}
		},
		"add": func() {
			if len(args) == 0 {
				messageCallback(genericMessageType, "missing migration filename")
				return
			}
			handleError(migrator.AddNewMigrationFiles(migrationPath, args[0]), messageCallback)
		},
		"validate": func() {
			result := migrator.ChecksumValidation(db, dbPrefix, migrationPath)
			messageCallback(genericMessageType, strings.Join(result, "\n"))
		},
		"save-baseline": func() {
			handleError(migrator.SaveBaseline(db, migrationPath), messageCallback)
		},
		"restore-baseline": func() {
			handleError(migrator.LoadBaseline(db, migrationPath), messageCallback)
		},
		"help": func() {
			displayFullHelp()
		},
	}

	if cmdFunc, exists := commands[command]; exists {
		cmdFunc()
	} else {
		messageCallback(genericMessageType, fmt.Sprintf("unknown command: %s", command))
	}
}

func parseMigrationCount(args []string) int {
	if len(args) > 0 {
		if count, err := strconv.Atoi(args[0]); err == nil {
			return count
		}
	}
	return 0
}

func handleError(err error, messageCallback messager.CallbackFunc) {
	if err != nil {
		messageCallback(
			genericMessageType,
			decorateErrorMessage(
				err.Error(),
			),
		)
	}
}

func decorateErrorMessage(message string) string {
	return red + message + reset
}
