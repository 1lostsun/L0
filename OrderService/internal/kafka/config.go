package kafka

import "github.com/segmentio/kafka-go"

type Config struct {
	Broker  string `env:"KAFKA_BROKER" envDefault:"kafka:9092"`
	Topic   string `env:"KAFKA_TOPIC" envDefault:"ordersTopic"`
	GroupID string `env:"KAFKA_GROUP_ID" envDefault:"order-service"`
}

func NewKafkaCfg(config Config) *Kafka {
	return &Kafka{
		kr: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.Broker},
			Topic:   config.Topic,
			GroupID: config.GroupID,
		}),
	}
}
