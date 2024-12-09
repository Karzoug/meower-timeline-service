package post

import (
	"context"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	postApi "github.com/Karzoug/meower-timeline-service/pkg/proto/grpc/post/v1"
)

func (c Client) ListPostIDsByUserIDs(ctx context.Context, reqUserID xid.ID, userIDs []xid.ID, limit int) ([]entity.Post, error) {
	ctx = grpc.ContextWithUserID(ctx, reqUserID)

	parents := make([]string, len(userIDs))
	for i := range userIDs {
		parents[i] = userIDs[i].String()
	}

	postIDs, err := c.c.ListPostIdProjections(ctx, &postApi.ListPostIdProjectionsRequest{
		Parents:  parents,
		PageSize: int32(limit), //nolint:gosec
	})
	if err != nil {
		return nil, err
	}

	res := make([]entity.Post, 0, len(postIDs.PostIdProjections))
	for i := range postIDs.PostIdProjections {
		if postIDs.PostIdProjections[i] == nil {
			continue
		}
		postID, _ := xid.FromString(postIDs.PostIdProjections[i].Id)
		authorID, _ := xid.FromString(postIDs.PostIdProjections[i].AuthorId)
		res = append(res, entity.Post{
			PostID:   postID,
			AuthorID: authorID,
		})
	}

	return res, nil
}
