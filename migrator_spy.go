package main

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
	lastAddCustomText string
}

func newSpyMigrator() *spyMigrator {
	return &spyMigrator{}
}

func (a *spyMigrator) Migrate(db *sql.DB, provider migrator.MigrationProvider, migrationPath string, count int) error {
	a.migrateCalled++
	a.lastCount = count
	return nil
}

func (a *spyMigrator) Rollback(db *sql.DB, provider migrator.MigrationProvider, migrationPath string, count int) error {
	a.rollbackCalled++
	a.lastCount = count
	return nil
}

func (a *spyMigrator) Refresh(db *sql.DB, provider migrator.MigrationProvider, migrationPath string) error {
	a.refreshCalled++
	return nil
}

func (a *spyMigrator) AddNewMigrationFiles(migrationPath string, customText string) {
	a.lastAddCustomText = customText
	a.addCalled++
}
