package service

import (
	"context"

	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
)

func (ts TimelineService) GetTimeline(ctx context.Context, reqUserID, userID xid.ID, pgn PaginationOptions) ([]entity.Post, error) {
	return nil, ucerr.NewError(nil, "not implemented", codes.Unimplemented)
}
