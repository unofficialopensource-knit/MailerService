package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unofficialopensource-knit/MailerService/pkg/config"
)

func TestConfigPasses(t *testing.T) {
	os.Setenv("MAILER_ENVIRONMENT", "test")
	t.Run("Check if config is set from env vars", func(t *testing.T) {
		conf := config.Config()
		assert.Equal(t, conf.Environment, "test")
	})
}
