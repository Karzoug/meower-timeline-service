package redis

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	*redis.ClusterClient
}

func NewClusterDB(cfg Config) (DB, error) {
	opts, err := redis.ParseClusterURL(cfg.URI)
	if err != nil {
		return DB{}, err
	}
	// To route commands by latency or randomly, enable one of the following.
	// opts.RouteByLatency = true
	// opts.RouteRandomly = true

	rdb := redis.NewClusterClient(opts)

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		return DB{}, err
	}

	return DB{ClusterClient: rdb}, nil
}

func (db DB) Close(_ context.Context) error {
	return db.ClusterClient.Close()
}
