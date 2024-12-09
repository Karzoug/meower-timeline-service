package service

import "time"

type Config struct {
	// Limit of timeline records in cache.
	Limit int `env:"LIMIT,notEmpty" envDefault:"1000"`
	// TTL of timeline records in cache.
	TTL time.Duration `env:"TTL,notEmpty" envDefault:"72h"`
	// BuildTimeout  is timeout for build timeline from scratch.
	BuildTimeout time.Duration `env:"BUILD_TIMEOUT,notEmpty" envDefault:"180s"`
}
