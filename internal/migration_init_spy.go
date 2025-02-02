package migrator

import (
	"database/sql"
	"strconv"

	migrator "github.com/olbrichattila/godbmigrator"
)

type migrationInitSpy struct {
}

func newMigrationInitSpy() *migrationInitSpy {
	return &migrationInitSpy{}
}

func (m *migrationInitSpy) migrationInit(args []string, initMigrationTables bool) (*sql.DB, migrator.MigrationProvider, int, error) {
	conn, err := sql.Open(driverTypeSqLite, ":memory:")
	if err != nil {
		return nil, nil, 0, err
	}

	count := 0
	if len(args) > 2 {
		count, err = strconv.Atoi(args[2])
		if err != nil {
			return nil, nil, 0, err
		}
	}
	provider := newMigrationSpyProvider()

	return conn, provider, count, err
}
