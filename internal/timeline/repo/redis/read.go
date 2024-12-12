package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	repoerr "github.com/Karzoug/meower-timeline-service/internal/timeline/repo"
)

const setExpireTimeout = 5 * time.Second

func (r repo) ListGetOlder(ctx context.Context, userID xid.ID, token *entity.Post, limit int, ttl *time.Duration) ([]entity.Post, *entity.Post, error) {
	expfn := func() {
		ctx, cancel := context.WithTimeout(context.Background(), setExpireTimeout)
		defer cancel()

		if cmd := r.db.Expire(ctx, userID.String(), *ttl); cmd.Err() != nil {
			r.logger.Error().
				Err(cmd.Err()).
				Str("user_id", userID.String()).
				Msg("failed to set expire")
			return
		}
	}

	var pos int64
	if token != nil {
		pipe := r.db.Pipeline()
		lenCmd := pipe.LLen(ctx, userID.String())
		posCmd := pipe.LPos(ctx, userID.String(), token.String(), redis.LPosArgs{})
		if _, err := pipe.Exec(ctx); err != nil {
			return nil, nil, err
		}

		if lenCmd.Val() == 0 {
			return nil, nil, repoerr.ErrKeyNotFound
		}
		if errors.Is(posCmd.Err(), redis.Nil) {
			return nil, nil, repoerr.ErrValueNotFound
		}

		pos = posCmd.Val()
	}

	// TODO: perf problems: 2 redis slow calls O(n) - LRANGE and LPOS,
	// rewrite using redis functions
	listCmd := r.db.LRange(ctx, userID.String(), pos, pos+int64(limit))
	if len(listCmd.Val()) != 0 {
		if ttl != nil {
			go expfn()
		}

		res := make([]entity.Post, 0, len(listCmd.Val()))
		if err := listCmd.ScanSlice(&res); err != nil {
			return nil, nil, err
		}

		// remove empty elements, if found -> it means we reached the end
		if v := res[len(res)-1]; v.PostID.IsZero() && v.AuthorID.IsZero() {
			res = res[:len(res)-1]
			return res, nil, nil
		}

		nextToken := res[len(res)-1]

		return res[0 : len(res)-1], &nextToken, nil
	}

	return []entity.Post{}, nil, repoerr.ErrKeyNotFound
}

func (r repo) ListGetNewer(ctx context.Context, userID xid.ID, token entity.Post, limit int, ttl *time.Duration) ([]entity.Post, *entity.Post, error) {
	expfn := func() {
		ctx, cancel := context.WithTimeout(context.Background(), setExpireTimeout)
		defer cancel()

		if cmd := r.db.Expire(ctx, userID.String(), *ttl); cmd.Err() != nil {
			r.logger.Error().
				Err(cmd.Err()).
				Str("user_id", userID.String()).
				Msg("failed to set expire")
			return
		}
	}

	pipe := r.db.Pipeline()

	lenCmd := pipe.LLen(ctx, userID.String())
	posCmd := pipe.LPos(ctx, userID.String(), token.String(), redis.LPosArgs{})

	if _, err := pipe.Exec(ctx); err != nil {
		return nil, nil, err
	}

	if lenCmd.Val() == 0 {
		return nil, nil, repoerr.ErrKeyNotFound
	}

	if errors.Is(posCmd.Err(), redis.Nil) {
		return nil, nil, repoerr.ErrValueNotFound
	}

	// TODO: perf problems: 2 redis slow calls O(n) - LRANGE and LPOS,
	// rewrite using redis functions
	listCmd := r.db.LRange(ctx, userID.String(), min(0, posCmd.Val()-int64(limit)), min(0, posCmd.Val()-1))
	if len(listCmd.Val()) != 0 {
		if ttl != nil {
			go expfn()
		}
		res := make([]entity.Post, 0, len(listCmd.Val()))
		if err := listCmd.ScanSlice(&res); err != nil {
			return nil, nil, err
		}

		return res, &token, nil
	}

	return []entity.Post{}, nil, repoerr.ErrKeyNotFound
}
