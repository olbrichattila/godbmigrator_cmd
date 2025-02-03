package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type envTestSuite struct {
	suite.Suite
}

func TestEnvRunner(t *testing.T) {
	suite.Run(t, new(envTestSuite))
}

func (t *envTestSuite) SetupTest() {
	os.Setenv("DB_CONNECTION", "conection")
	os.Setenv("DB_HOST", "host")
	os.Setenv("DB_PORT", "1234")
	os.Setenv("DB_DATABASE", "db")
	os.Setenv("DB_USERNAME", "usenName")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_SSLMODE", "verify-ca")
}

func (t *envTestSuite) TestEnvReturnsCorrectValues() {
	env, err := New()
	t.NoError(err)

	t.Equal("conection", env.GetDBConnection())
	t.Equal("usenName", env.GetDBUserName())
	t.Equal("password", env.GetDBPassword())
	t.Equal("host", env.GetDBHost())
	t.Equal(1234, env.GetDBPort())
	t.Equal("db", env.GetDBDatabase())
	t.Equal("verify-ca", env.GetDBPostgresSSLMode())
}
