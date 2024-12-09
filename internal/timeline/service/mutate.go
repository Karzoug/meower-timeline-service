package service

import (
	"context"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
)

func (ts TimelineService) PushTimelinePost(ctx context.Context, userID xid.ID, post entity.Post) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (ts TimelineService) DeleteTimelinePost(ctx context.Context, userID xid.ID, post entity.Post) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (ts TimelineService) DeleteTimeline(ctx context.Context, userID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (ts TimelineService) SubscribeOnUser(ctx context.Context, userID, targetUserID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}

func (ts TimelineService) UnsubscribeFromUser(ctx context.Context, userID, targetUserID xid.ID) error {
	return ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}
