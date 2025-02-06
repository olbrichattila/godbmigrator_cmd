// Package migratorcommand is a wrapper around db migrator github.com/olbrichattila/godbmigrator to expose it command line
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

	// Parse command-line arguments
	args := os.Args
	if len(args) < 2 {
		messageCallback(genericMessageType, fmt.Sprintf("invalid parameter count.\nUsage:\n%s", getUsageAsString()))
		return
	}

	command := args[1]
	commandArgs := args[2:]

	runCommand(command, commandArgs, messageCallback)
}

func connect(messageCallback messager.CallbackFunc) (*sql.DB, string, string, bool) {
	environment, err := env.New()
	if err != nil {
		messageCallback(genericMessageType, err.Error())
		return nil, "", "", false
	}

	db, err := dbconnector.New(environment).GetConnection(dbconnector.NewDB())
	if err != nil {
		messageCallback(genericMessageType, err.Error())
		return nil, "", "", false
	}

	dbPrefix := environment.GetDBPrefix()
	migrationFilePath := environment.GetMigrationPath()

	return db, dbPrefix, migrationFilePath, true

}

func runCommand(command string, args []string, messageCallback messager.CallbackFunc) {
	commands := map[string]func(){
		"migrate": func() {
			if db, dbPrefix, migrationPath, ok := connect(messageCallback); ok {
				count := parseMigrationCount(args)
				handleError(migrator.Migrate(db, dbPrefix, migrationPath, count), messageCallback)
			}
		},
		"rollback": func() {
			if db, dbPrefix, migrationPath, ok := connect(messageCallback); ok {
				count := parseMigrationCount(args)
				handleError(migrator.Rollback(db, dbPrefix, migrationPath, count), messageCallback)
			}
		},
		"refresh": func() {
			if db, dbPrefix, migrationPath, ok := connect(messageCallback); ok {
				handleError(migrator.Refresh(db, dbPrefix, migrationPath), messageCallback)
			}
		},
		"report": func() {
			if db, dbPrefix, migrationPath, ok := connect(messageCallback); ok {
				report, err := migrator.Report(db, dbPrefix, migrationPath)
				if err != nil {
					messageCallback(genericMessageType, err.Error())
				} else {
					messageCallback(genericMessageType, report)
				}
			}

		},
		"add": func() {
			if len(args) == 0 {
				messageCallback(genericMessageType, "missing migration filename")
				return
			}
			if _, _, migrationPath, ok := connect(messageCallback); ok {
				handleError(migrator.AddNewMigrationFiles(migrationPath, args[0]), messageCallback)
			}
		},
		"validate": func() {
			if db, dbPrefix, migrationPath, ok := connect(messageCallback); ok {
				result := migrator.ChecksumValidation(db, dbPrefix, migrationPath)
				messageCallback(genericMessageType, strings.Join(result, "\n"))
			}
		},
		"save-baseline": func() {
			if db, _, migrationPath, ok := connect(messageCallback); ok {
				handleError(migrator.SaveBaseline(db, migrationPath), messageCallback)
			}
		},
		"restore-baseline": func() {
			if db, _, migrationPath, ok := connect(messageCallback); ok {
				handleError(migrator.LoadBaseline(db, migrationPath), messageCallback)
			}
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

func getUsageAsString() string {
	return `migrator migrate
migrator rollback
migrator migrate 2
migrator rollback 2
migrator refresh
migrator report
migrator validate
migrator save-baseline
migrator restore-baseline`
}
