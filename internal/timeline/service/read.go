package service

import (
	"context"
	"errors"
	"fmt"

	ocodes "go.opentelemetry.io/otel/codes"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-common-go/trace/otlp"
	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	repoerr "github.com/Karzoug/meower-timeline-service/internal/timeline/repo"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service/option"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service/result"
)

const preffixSpanName = "TimelineService.Service/"

func (ts TimelineService) GetTimeline(ctx context.Context, authUserID xid.ID, pgn option.Pagination) (result.ListPost, error) {
	ctx, span := otlp.AddSpan(ctx, preffixSpanName+"GetTimeline")
	defer span.End()

	if pgn.Size < 0 {
		return result.ListPost{}, ucerr.NewError(
			nil,
			"invalid pagination parameter: negative size",
			codes.InvalidArgument,
		)
	}
	if pgn.Size == 0 {
		pgn.Size = 100
	} else if pgn.Size > 100 {
		pgn.Size = 100
	}

	var err error
	if pgn.PrevToken != "" {
		token, err := entity.FromString(pgn.PrevToken)
		if err != nil {
			return result.ListPost{}, ucerr.NewError(
				err,
				"invalid pagination parameter: previous token",
				codes.InvalidArgument,
			)
		}

		var (
			res       []entity.Post
			nextToken *entity.Post
		)
		res, nextToken, err = ts.repo.ListGetNewer(ctx, authUserID, token, pgn.Size, &ts.cfg.TTL)
		if nil == err {
			res := result.ListPost{
				Posts:     res,
				NextToken: pgn.PrevToken,
			}
			if nextToken != nil {
				res.PrevToken = nextToken.String()
			}
			return res, nil
		}

		if errors.Is(err, repoerr.ErrValueNotFound) {
			return result.ListPost{}, ucerr.NewError(
				err,
				"invalid pagination token",
				codes.InvalidArgument,
			)
		}

		return result.ListPost{}, ucerr.NewError(
			err,
			"failed to get timeline",
			codes.Internal,
		)
	}

	var token *entity.Post
	if pgn.NextToken != "" {
		t, err := entity.FromString(pgn.NextToken)
		if err != nil {
			return result.ListPost{}, ucerr.NewError(
				err,
				"invalid pagination parameter: next token",
				codes.InvalidArgument,
			)
		}
		token = &t
	}

	res, nextToken, err := ts.repo.ListGetOlder(ctx, authUserID, token, pgn.Size, &ts.cfg.TTL)
	if nil == err {
		res := result.ListPost{
			Posts:     res,
			PrevToken: pgn.NextToken,
		}
		if nextToken != nil {
			res.NextToken = nextToken.String()
		}
		return res, nil
	}

	if errors.Is(err, repoerr.ErrValueNotFound) {
		return result.ListPost{}, ucerr.NewError(
			err,
			"invalid pagination token",
			codes.InvalidArgument,
		)
	}

	if errors.Is(err, repoerr.ErrKeyNotFound) {
		// not found timeline in cache -> build it from scratch
		span.AddEvent("timeline not found in cache")
		posts, err := ts.getTimelineFromScratch(authUserID)
		if err != nil {
			return result.ListPost{}, err
		}

		if pgn.NextToken != "" || pgn.PrevToken != "" {
			return result.ListPost{
					Posts: posts[0:pgn.Size],
				},
				ucerr.NewError(nil, "token is expired, returned only first page of timeline", codes.FailedPrecondition)
		}
		return result.ListPost{Posts: posts}, nil
	}

	switch {
	case errors.Is(err, context.Canceled):
		return result.ListPost{}, ucerr.NewError(err, "request canceled", codes.Canceled)
	case errors.Is(err, context.DeadlineExceeded):
		return result.ListPost{}, ucerr.NewError(err, "request timeout", codes.DeadlineExceeded)
	default:
		return result.ListPost{}, ucerr.NewInternalError(err)
	}
}

func (ts TimelineService) getTimelineFromScratch(userID xid.ID) ([]entity.Post, error) {
	ctx, cancel := context.WithTimeout(ts.shutdownCtx, ts.cfg.BuildTimeout)
	defer cancel()

	ctx, span := otlp.AddSpan(ctx, preffixSpanName+"BuildTimelineFromScratch")
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

		if err := ts.repo.ListSet(ctx, userID, res, ts.cfg.TTL); err != nil {
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
