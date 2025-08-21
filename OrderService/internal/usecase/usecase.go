package usecase

import (
	"OrderService/OrderService/internal/kafka"
	"OrderService/OrderService/internal/service"
	"golang.org/x/net/context"
	"log"
)

type Usecase struct {
	kafka        *kafka.Kafka
	orderService *service.OrderService
}

func New(kafka *kafka.Kafka) *Usecase {
	return &Usecase{
		kafka: kafka,
	}
}

func (uc *Usecase) ReadKafkaMessage(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := uc.kafka.Consume(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return ctx.Err()
				}
				log.Println("kafka consume err: ", err)
				continue
			}

		}
	}
}
