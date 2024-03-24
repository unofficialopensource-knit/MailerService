package schema

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Environment string `encvonfig:"environment"`
	BindAddress string `envconfig:"bind_addr"`
}

func LoadConfig() (Config, error) {
	var conf Config
	err := envconfig.Process("mailer", conf)
	return conf, err
}
