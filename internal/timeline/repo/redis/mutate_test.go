package redis

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	rc "github.com/testcontainers/testcontainers-go/modules/redis"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	"github.com/Karzoug/meower-timeline-service/pkg/redis"
)

func Test_repo_ExistedListDeletePost(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	redisContainer, err := rc.Run(ctx, "redis:6")
	require.NoError(t, err)
	defer redisContainer.Terminate(context.TODO()) //nolint:errcheck

	url, err := redisContainer.ConnectionString(ctx)
	require.NoError(t, err)

	url, found := strings.CutPrefix(url, "redis://")
	require.True(t, found)

	db, err := redis.NewDB(ctx, redis.Config{Addrs: []string{url}})
	require.NoError(t, err)

	r := repo{
		db:     db,
		logger: zerolog.New(os.Stdout),
	}

	userID := xid.New()

	post1 := entity.Post{
		PostID:   xid.New(),
		AuthorID: xid.New(),
	}

	err = r.ListSet(ctx,
		userID,
		[]entity.Post{
			post1,
			{
				PostID:   xid.New(),
				AuthorID: xid.New(),
			},
		}, time.Hour)
	require.NoError(t, err)

	err = r.ExistedListDeletePost(ctx, userID, post1)
	require.NoError(t, err)

	resp, err := r.ListGet(ctx, userID, 0, 10, nil)
	require.NoError(t, err)

	assert.Len(t, resp, 1)
}
