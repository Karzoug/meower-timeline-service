package config

import (
	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	grpcConfig "github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/server"
	"github.com/Karzoug/meower-timeline-service/internal/delivery/kafka"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service"
	"github.com/Karzoug/meower-timeline-service/pkg/redis"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel        zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	GRPC            grpcConfig.Config `envPrefix:"GRPC_"`
	PromHTTP        prom.ServerConfig `envPrefix:"PROM_"`
	OTLP            otlp.Config       `envPrefix:"OTLP_"`
	ConsumerKafka   kafka.Config      `envPrefix:"CONSUMER_KAFKA_"`
	Service         service.Config    `envPrefix:"SERVICE_"`
	Redis           redis.Config      `envPrefix:"REDIS_"`
	PostService     grpc.Config       `envPrefix:"POST_SERVICE_"`
	RelationService grpc.Config       `envPrefix:"RELATION_SERVICE_"`
}
