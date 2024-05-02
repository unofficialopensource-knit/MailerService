package app_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigTestSuite(t *testing.T) {
	t.Log("Running suite ConfigTestSuite")

	suite.Run(t, new(ConfigTestSuite))

	t.Log("ConfigTestSuite completed")
}

func (s *ConfigTestSuite) SetupSuite() {
	s.T().Log("Running Set up for ConfigTestSuite")

	s.T().Log("Value of environment variable SMTP_HOST=" + os.Getenv("SMTP_HOST"))

	s.T().Log("Set up method for ConfigTestSuite completed")
}

func (s *ConfigTestSuite) TearDownSuite() {
	s.T().Log("Running Tear Down for ConfigTestSuite")

	s.T().Log("Value of environment variable SMTP_HOST=" + os.Getenv("SMTP_HOST"))

	s.T().Log("Tear Down method for ConfigTestSuite completed")
}

func (s *ConfigTestSuite) TestConfigSetCorrectly() {
	s.T().Log("Checks if config is correctly set from env vars")
}

func (s *ConfigTestSuite) TestConfigMissing() {
	s.T().Log("Checks if config fails due to missing env var")
}
