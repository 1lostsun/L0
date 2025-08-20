package usecase

import (
	"OrderService/OrderService/internal/kafka"
	"OrderService/OrderService/internal/repository/postgres"
)

type Usecase struct {
	kafka        kafka.Kafka
	pgRepository postgres.OrderRepository
}

func New(kafka kafka.Kafka) *Usecase {
	return &Usecase{
		kafka: kafka,
	}
}

func (uc *Usecase) SaveOrderAndSendKafkaMsg() error {
	return nil
}
