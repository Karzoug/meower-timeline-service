package redis

import (
	"context"
	"time"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	repoerr "github.com/Karzoug/meower-timeline-service/internal/timeline/repo"
)

const setExpireTimeout = 5 * time.Second

func (r repo) ListGet(ctx context.Context, userID xid.ID, offset, limit int, ttl *time.Duration) ([]entity.Post, error) {
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
	listCmd := pipe.LRange(ctx, userID.String(), int64(offset), int64(offset+limit-1))

	if _, err := pipe.Exec(ctx); err != nil {
		return nil, err
	}

	if len(listCmd.Val()) != 0 {
		res := make([]entity.Post, 0, len(listCmd.Val()))
		if err := listCmd.ScanSlice(&res); err != nil {
			return nil, err
		}

		if ttl != nil {
			go expfn()
		}

		// remove empty elements
		if v := res[len(res)-1]; v.PostID.IsZero() && v.AuthorID.IsZero() {
			res = res[:len(res)-1]
		}

		return res, nil
	}

	if lenCmd.Val() == 0 {
		return nil, repoerr.ErrNotFound
	}

	go expfn()

	return []entity.Post{}, nil
}
