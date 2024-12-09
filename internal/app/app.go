package app

import (
	"context"
	"runtime"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"

	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	"github.com/Karzoug/meower-timeline-service/internal/config"
	healthHandler "github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/handler/health"
	timelineHandler "github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/handler/timeline"
	grpcServer "github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/server"
	"github.com/Karzoug/meower-timeline-service/internal/delivery/kafka"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service"
	"github.com/Karzoug/meower-timeline-service/pkg/buildinfo"
)

const (
	serviceName     = "TimelineService"
	metricNamespace = "timeline_service"
	pkgName         = "github.com/Karzoug/meower-timeline-service"
	initTimeout     = 10 * time.Second
	shutdownTimeout = 10 * time.Second
)

var serviceVersion = buildinfo.Get().ServiceVersion

func Run(ctx context.Context, logger zerolog.Logger) error {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(cfg.LogLevel)

	logger.Info().
		Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
		Str("log level", cfg.LogLevel.String()).
		Msg("starting up")

	// set timeout for initialization
	ctxInit, closeCtx := context.WithTimeout(ctx, initTimeout)
	defer closeCtx()

	// set up tracer
	cfg.OTLP.ServiceName = serviceName
	cfg.OTLP.ServiceVersion = serviceVersion
	cfg.OTLP.ExcludedGrpcMethods = map[string]string{
		"grpc.health.v1.Health": "Check",
	}
	shutdownTracer, err := otlp.RegisterGlobal(ctxInit, cfg.OTLP)
	if err != nil {
		return err
	}
	defer doClose(shutdownTracer, logger)

	tracer := otel.GetTracerProvider().Tracer(pkgName)

	// set up meter
	shutdownMeter, err := prom.RegisterGlobal(ctxInit, serviceName, serviceVersion, metricNamespace)
	if err != nil {
		return err
	}
	defer doClose(shutdownMeter, logger)

	// set up service
	ts, err := service.NewTimelineService(cfg.Service, nil, nil, nil, ctx.Done(), tracer, logger)
	if err != nil {
		return err
	}

	// set up kafka producer
	kafkaConsumer, err := kafka.NewConsumer(ctxInit, cfg.ConsumerKafka, ts, tracer, logger)
	if err != nil {
		return err
	}

	// set up grpc server
	grpcSrv := grpcServer.New(
		cfg.GRPC,
		[]grpcServer.ServiceRegister{
			healthHandler.RegisterService(),
			timelineHandler.RegisterService(ts),
		},
		tracer,
		logger,
	)

	eg, ctx := errgroup.WithContext(ctx)
	// run service grpc server
	eg.Go(func() error {
		return grpcSrv.Run(ctx)
	})
	// run kafka consumer
	eg.Go(func() error {
		return kafkaConsumer.Run(ctx)
	})
	// run prometheus metrics http server
	eg.Go(func() error {
		return prom.Serve(ctx, cfg.PromHTTP, logger)
	})

	return eg.Wait()
}

func doClose(fn func(context.Context) error, logger zerolog.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := fn(ctx); err != nil {
		logger.Error().
			Err(err).
			Msg("error closing")
	}
}
