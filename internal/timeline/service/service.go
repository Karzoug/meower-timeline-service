package service

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/singleflight"
)

const defaultTimeout = 5 * time.Second

type TimelineService struct {
	repo repo
	relationService
	postService
	cfg         Config
	shutdownCtx context.Context // for background workers
	syncGroup   *singleflight.Group
	tracer      trace.Tracer
	logger      zerolog.Logger
}

func NewTimelineService(cfg Config,
	repo repo,
	relationService relationService,
	postService postService,
	closeChan <-chan struct{},
	tracer trace.Tracer,
	logger zerolog.Logger,
) (TimelineService, error) {
	logger = logger.With().
		Str("component", "timeline service").
		Logger()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-closeChan
		cancel()
	}()

	return TimelineService{
		cfg:             cfg,
		repo:            repo,
		relationService: relationService,
		postService:     postService,
		syncGroup:       new(singleflight.Group),
		shutdownCtx:     ctx,
		tracer:          tracer,
		logger:          logger,
	}, nil
}
