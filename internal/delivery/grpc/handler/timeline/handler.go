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
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service/option"
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

	authUserID := auth.UserIDFromContext(ctx)

	userID, err := xid.FromString(req.Parent)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id: "+req.Parent)
	}

	if userID.Compare(authUserID) != 0 {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	opt := option.Pagination{
		Size: int(req.PageSize),
	}
	switch t := req.ByOneof.(type) {
	case *gen.ListTimelineRequest_PrevPageToken:
		opt = option.Pagination{
			PrevToken: t.PrevPageToken,
		}
	case *gen.ListTimelineRequest_NextPageToken:
		opt = option.Pagination{
			NextToken: t.NextPageToken,
		}
	}

	res, err := h.timelineService.GetTimeline(ctx, authUserID, opt)
	if err != nil {
		return nil, err
	}

	return &gen.ListTimelineResponse{
		Posts:         converter.ToProtoPosts(res.Posts),
		NextPageToken: res.NextToken,
		PrevPageToken: res.PrevToken,
	}, nil
}
