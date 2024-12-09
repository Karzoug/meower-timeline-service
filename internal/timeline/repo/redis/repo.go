package redis

import (
	"github.com/rs/xid"
	"github.com/rs/zerolog"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	"github.com/Karzoug/meower-timeline-service/pkg/redis"
)

var emptyPost = entity.Post{
	AuthorID: xid.NilID(),
	PostID:   xid.NilID(),
}

type repo struct {
	db     redis.DB
	logger zerolog.Logger
}

func NewTimelineRepo(db redis.DB, logger zerolog.Logger) repo {
	logger = logger.With().
		Str("component", "redis repo").
		Logger()

	return repo{
		db:     db,
		logger: logger,
	}
}
