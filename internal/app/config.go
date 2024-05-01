package app

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type HTTPConfig struct {
	LambdaTaskRoot            string `envconfig:"LAMBDA_TASK_ROOT"`
	Environment               string `encvonfig:"ENVIRONMENT"`
	BindAddress               string `envconfig:"BIND_ADDR"`
	SMTPIdentity              string `envconfig:"SMTP_IDENTITY"`
	SMTPUsername              string `envconfig:"SMTP_USERNAME"`
	SMTPPassword              string `envconfig:"SMTP_PASSWORD"`
	SMTPHost                  string `envconfig:"SMTP_HOST"`
	SMTPPort                  string `envconfig:"SMTP_PORT"`
	ContactUsDefaultRecipient string `envconfig:"CONTACT_US_DEFAULT_RECIPIENT"`
}

func LoadConfig(c context.Context) (HTTPConfig, error) {
	var conf HTTPConfig
	err := envconfig.Process(c, &conf)
	return conf, err
}
