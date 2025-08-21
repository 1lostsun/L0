package redis

import "github.com/redis/go-redis/v9"

type Config struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"order-redis:6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

func NewConfig(config Config) *Redis {
	return &Redis{
		rdb: redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		}),
	}
}
