package factory

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func Config() schema.Config {
	var conf schema.Config
	envconfig.MustProcess("MAILER", &conf)
	return conf
}
