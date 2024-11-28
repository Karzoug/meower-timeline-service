package kafka

type Config struct {
	// Kafka brokers addresses separated by comma
	Brokers string `env:"BROKERS,notEmpty"`
}
