package kafka

import (
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Kafka struct {
	kw *kafka.Writer
}

func (k *Kafka) Produce(ctx context.Context, msg []byte) error {
	err := k.kw.WriteMessages(
		ctx, kafka.Message{
			Value: msg,
		})

	if err != nil {
		return err
	}

	return nil
}
