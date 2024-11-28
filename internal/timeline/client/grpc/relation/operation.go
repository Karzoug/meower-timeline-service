package relation

import (
	"context"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	relationApi "github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc/gen/relation/v1"
)

func (c Client) ListAllFollowerIDs(ctx context.Context, reqUserID, userID string) ([]string, error) {
	ctx = grpc.ContextWithUserID(ctx, reqUserID)

	fresp, err := c.c.ListFollowers(ctx,
		&relationApi.ListFollowersRequest{
			UserId: userID,
			Pagination: &relationApi.PaginationRequest{
				Limit: -1, // require all followings
			},
		})
	if err != nil {
		return nil, err
	}

	followerIDs := make([]string, 0, len(fresp.Followers))
	for i := range fresp.Followers {
		if fresp.Followers[i] == nil || fresp.Followers[i].Hidden {
			continue
		}
		followerIDs = append(followerIDs, fresp.Followers[i].Id)
	}

	return followerIDs, nil
}

func (c Client) ListAllNotHiddenFollowingIDs(ctx context.Context, reqUserID, userID string) ([]string, error) {
	ctx = grpc.ContextWithUserID(ctx, reqUserID)

	fresp, err := c.c.ListFollowings(ctx,
		&relationApi.ListFollowingsRequest{
			UserId: userID,
			Pagination: &relationApi.PaginationRequest{
				Limit: -1, // require all followings
			},
		})
	if err != nil {
		return nil, err
	}

	followingIDs := make([]string, 0, len(fresp.Followings))
	for i := range fresp.Followings {
		if fresp.Followings[i] == nil ||
			fresp.Followings[i].Hidden {
			continue
		}
		followingIDs = append(followingIDs, fresp.Followings[i].Id)
	}

	return followingIDs, nil
}
