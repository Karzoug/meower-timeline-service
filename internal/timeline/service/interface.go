package service

import (
	"context"
	"time"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
)

type repo interface {
	// ListGet returns timeline list from cache.
	ListGet(ctx context.Context, userID xid.ID, offset, limit int, ttl *time.Duration) ([]entity.Post, error)
	// ExistedListPush push timeline record to existed timeline list or do nothing if timeline list does not exist.
	ExistedListPushPost(ctx context.Context, userID xid.ID, post entity.Post, limit int64) error
	// ListSet set timeline list to cache (new records must be at the end).
	ListSet(ctx context.Context, userID xid.ID, posts []entity.Post, ttl time.Duration) error
	ExistedListDeletePost(ctx context.Context, userID xid.ID, post entity.Post) error
	// ExistedListDelete romoves timeline list by userID or do nothing if timeline list does not exist.
	ExistedListDelete(ctx context.Context, userID xid.ID) error
}

type relationService interface {
	ListFollowerIDs(ctx context.Context, userID xid.ID) ([]xid.ID, error)
	ListNotMutedFollowingIDs(ctx context.Context, userID xid.ID) ([]xid.ID, error)
}

type postService interface {
	ListPostIDsByUserIDs(ctx context.Context, reqUserID xid.ID, userIDs []xid.ID, limit int) ([]entity.Post, error)
}
