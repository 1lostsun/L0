package config

import (
	"OrderService/internal/kafka"
	"OrderService/internal/repository/postgres"
	"OrderService/internal/repository/redis"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	KafkaCfg kafka.Config
	PgCfg    postgres.Config
	RedisCfg redis.Config
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg.KafkaCfg); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.PgCfg); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.RedisCfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
