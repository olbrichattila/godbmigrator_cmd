// Package main initiates the DB migrator command
package main

import (
	"fmt"
	"os"

	migratorcommand "github.com/olbrichattila/godbmigrator_cmd/internal"
	"github.com/olbrichattila/godbmigrator_cmd/internal/messagedecorator"
)

var decorator messagedecorator.MessageDecorator

func main() {
	decorator = messagedecorator.New()
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		if err := migratorcommand.Serve(); err != nil {
			fmt.Println("Error running server: " + err.Error())
		}
		return
	}

	migratorcommand.Init(showMessage)
}

// showMessage formats and prints messages based on the event type.
func showMessage(eventType int, msg string) {
	fmt.Println(decorator.DecorateMessage(eventType, msg))
}
