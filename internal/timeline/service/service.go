package service

import (
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

const defaultTimeout = 5 * time.Second

type TimelineService struct {
	tracer trace.Tracer
	logger zerolog.Logger
}

func NewTimelineService(
	tracer trace.Tracer,
	logger zerolog.Logger,
) (TimelineService, error) {
	logger = logger.With().
		Str("component", "user service").
		Logger()

	return TimelineService{
		tracer: tracer,
		logger: logger,
	}, nil
}
