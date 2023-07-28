package main

import (
	"database/sql"

	migrator "github.com/olbrichattila/godbmigrator"
)

type migrationAdapter struct {
}

func NewMigrationAdapter() *migrationAdapter {
	return &migrationAdapter{}
}

func (a *migrationAdapter) Migrate(db *sql.DB, provider migrator.MigrationProvider, migrationPath string, count int) error {
	return migrator.Migrate(db, provider, migrationPath, count)
}

func (a *migrationAdapter) Rollback(db *sql.DB, provider migrator.MigrationProvider, migrationPath string, count int) error {
	return migrator.Rollback(db, provider, migrationPath, count)
}

func (a *migrationAdapter) Refresh(db *sql.DB, provider migrator.MigrationProvider, migrationPath string) error {
	return migrator.Refresh(db, provider, migrationPath)
}

func (a *migrationAdapter) AddNewMigrationFiles(migrationPath string, customText string) {
	migrator.AddNewMigrationFiles(migrationPath, customText)
}
