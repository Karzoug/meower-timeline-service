package relation

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gcfg "github.com/Karzoug/meower-timeline-service/internal/timeline/client/grpc"
	relationApi "github.com/Karzoug/meower-timeline-service/pkg/proto/grpc/relation/v1"
)

type Client struct {
	c relationApi.RelationServiceClient
}

func NewServiceClient(cfg gcfg.Config) (Client, error) {
	conn, err := grpc.NewClient(
		cfg.URI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return Client{}, fmt.Errorf("could not connect to relation microservice: %w", err)
	}
	return Client{c: relationApi.NewRelationServiceClient(conn)}, nil
}
