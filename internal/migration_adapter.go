package migrator

import (
	"database/sql"

	migrator "github.com/olbrichattila/godbmigrator"
)

type migrationAdapter struct {
}

func newMigrationAdapter() *migrationAdapter {
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

func (a *migrationAdapter) AddNewMigrationFiles(migrationPath string, customText string) error {
	return migrator.AddNewMigrationFiles(migrationPath, customText)
}

func (a *migrationAdapter) Report(db *sql.DB, provider migrator.MigrationProvider, migrationPath string) (string, error) {
	return migrator.Report(db, provider, migrationPath)
}

func (a *migrationAdapter) ChecksumValidation(db *sql.DB, provider migrator.MigrationProvider, migrationPath string) ([]string) {
	return migrator.ChecksumValidation(db, provider, migrationPath)
}
