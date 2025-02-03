package dbconnector

import "database/sql"

// NewDBMock creates a mock sql
func NewDBMock() *DbMock {
	return &DbMock{}
}

// DbMock to test if db open was called and which db engine, connection string was passed
type DbMock struct {
	callCnt                 int
	calledWithDBengine      string
	calledWithConnectionStr string
}

// Open will call sql.Open to open the database
func (d *DbMock) Open(dbEngine, connectionStr string) (*sql.DB, error) {
	d.callCnt++
	d.calledWithDBengine = dbEngine
	d.calledWithConnectionStr = connectionStr

	return &sql.DB{}, nil
}
