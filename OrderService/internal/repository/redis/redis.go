package redis

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RedisDB *redis.Client
}
