package redis

import (
	"context"
	"time"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
)

func (r repo) ExistedListPushPost(ctx context.Context, userID xid.ID, record entity.Post, limit int64) error {
	pipe := r.db.Pipeline()

	pipe.LPushX(ctx, userID.String(), record)
	pipe.LTrim(ctx, userID.String(), 0, limit+1)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r repo) ListSet(ctx context.Context, userID xid.ID, records []entity.Post, ttl time.Duration) error {
	pipe := r.db.Pipeline()

	pipe.LTrim(ctx, userID.String(), 1, 0)
	pipe.LPush(ctx, userID.String(), emptyPost)

	// TODO: perf: rewrite it for batch insert
	for i := range records {
		pipe.LPush(ctx, userID.String(), records[i])
	}
	pipe.Expire(ctx, userID.String(), ttl)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r repo) ExistedListDeletePost(ctx context.Context, userID xid.ID, post entity.Post) error {
	cmd := r.db.LRem(ctx, userID.String(), -1, post)
	return cmd.Err()
}

func (r repo) ExistedListDelete(ctx context.Context, userID xid.ID) error {
	return r.db.Del(ctx, userID.String()).Err()
}
