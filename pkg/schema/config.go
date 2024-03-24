package schema

type Config struct {
	Environment string `encvonfig:"environment"`
	BindAddress string `envconfig:"bind_addr"`
}
