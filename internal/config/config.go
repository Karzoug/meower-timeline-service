package config

import (
	grpcConfig "github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/server"
	inkafka "github.com/Karzoug/meower-timeline-service/internal/delivery/kafka"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	outkafka "github.com/Karzoug/meower-timeline-service/internal/timeline/client/kafka"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service"
	"github.com/Karzoug/meower-timeline-service/pkg/metric/prom"
	"github.com/Karzoug/meower-timeline-service/pkg/trace/otlp"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel        zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	GRPC            grpcConfig.Config `envPrefix:"GRPC_"`
	PromHTTP        prom.ServerConfig `envPrefix:"PROM_"`
	OTLP            otlp.Config       `envPrefix:"OTLP_"`
	ProducerKafka   outkafka.Config   `envPrefix:"PRODUCER_KAFKA_"`
	ConsumerKafka   inkafka.Config    `envPrefix:"CONSUMER_KAFKA_"`
	Service         service.Config    `envPrefix:"SERVICE_"`
	PostService     grpc.Config       `envPrefix:"POST_SERVICE"`
	RelationService grpc.Config       `envPrefix:"RELATION_SERVICE"`
}
