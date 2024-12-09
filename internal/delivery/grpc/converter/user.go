package converter

import (
	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	gen "github.com/Karzoug/meower-timeline-service/pkg/proto/grpc/timeline/v1"
)

func ToProtoPost(p entity.Post) *gen.Post {
	return &gen.Post{
		AuthorId: p.AuthorID.String(),
		PostId:   p.PostID.String(),
		IsRepost: p.IsRepost,
	}
}

func ToProtoPosts(pp []entity.Post) []*gen.Post {
	res := make([]*gen.Post, len(pp))
	for i := range pp {
		res[i] = ToProtoPost(pp[i])
	}
	return res
}
