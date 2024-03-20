package server

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment string `encvonfig:"environment"`
	BindAddress string `envconfig:"bind_addr"`
}

func LoadConfig() Config {
	var config Config

	envconfig.MustProcess("mailer", &config)

	return config
}
