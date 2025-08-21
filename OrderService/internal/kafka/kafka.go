package kafka

import (
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Kafka struct {
	kr *kafka.Reader
}

func (k *Kafka) Consume(ctx context.Context) ([]byte, error) {
	msg, err := k.kr.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return msg.Value, nil
}
