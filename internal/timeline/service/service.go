package service

import (
	"context"

	"github.com/rs/zerolog"
	"golang.org/x/sync/singleflight"
)

type TimelineService struct {
	repo repo
	relationService
	postService
	cfg         Config
	shutdownCtx context.Context // for background workers
	syncGroup   *singleflight.Group
	logger      zerolog.Logger
}

func NewTimelineService(cfg Config,
	repo repo,
	relationService relationService,
	postService postService,
	closeChan <-chan struct{},
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
		logger:          logger,
	}, nil
}
