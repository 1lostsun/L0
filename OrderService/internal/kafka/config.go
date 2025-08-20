package kafka

import (
	"github.com/segmentio/kafka-go"
)

type Config struct {
	Addr  string `env:"KAFKA_ADDR" envDefault:"localhost:9092"`
	Topic string `env:"KAFKA_TOPIC" envDefault:"testTopic"`
}

func NewKafkaCfg(config Config) *Kafka {
	return &Kafka{
		&kafka.Writer{
			Addr:         kafka.TCP(config.Addr),
			Topic:        config.Topic,
			RequiredAcks: kafka.RequireAll,
			Async:        false,
		},
	}
}
