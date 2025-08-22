package kafka

import (
	"github.com/caarlos0/env/v10"
	"github.com/segmentio/kafka-go"
	"log"
)

type Config struct {
	Addr  string `env:"KAFKA_ADDR" envDefault:"kafka:9092"`
	Topic string `env:"KAFKA_TOPIC" envDefault:"testTopic"`
}

func NewKafkaCfg() *Kafka {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &Kafka{
		&kafka.Writer{
			Addr:         kafka.TCP(cfg.Addr),
			Topic:        cfg.Topic,
			RequiredAcks: kafka.RequireAll,
			Async:        false,
		},
	}
}
