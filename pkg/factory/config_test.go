package factory_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unofficialopensource-knit/MailerService/pkg/factory"
)

func TestConfigPasses(t *testing.T) {
	os.Setenv("MAILER_ENVIRONMENT", "test")
	conf, err := factory.Config()
	fmt.Println(conf)

	assert.Empty(t, err)
	assert.Equal(t, conf.Environment, "test")
}
