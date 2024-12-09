package redis

type Config struct {
	Addrs         []string `env:"ADDRS,notEmpty" envSeparator:","`
	EnableTracing bool     `env:"ENABLE_TRACING" envDefault:"false"`
}
