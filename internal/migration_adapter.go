package migrator

import (
	"database/sql"

	migrator "github.com/olbrichattila/godbmigrator"
	"github.com/olbrichattila/godbmigrator/messager"
)

type migrationAdapter struct {
}

func newMigrationAdapter() *migrationAdapter {
	return &migrationAdapter{}
}

func (a *migrationAdapter) subscribeToMessages(callback messager.CallbackFunc) {
	migrator.SubscribeToMessages(callback)
}

func (a *migrationAdapter) Migrate(db *sql.DB, provider string, migrationPath string, count int) error {
	return migrator.Migrate(db, provider, migrationPath, count)
}

func (a *migrationAdapter) Rollback(db *sql.DB, provider string, migrationPath string, count int) error {
	return migrator.Rollback(db, provider, migrationPath, count)
}

func (a *migrationAdapter) Refresh(db *sql.DB, provider string, migrationPath string) error {
	return migrator.Refresh(db, provider, migrationPath)
}

func (a *migrationAdapter) AddNewMigrationFiles(migrationPath string, customText string) error {
	return migrator.AddNewMigrationFiles(migrationPath, customText)
}

func (a *migrationAdapter) Report(db *sql.DB, provider string, migrationPath string) (string, error) {
	return migrator.Report(db, provider, migrationPath)
}

func (a *migrationAdapter) ChecksumValidation(db *sql.DB, provider string, migrationPath string) []string {
	return migrator.ChecksumValidation(db, provider, migrationPath)
}

func (a *migrationAdapter) SaveBaseline(db *sql.DB, migrationFilePath string) error {
	return migrator.SaveBaseline(db, migrationFilePath)
}

func (a *migrationAdapter) LoadBaseline(db *sql.DB, migrationFilePath string) error {
	return migrator.LoadBaseline(db, migrationFilePath)
}
