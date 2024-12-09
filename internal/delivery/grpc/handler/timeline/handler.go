package user

import (
	"context"

	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Karzoug/meower-common-go/auth"

	"github.com/Karzoug/meower-timeline-service/internal/delivery/grpc/converter"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service"
	gen "github.com/Karzoug/meower-timeline-service/pkg/proto/grpc/timeline/v1"
)

func RegisterService(us service.TimelineService) func(grpcServer *grpc.Server) {
	hdl := handlers{
		timelineService: us,
	}
	return func(grpcServer *grpc.Server) {
		gen.RegisterTimelineServiceServer(grpcServer, hdl)
	}
}

type handlers struct {
	gen.UnimplementedTimelineServiceServer
	timelineService service.TimelineService
}

func (h handlers) ListTimeline(ctx context.Context, req *gen.ListTimelineRequest) (*gen.ListTimelineResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	userID, err := xid.FromString(req.Parent)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id: "+req.Parent)
	}

	posts, err := h.timelineService.GetTimeline(ctx, auth.UserIDFromContext(ctx), userID, service.PaginationOptions{
		Offset: int(req.PageOffset),
		Limit:  int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	return &gen.ListTimelineResponse{
		Posts: converter.ToProtoPosts(posts),
	}, nil
}
