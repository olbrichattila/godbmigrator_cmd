package dbconnector

import "database/sql"

// NewDB creates an abstracted database implementation
func NewDB() Database {
	return &db{}
}

// Database purpose of this interface to make DB testable, this way it can be mocked
type Database interface {
	Open(string, string) (*sql.DB, error)
}

type db struct {
}

// Open will call sql.Open to open the database
func (*db) Open(dbEngine, connectionStr string) (*sql.DB, error) {
	return sql.Open(dbEngine, connectionStr)
}
