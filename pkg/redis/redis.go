package redis

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	redis.UniversalClient
}

func NewDB(ctx context.Context, cfg Config) (DB, error) {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: cfg.Addrs,
	})

	if cmd := rdb.Ping(ctx); cmd.Err() != nil {
		return DB{}, cmd.Err()
	}

	if cfg.EnableTracing {
		if err := redisotel.InstrumentTracing(rdb); err != nil {
			return DB{}, err
		}
	}

	return DB{UniversalClient: rdb}, nil
}

func (db DB) Close(_ context.Context) error {
	return db.UniversalClient.Close()
}
