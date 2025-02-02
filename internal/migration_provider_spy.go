package migrator

import migrator "github.com/olbrichattila/godbmigrator"

type migrationProviderSpy struct {
}

func (s *migrationProviderSpy) SetJSONFilePath(string) {
}

func newMigrationSpyProvider() *migrationProviderSpy {
	return &migrationProviderSpy{}
}

func (s migrationProviderSpy) Migrations(bool) ([]migrator.MigrationRow, error) {
	return []migrator.MigrationRow{}, nil
}

func (s migrationProviderSpy) AddToMigration(string, string) error {
	return nil
}

func (s migrationProviderSpy) RemoveFromMigration(string) error {
	return nil
}

func (s migrationProviderSpy) MigrationExistsForFile(string) (bool, error) {
	return false, nil
}

func (s migrationProviderSpy) ResetDate() {
}

func (s migrationProviderSpy) GetJSONFileName() string {
	return "./mockjson.json"
}

func (s migrationProviderSpy) Report() (string, error) {
	return "", nil
}

func (s *migrationProviderSpy) AddToMigrationReport(string, error) error {
	return nil
}

func (s *migrationProviderSpy) CreateMigrationTables() error {
	return nil
}
