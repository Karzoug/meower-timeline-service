package service

import (
	"context"
	"time"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
)

type repo interface {
	// ListGet returns timeline list from cache.
	ListGet(ctx context.Context, userID string, offset, limit int, ttl time.Duration) ([]entity.Post, error)
	// ExistedListPush push timeline record to existed timeline list or do nothing if timeline list does not exist.
	ExistedListPush(ctx context.Context, userID string, record entity.Post, limit int64) error
	// ListSet set timeline list to cache.
	ListSet(ctx context.Context, userID string, records []entity.Post, ttl time.Duration) error
}

type eventProducer interface {
	SendFollowersGotNewPostEvent(followerIDs []string, authorID, postID string) error
}
