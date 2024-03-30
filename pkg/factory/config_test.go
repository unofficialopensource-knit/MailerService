package factory_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unofficialopensource-knit/MailerService/pkg/factory"
)

func TestConfigPasses(t *testing.T) {
	os.Setenv("MAILER_ENVIRONMENT", "test")
	conf := factory.Config()

	assert.Equal(t, conf.Environment, "test")
}
