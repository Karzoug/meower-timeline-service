package service

import (
	"context"
	"errors"
	"slices"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	repoerr "github.com/Karzoug/meower-timeline-service/internal/timeline/repo"
)

func (ts TimelineService) PushTimelinePost(ctx context.Context, userID xid.ID, post entity.Post) error {
	if err := ts.repo.ExistedListPushPost(ctx, userID, post, int64(ts.cfg.Limit)); err != nil {
		return ucerr.NewInternalError(err)
	}

	return nil
}

func (ts TimelineService) DeleteTimelinePost(ctx context.Context, userID xid.ID, post entity.Post) error {
	if err := ts.repo.ExistedListDeletePost(ctx, userID, post); err != nil {
		return ucerr.NewInternalError(err)
	}

	return nil
}

func (ts TimelineService) DeleteTimeline(ctx context.Context, userID xid.ID) error {
	if err := ts.repo.ExistedListDelete(ctx, userID); err != nil {
		return ucerr.NewInternalError(err)
	}

	return nil
}

func (ts TimelineService) SubscribeOnUser(ctx context.Context, userID, targetUserID xid.ID) error {
	posts, _, err := ts.repo.ListGetOlder(ctx, userID, nil, ts.cfg.Limit, nil)
	if err != nil {
		if errors.Is(err, repoerr.ErrKeyNotFound) {
			return nil
		}
		return ucerr.NewInternalError(err)
	}

	targetPosts, err := ts.postService.ListPostIDsByUserIDs(ctx, targetUserID, []xid.ID{targetUserID}, ts.cfg.Limit)
	if err != nil {
		return ucerr.NewInternalError(err)
	}

	// merge results by time
	res := make([]entity.Post, min(ts.cfg.Limit, len(posts)+len(targetPosts)))
	var i, j int
	for k := 0; k < len(res); k++ {
		if i == len(posts) {
			res[k] = targetPosts[j]
			j++
			continue
		}
		if j == len(targetPosts) {
			res[k] = posts[i]
			i++
			continue
		}
		if posts[i].PostID.Time().After(targetPosts[j].PostID.Time()) {
			res[k] = posts[i]
			i++
		} else {
			res[k] = targetPosts[j]
			j++
		}
	}

	if err := ts.repo.ListSet(ctx, userID, res, ts.cfg.TTL); err != nil {
		return ucerr.NewInternalError(err)
	}

	return nil
}

func (ts TimelineService) UnsubscribeFromUser(ctx context.Context, userID, targetUserID xid.ID) error {
	posts, _, err := ts.repo.ListGetOlder(ctx, userID, nil, ts.cfg.Limit, nil)
	if err != nil {
		if errors.Is(err, repoerr.ErrKeyNotFound) || errors.Is(err, repoerr.ErrValueNotFound) {
			return nil
		}
		return ucerr.NewInternalError(err)
	}

	res := slices.DeleteFunc(posts, func(p entity.Post) bool {
		return p.AuthorID.Compare(targetUserID) == 0
	})

	// TODO: check if user has too low number of posts
	// and rebuild timeline if needed

	if err := ts.repo.ListSet(ctx, userID, res, ts.cfg.TTL); err != nil {
		return ucerr.NewInternalError(err)
	}

	return nil
}
