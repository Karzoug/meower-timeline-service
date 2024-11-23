package config

import (
	grpcConfig "github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/server"
	"github.com/Karzoug/meower-timeline-service/internal/delivery/kafka"
	"github.com/Karzoug/meower-timeline-service/pkg/metric/prom"
	"github.com/Karzoug/meower-timeline-service/pkg/trace/otlp"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	GRPC     grpcConfig.Config `envPrefix:"GRPC_"`
	PromHTTP prom.ServerConfig `envPrefix:"PROM_"`
	OTLP     otlp.Config       `envPrefix:"OTLP_"`
	Kafka    kafka.Config      `envPrefix:"KAFKA_"`
}
