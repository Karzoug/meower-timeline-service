package service

import (
	"context"
	"errors"
	"fmt"

	ocodes "go.opentelemetry.io/otel/codes"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	repoerr "github.com/Karzoug/meower-timeline-service/internal/timeline/repo"
)

const preffixSpanName = "TimelineService.Service/"

func (ts TimelineService) GetTimeline(ctx context.Context, reqUserID, userID xid.ID, pgn PaginationOptions) ([]entity.Post, error) {
	ctx, span := ts.tracer.Start(ctx, preffixSpanName+"GetTimeline")
	defer span.End()

	if pgn.Offset < 0 {
		return nil, ucerr.NewError(
			nil,
			"invalid pagination parameter: negative offset",
			codes.InvalidArgument,
		)
	}

	if pgn.Limit < 0 {
		return nil, ucerr.NewError(
			nil,
			"invalid pagination parameter: negative size",
			codes.InvalidArgument,
		)
	}
	if pgn.Limit == 0 {
		pgn.Limit = 100
	} else if pgn.Limit > 100 {
		pgn.Limit = 100
	}

	if reqUserID.Compare(userID) != 0 {
		return nil, ucerr.NewError(nil, "user id mismatch", codes.PermissionDenied)
	}
	if pgn.Offset >= ts.cfg.Limit {
		return nil, ucerr.NewError(nil, "end of timeline", codes.OutOfRange)
	}
	if pgn.Offset+pgn.Limit > ts.cfg.Limit {
		pgn.Limit = ts.cfg.Limit - pgn.Offset
	}

	res, err := ts.repo.ListGet(ctx, userID, pgn.Offset, pgn.Limit, &ts.cfg.TTL)
	if nil == err {
		return res, nil
	}
	if errors.Is(err, repoerr.ErrNotFound) {
		// not found timeline in cache -> build it from scratch
		span.AddEvent("timeline not found in cache")
		return ts.getTimelineFromScratch(userID)
	}

	switch {
	case errors.Is(err, context.Canceled):
		return nil, ucerr.NewError(err, "request canceled", codes.Canceled)
	case errors.Is(err, context.DeadlineExceeded):
		return nil, ucerr.NewError(err, "request timeout", codes.DeadlineExceeded)
	default:
		return nil, ucerr.NewInternalError(err)
	}
}

func (ts TimelineService) getTimelineFromScratch(userID xid.ID) ([]entity.Post, error) {
	ctx, cancel := context.WithTimeout(ts.shutdownCtx, ts.cfg.BuildTimeout)
	defer cancel()

	ctx, span := ts.tracer.Start(ts.shutdownCtx, preffixSpanName+"BuildTimelineFromScratch")
	defer span.End()

	res := ts.buildTimelineFromScratch(ctx, userID)

	fn := func(r singleflight.Result) ([]entity.Post, error) {
		if r.Err != nil {
			return nil, ucerr.NewInternalError(r.Err)
		}
		res, ok := r.Val.([]entity.Post)
		if !ok {
			return nil, ucerr.NewInternalError(fmt.Errorf("invalid type assertion: want []entity.Post, got %T", r.Val))
		}

		if err := ts.set(ctx, userID, res); err != nil {
			ts.logger.Error().
				Err(err).
				Msg("failed to set timeline")
		}

		return res, nil
	}

	var (
		err   error
		posts []entity.Post
	)

	select {
	case <-ctx.Done():
		err = ucerr.NewError(ctx.Err(), "request canceled", codes.Canceled)
	case r := <-res:
		posts, err = fn(r)
	}

	if err != nil {
		span.SetStatus(ocodes.Error, err.Error())
	}

	return posts, err
}

func (ts TimelineService) buildTimelineFromScratch(ctx context.Context, userID xid.ID) <-chan singleflight.Result {
	// suppression mechanism for set of the same requests
	return ts.syncGroup.DoChan(userID.String(), func() (any, error) {
		followingIDs, err := ts.relationService.ListNotMutedFollowingIDs(ctx, userID)
		if err != nil {
			return nil, err
		}

		posts, err := ts.postService.ListPostIDsByUserIDs(ctx, userID, followingIDs, ts.cfg.Limit)
		if err != nil {
			return nil, err
		}

		return posts, nil
	})
}

func (ts TimelineService) set(ctx context.Context, userID xid.ID, records []entity.Post) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if err := ts.repo.ListSet(ctx, userID, records, ts.cfg.TTL); err != nil {
		return err
	}

	return nil
}
