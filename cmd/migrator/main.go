// Package main initiate db migrator command
package main

import (
	"fmt"

	"github.com/olbrichattila/godbmigrator/messager"
	migrator "github.com/olbrichattila/godbmigrator_cmd/internal"
)

func main() {
	migrator.Init(func(et int, msg string) {
		switch et {
		case messager.MigratedItems:
			fmt.Printf("Migrated %s items\n", msg)
		case messager.NothingToRollback:
			fmt.Println("Nothing to roll back")
		case messager.RolledBack:
			fmt.Printf("Rolled back %s items\n", msg)
		case messager.RunningMigrations:
			fmt.Printf("Running migration: %s\n", msg)
		case messager.SkipRollback:
			fmt.Printf("Skip rollback as file '%s' not exists\n", msg)
		case messager.RunningRollback:
			fmt.Printf("Running rollback %s\n", msg)
		case messager.MigrationFileCreated:
			fmt.Printf("Migration file created: %s\n", msg)
		default:
			fmt.Println(msg)
		}

	})
}
