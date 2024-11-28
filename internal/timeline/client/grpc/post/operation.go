package post

import (
	"context"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	postApi "github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc/gen/post/v1"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
)

func (c Client) ListPostByUserIDs(ctx context.Context, reqUserID string, userIDs []string, limit, offset int) ([]entity.Post, error) {
	ctx = grpc.ContextWithUserID(ctx, reqUserID)

	postIDs, err := c.c.ListPostIds(ctx, &postApi.ListPostIdsRequest{
		AuthorUserIds: userIDs,
		Pagination: &postApi.PaginationRequest{
			Offset: int64(offset),
			Limit:  int64(limit),
		},
	})
	if err != nil {
		return nil, err
	}

	res := make([]entity.Post, 0, len(postIDs.PostIds))

	for i := range postIDs.PostIds {
		if postIDs.PostIds[i] == nil {
			continue
		}
		res = append(res, entity.Post{
			PostID:   postIDs.PostIds[i].Id,
			AuthorID: postIDs.PostIds[i].AuthorId,
		})
	}

	return res, nil
}
