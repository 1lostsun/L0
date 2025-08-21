package config

import (
	"OrderService/OrderService/internal/kafka"
	"OrderService/OrderService/internal/repository/postgres"
	"OrderService/OrderService/internal/repository/redis"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	KafkaCfg kafka.Config
	PgCfg    postgres.Config
	RedisCfg redis.Config
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
