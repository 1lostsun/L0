package usecase

import (
	"OrderService/ProducerService/internal/kafka"
	"OrderService/ProducerService/internal/model"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
)

type Usecase struct {
	kafka *kafka.Kafka
}

func New(kafka *kafka.Kafka) *Usecase {
	return &Usecase{
		kafka: kafka,
	}
}

func (uc *Usecase) SendKafkaMsg(ctx context.Context, req *model.OrderRequest) error {
	msg, err := json.Marshal(req)
	if err != nil {
		return err
	}

	fmt.Println(msg)

	if err := uc.kafka.Produce(ctx, msg); err != nil {
		return err
	}

	return nil
}
