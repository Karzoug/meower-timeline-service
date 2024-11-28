package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const userKey string = "x-user"

func ContextWithUserID(ctx context.Context, userID string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(userKey, userID))
}
