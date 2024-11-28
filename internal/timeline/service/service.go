package service

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/sync/singleflight"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc/post"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc/relation"
)

const defaultTimeout = 5 * time.Second

type TimelineService struct {
	repo           repo
	cfg            Config
	postClient     post.Client
	relationClient relation.Client
	eventProducer  eventProducer
	shutdownCtx    context.Context
	syncGroup      *singleflight.Group
	logger         zerolog.Logger
}

func NewTimelineService(cfg Config,
	repo repo,
	pclient post.Client,
	rclient relation.Client,
	eventProducer eventProducer,
	shutdownSig <-chan struct{},
	logger zerolog.Logger,
) (TimelineService, error) {
	logger = logger.With().
		Str("component", "user service").
		Logger()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-shutdownSig
		cancel()
	}()

	return TimelineService{
		repo:           repo,
		cfg:            cfg,
		postClient:     pclient,
		relationClient: rclient,
		shutdownCtx:    ctx,
		syncGroup:      new(singleflight.Group),
		logger:         logger,
	}, nil
}
