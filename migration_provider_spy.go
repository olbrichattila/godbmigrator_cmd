package main

type migrationProviderSpy struct {
}

func (s *migrationProviderSpy) SetJsonFilePath(string) {
}

func newMigrationSpyProvider() *migrationProviderSpy {
	return &migrationProviderSpy{}
}

func (s migrationProviderSpy) Migrations(bool) ([]string, error) {
	return []string{}, nil
}

func (s migrationProviderSpy) AddToMigration(string) error {
	return nil
}

func (s migrationProviderSpy) RemoveFromMigration(string) error {
	return nil
}

func (s migrationProviderSpy) MigrationExistsForFile(string) bool {
	return false
}

func (s migrationProviderSpy) ResetDate() {
}

func (s migrationProviderSpy) GetJsonFileName() string {
	return "./mockjson.json"
}

func (s migrationProviderSpy) Report() (string, error) {
	return "", nil
}

func (s *migrationProviderSpy) AddToMigrationReport(string, error) error {
	return nil
}
