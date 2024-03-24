package factory

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func Config() (schema.Config, error) {
	var conf schema.Config
	err := envconfig.Process("mailer", &conf)
	return conf, err
}
