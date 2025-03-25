// Package messagedecorator formats a message
package messagedecorator

import (
	"fmt"

	"github.com/olbrichattila/godbmigrator/config"
)

// New message decorator
func New() MessageDecorator {
	return &decorator{}
}

// MessageDecorator is the interface to decorate messages
type MessageDecorator interface {
	DecorateMessage(eventType int, message string) string
}

type decorator struct {
}

// DecorateMessage will return a message with the correct context from the event type and message
func (d *decorator) DecorateMessage(eventType int, message string) string {
	if formattedMsg, exists := d.getMessageFormat(eventType); exists {
		return fmt.Sprintf(formattedMsg, message)
	}

	return message
}

// getMessageFormat returns the message format string based on the event type.
func (*decorator) getMessageFormat(eventType int) (string, bool) {
	messages := map[int]string{
		config.MigratedItems:        "Migrated %s items",
		config.NothingToRollback:    "Nothing to roll back%s",
		config.RolledBack:           "Rolled back %s items",
		config.RunningMigrations:    "Running migration: %s",
		config.SkipRollback:         "Skip rollback as file '%s' not exists",
		config.RunningRollback:      "Running rollback %s",
		config.MigrationFileCreated: "Migration file created: %s",
	}

	format, exists := messages[eventType]
	return format, exists
}
