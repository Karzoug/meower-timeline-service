package result

import "github.com/Karzoug/meower-timeline-service/internal/timeline/entity"

type ListPost struct {
	Posts     []entity.Post
	PrevToken string
	NextToken string
}
