package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	Environment               string `encvonfig:"environment"`
	BindAddress               string `envconfig:"bind_addr"`
	SMTPIdentity              string `envconfig:"smtp_identity"`
	SMTPUsername              string `envconfig:"smtp_username"`
	SMTPPassword              string `envconfig:"smtp_password"`
	SMTPHost                  string `envconfig:"smtp_host"`
	SMTPPort                  string `envconfig:"smtp_port"`
	ContactUsDefaultRecipient string `envconfig:"contact_us_default_recipient"`
}

func Config() Settings {
	var conf Settings
	envconfig.MustProcess("MAILER", &conf)
	return conf
}
