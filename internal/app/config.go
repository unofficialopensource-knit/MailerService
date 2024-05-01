package app

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type HTTPConfig struct {
	LambdaTaskRoot            string `envconfig:"LAMBDA_TASK_ROOT"`
	BindAddress               string `envconfig:"BIND_ADDR"`
	SMTPIdentity              string `envconfig:"SMTP_IDENTITY, required"`
	SMTPUsername              string `envconfig:"SMTP_USERNAME, required"`
	SMTPPassword              string `envconfig:"SMTP_PASSWORD, required"`
	SMTPHost                  string `envconfig:"SMTP_HOST, required"`
	SMTPPort                  string `envconfig:"SMTP_PORT, required"`
	ContactUsDefaultRecipient string `envconfig:"CONTACT_US_DEFAULT_RECIPIENT, required"`
}

func NewHTTPConfig(c context.Context) (HTTPConfig, error) {
	var conf HTTPConfig
	err := envconfig.Process(c, &conf)
	return conf, err
}
