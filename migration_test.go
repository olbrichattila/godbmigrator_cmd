package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) TestRunning() {
	t.True(true)
}

func (t *TestSuite) TestCommandLineMigrateCalled() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "migrate"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(1, spyMigrator.migrateCalled)

	t.Equal(0, spyMigrator.rollbackCalled)

	t.Equal(0, spyMigrator.refreshCalled)

	t.Equal(0, spyMigrator.addCalled)
}

func (t *TestSuite) TestCommandLineRollbackCalled() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "rollback"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(0, spyMigrator.migrateCalled)

	t.Equal(1, spyMigrator.rollbackCalled)

	t.Equal(0, spyMigrator.refreshCalled)

	t.Equal(0, spyMigrator.addCalled)
}

func (t *TestSuite) TestCommandLineRefreshCalled() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "refresh"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(0, spyMigrator.migrateCalled)

	t.Equal(0, spyMigrator.rollbackCalled)

	t.Equal(1, spyMigrator.refreshCalled)

	t.Equal(0, spyMigrator.addCalled)
}

func (t *TestSuite) TestCommandLineAddCalled() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "add"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(0, spyMigrator.migrateCalled)

	t.Equal(0, spyMigrator.rollbackCalled)

	t.Equal(0, spyMigrator.refreshCalled)

	t.Equal(1, spyMigrator.addCalled)
}

func (t *TestSuite) TestCommandLineMigrateCalledWithCount3() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "migrate", "3"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(1, spyMigrator.migrateCalled)

	t.Equal(3, spyMigrator.lastCount)
}

func (t *TestSuite) TestCommandLineRollbackCalledWithCount3() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "rollback", "3"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(1, spyMigrator.rollbackCalled)

	t.Equal(3, spyMigrator.lastCount)
}

func (t *TestSuite) TestCommandLineRefreshCalledWithCount3ButCalledWith0() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "refresh", "3"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)
	t.Nil(err)

	t.Equal(1, spyMigrator.refreshCalled)

	t.Equal(0, spyMigrator.lastCount)
}

func (t *TestSuite) TestCommandLineCalledWithIncorrectCommandThrowsError() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd", "incorrect"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)

	t.NotNil(err)
}

func (t *TestSuite) TestCommandLineCalledWithIncorrectNumberOfParametersThrowsError() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"cmd"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)

	t.NotNil(err)

	args = []string{}

	err = routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)

	t.NotNil(err)
}

func (t *TestSuite) TestCommandLineCalledWithAddNoCustomText() {
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"command", "add"}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)

	t.Nil(err)

	t.Equal("", spyMigrator.lastAddCustomText)
}

func (t *TestSuite) TestCommandLineCalledWithAddCustomText() {
	expectedText := "custom-example-text"
	spyMigrator := newSpyMigrator()
	newMigrationInitSpy := newMigrationInitSpy()
	args := []string{"command", "add", expectedText}

	err := routeCommandLineParameters(args, spyMigrator, newMigrationInitSpy)

	t.Nil(err)

	// Assert hyphen is added
	t.Equal("-"+expectedText, spyMigrator.lastAddCustomText)
}
