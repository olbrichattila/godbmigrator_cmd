package main

type migrationProviderSpy struct {
	calledMigrations             int
	calledAddToMigrations        int
	calledRemoveFromMigrations   int
	calledMigrationExistsForFile int
	calledResetDate              int
	calledGetJsonFileName        int
	calledSetJsonFileName        int
}

func newMigrationSpyProvider() *migrationProviderSpy {
	return &migrationProviderSpy{}
}

func (s migrationProviderSpy) Migrations(bool) ([]string, error) {
	s.calledMigrations++

	return []string{}, nil
}

func (s migrationProviderSpy) AddToMigration(string) error {
	s.calledAddToMigrations++

	return nil
}

func (s migrationProviderSpy) RemoveFromMigration(string) error {
	s.calledRemoveFromMigrations++

	return nil
}

func (s migrationProviderSpy) MigrationExistsForFile(string) bool {
	s.calledMigrationExistsForFile++

	return false
}

func (s migrationProviderSpy) ResetDate() {
	s.calledResetDate++
}

func (s migrationProviderSpy) GetJsonFileName() string {
	s.calledGetJsonFileName++

	return "./mockjson.json"

}

func (s migrationProviderSpy) SetJsonFileName(string) {
	s.calledSetJsonFileName++
}
