package migrator

import (
	"database/sql"
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

func (a *spyMigrator) Migrate(_ *sql.DB, _ string, _ string, count int) error {
	a.migrateCalled++
	a.lastCount = count
	return nil
}

func (a *spyMigrator) Rollback(_ *sql.DB, _ string, _ string, count int) error {
	a.rollbackCalled++
	a.lastCount = count
	return nil
}

func (a *spyMigrator) Refresh(_ *sql.DB, _ string, _ string) error {
	a.refreshCalled++
	return nil
}

func (a *spyMigrator) AddNewMigrationFiles(_ string, customText string) error {
	a.lastAddCustomText = customText
	a.addCalled++
	return nil
}

func (a *spyMigrator) Report(*sql.DB, string, string) (string, error) {
	a.reportCallCount++

	return "report result", nil
}

func (a *spyMigrator) ChecksumValidation(*sql.DB, string, string) []string {
	return []string{}
}

func (a *spyMigrator) SaveBaseline(*sql.DB, string) error {
	return nil
}

func (a *spyMigrator) LoadBaseline(*sql.DB, string) error {
	return nil
}
