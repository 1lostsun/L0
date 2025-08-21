package redis

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

const defaultTTL = 10 * time.Minute

type Redis struct {
	rdb *redis.Client
}

func (r *Redis) Set(key string, value []byte) error {
	return r.rdb.Set(context.Background(), key, value, defaultTTL).Err()
}

func (r *Redis) Get(key string) ([]byte, error) {
	return r.rdb.Get(context.Background(), key).Bytes()
}
