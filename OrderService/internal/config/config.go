package config

import (
	"OrderService/internal/kafka"
	"OrderService/internal/repository/postgres"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	KafkaCfg kafka.Config
	PgCfg    postgres.Config
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
