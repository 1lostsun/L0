package usecase

import (
	"OrderService/internal/kafka"
	"OrderService/internal/model"
	"OrderService/internal/service"
	"encoding/json"
	"golang.org/x/net/context"
	"log"
)

type Usecase struct {
	kafka        *kafka.Kafka
	orderService *service.Service
}

func New(kafka *kafka.Kafka, orderService *service.Service) *Usecase {
	return &Usecase{
		kafka:        kafka,
		orderService: orderService,
	}
}

func (uc *Usecase) ReadKafkaMessage(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := uc.kafka.Consume(ctx)
			var ord model.OrderRequest
			if err := json.Unmarshal(msg, &ord); err != nil {
				log.Println(err)
			}

			if err != nil {
				if ctx.Err() != nil {
					return ctx.Err()
				}
				log.Println("kafka consume err: ", err)
				continue
			}
			if err := uc.orderService.ProcessOrder(ctx, ord); err != nil {
				log.Println("process order err: ", err)
			}
		}
	}
}
