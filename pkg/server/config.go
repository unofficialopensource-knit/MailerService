package server

type Config struct {
	AWSRegion   string `envconfig:"aws_region"`
	Environment string `encvonfig:"environment"`
	BindAddress string `envconfig:"bind_addr"`
}
