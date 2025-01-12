package migrator

import (
	"database/sql"

	migrator "github.com/olbrichattila/godbmigrator"
)

type spyMigrator struct {
	migrateCalled     int
	rollbackCalled    int
	refreshCalled     int
	addCalled         int
	lastCount         int
	reportCallCount   int
	lastAddCustomText string
}

func newSpyMigrator() *spyMigrator {
	return &spyMigrator{}
}

func (a *spyMigrator) Migrate(_ *sql.DB, _ migrator.MigrationProvider, _ string, count int) error {
	a.migrateCalled++
	a.lastCount = count
	return nil
}

func (a *spyMigrator) Rollback(_ *sql.DB, _ migrator.MigrationProvider, _ string, count int) error {
	a.rollbackCalled++
	a.lastCount = count
	return nil
}

func (a *spyMigrator) Refresh(_ *sql.DB, _ migrator.MigrationProvider, _ string) error {
	a.refreshCalled++
	return nil
}

func (a *spyMigrator) AddNewMigrationFiles(_ string, customText string) error {
	a.lastAddCustomText = customText
	a.addCalled++
	return nil
}

func (a *spyMigrator) Report(*sql.DB, migrator.MigrationProvider, string) (string, error) {
	a.reportCallCount++

	return "report result", nil
}
func (a *spyMigrator) ChecksumValidation(*sql.DB, migrator.MigrationProvider, string) []string {
	return []string{}
}
