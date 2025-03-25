// Package migratorcommand is a wrapper around db migrator github.com/olbrichattila/godbmigrator to expose it command line
package migratorcommand

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	migrator "github.com/olbrichattila/godbmigrator"
	"github.com/olbrichattila/godbmigrator_cmd/internal/dbconnector"
	"github.com/olbrichattila/godbmigrator_cmd/internal/env"
	"github.com/olbrichattila/godbmigrator_cmd/internal/parameters"
	"github.com/olbrichattila/godbmigrator_cmd/internal/server"
)

const (
	reset              = "\033[0m"
	red                = "\033[31m"
	genericMessageType = -1
)

// Init will initiate the command line utility
func Init(messageCallback func(int, string)) {
	// Parse command-line arguments
	args := os.Args
	if len(args) < 2 {
		messageCallback(genericMessageType, fmt.Sprintf("invalid parameter count.\nUsage:\n%s", getUsageAsString()))
		return
	}

	parManager := parameters.NewCommandLine()
	runCommand(parManager, messageCallback)
}

// Serve spins up a HTTP server and will listen for connection and execute migrator
func Serve() error {
	environment, err := env.New()
	if err != nil {
		return err
	}

	return server.Serve(environment.GetHTTPServerPort(), runCommand)
}

func connect(messageCallback func(int, string)) (migrator.DBMigrator, error) {
	environment, err := env.New()
	if err != nil {
		messageCallback(genericMessageType, err.Error())
		return nil, err
	}

	db, err := dbconnector.New(environment).GetConnection(dbconnector.NewDB())
	if err != nil {
		messageCallback(genericMessageType, err.Error())
		return nil, fmt.Errorf("cannot connect %w", err)
	}

	dbPrefix := environment.GetDBPrefix()
	migrationFilePath := environment.GetMigrationPath()

	m := migrator.New(db, migrationFilePath, dbPrefix)
	m.SubscribeToMessages(messageCallback)
	return m, nil
}

func runCommand(parManager parameters.ParameterManager, messageCallback func(int, string)) {
	command := parManager.Command()
	args := parManager.Params()
	commands := map[string]func(){
		"migrate": func() {
			m, err := connect(messageCallback)
			if err == nil {
				validation := m.ChecksumValidation()
				if len(validation) > 0 && !slices.Contains(args, parameters.ForceParam) {
					messageCallback(genericMessageType, fmt.Sprintf("validation error: %s", strings.Join(validation, ", ")))
					return
				}

				count := parseMigrationCount(args)
				handleError(m.Migrate(count), messageCallback)
				return
			}
			handleError(err, messageCallback)
		},
		"rollback": func() {
			m, err := connect(messageCallback)
			if err == nil {
				count := parseMigrationCount(args)
				handleError(m.Rollback(count), messageCallback)
				return
			}
			handleError(err, messageCallback)
		},
		"refresh": func() {
			m, err := connect(messageCallback)
			if err == nil {
				handleError(m.Refresh(), messageCallback)
				return
			}
			handleError(err, messageCallback)
		},
		"report": func() {
			m, err := connect(messageCallback)
			if err == nil {
				report, err := m.Report()
				if err != nil {
					messageCallback(genericMessageType, err.Error())
				} else {
					messageCallback(genericMessageType, report)
				}
				return
			}
			handleError(err, messageCallback)
		},
		"add": func() {
			if len(args) == 0 {
				messageCallback(genericMessageType, "missing migration filename")
				return
			}
			m, err := connect(messageCallback)
			if err == nil {
				handleError(m.AddNewMigrationFiles(args[0]), messageCallback)
				return
			}
			handleError(err, messageCallback)
		},
		"validate": func() {
			m, err := connect(messageCallback)
			if err == nil {
				result := m.ChecksumValidation()
				messageCallback(genericMessageType, strings.Join(result, "\n"))
				return
			}
			handleError(err, messageCallback)
		},
		"save-baseline": func() {
			m, err := connect(messageCallback)
			if err == nil {
				handleError(m.SaveBaseline(), messageCallback)
				return
			}
			handleError(err, messageCallback)
		},
		"restore-baseline": func() {
			m, err := connect(messageCallback)
			if err == nil {
				handleError(m.LoadBaseline(), messageCallback)
				return
			}
			handleError(err, messageCallback)
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

func handleError(err error, messageCallback func(int, string)) {
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
