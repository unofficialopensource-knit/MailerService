package app

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type HTTPConfig struct {
	LambdaTaskRoot            string `env:"LAMBDA_TASK_ROOT"`
	BindAddress               string `env:"BIND_ADDR"`
	SMTPIdentity              string `env:"SMTP_IDENTITY, required"`
	SMTPUsername              string `env:"SMTP_USERNAME, required"`
	SMTPPassword              string `env:"SMTP_PASSWORD, required"`
	SMTPHost                  string `env:"SMTP_HOST, required"`
	SMTPPort                  string `env:"SMTP_PORT, required"`
	ContactUsDefaultRecipient string `env:"CONTACT_US_DEFAULT_RECIPIENT, required"`
}

func NewHTTPConfig(c context.Context) (HTTPConfig, error) {
	var conf HTTPConfig
	err := envconfig.Process(c, &conf)
	return conf, err
}
