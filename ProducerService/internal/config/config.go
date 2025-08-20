package config

import (
	"OrderService/ProducerService/internal/kafka"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	KafkaCfg kafka.Config
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
