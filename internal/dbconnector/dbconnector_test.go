package dbconnector

import (
	"os"
	"testing"

	"github.com/olbrichattila/godbmigrator_cmd/internal/env"
	"github.com/stretchr/testify/suite"
)

type dbConnectionTestSuite struct {
	suite.Suite
}

func TestEnvRunner(t *testing.T) {
	suite.Run(t, new(dbConnectionTestSuite))
}

func (t *dbConnectionTestSuite) TestIncorrectConnection() {
	os.Setenv("DB_CONNECTION", "dbConection")
	os.Setenv("DB_DATABASE", "database.sqlite")

	env, _ := env.New()
	connector := New(env)

	dbMock := NewDBMock()
	_, err := connector.GetConnection(dbMock)
	t.Error(err)
}

func (t *dbConnectionTestSuite) TestSqLiteConnection() {
	os.Setenv("DB_CONNECTION", "sqlite")
	os.Setenv("DB_DATABASE", "database.sqlite")

	env, _ := env.New()
	connector := New(env)

	dbMock := NewDBMock()
	_, err := connector.GetConnection(dbMock)
	t.NoError(err)

	t.Equal(1, dbMock.callCnt)
	t.Equal("sqlite3", dbMock.calledWithDBengine)
	t.Equal("database.sqlite", dbMock.calledWithConnectionStr)
}
