package post

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gcfg "github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	postApi "github.com/Karzoug/meower-timeline-service/pkg/proto/grpc/post/v1"
)

type Client struct {
	c postApi.PostServiceClient
}

func NewServiceClient(cfg gcfg.Config) (Client, error) {
	conn, err := grpc.NewClient(
		cfg.URI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return Client{}, fmt.Errorf("could not connect to post microservice: %w", err)
	}
	return Client{c: postApi.NewPostServiceClient(conn)}, nil
}
