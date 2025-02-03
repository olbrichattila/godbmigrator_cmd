// Package main initiates the DB migrator command
package main

import (
	"fmt"

	"github.com/olbrichattila/godbmigrator/messager"
	migratorcommand "github.com/olbrichattila/godbmigrator_cmd/internal"
)

func main() {
	migratorcommand.Init(showMessage)
}

// showMessage formats and prints messages based on the event type.
func showMessage(eventType int, msg string) {
	if formattedMsg, exists := getMessageFormat(eventType); exists {
		fmt.Printf(formattedMsg+"\n", msg)
	} else {
		fmt.Println(msg)
	}
}

// getMessageFormat returns the message format string based on the event type.
func getMessageFormat(eventType int) (string, bool) {
	messages := map[int]string{
		messager.MigratedItems:        "Migrated %s items",
		messager.NothingToRollback:    "Nothing to roll back%s",
		messager.RolledBack:           "Rolled back %s items",
		messager.RunningMigrations:    "Running migration: %s",
		messager.SkipRollback:         "Skip rollback as file '%s' not exists",
		messager.RunningRollback:      "Running rollback %s",
		messager.MigrationFileCreated: "Migration file created: %s",
	}

	format, exists := messages[eventType]
	return format, exists
}
