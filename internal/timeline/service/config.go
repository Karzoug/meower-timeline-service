package service

import "time"

type Config struct {
	Timeline struct {
		Limit int64         `env:"LIMIT,notEmpty" envDefault:"100"`
		TTL   time.Duration `env:"TTL,notEmpty" envDefault:"72h"`
	} `envPrefix:"TIMELINE_"`
}
