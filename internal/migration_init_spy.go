package migrator

import (
	"database/sql"
	"strconv"
)

type migrationInitSpy struct {
}

func newMigrationInitSpy() *migrationInitSpy {
	return &migrationInitSpy{}
}

func (m *migrationInitSpy) migrationInit(args []string, _ bool) (*sql.DB, string, int, error) {
	conn, err := sql.Open(driverTypeSqLite, ":memory:")
	if err != nil {
		return nil, "", 0, err
	}

	count := 0
	if len(args) > 2 {
		count, err = strconv.Atoi(args[2])
		if err != nil {
			return nil, "", 0, err
		}
	}

	return conn, "", count, err
}
