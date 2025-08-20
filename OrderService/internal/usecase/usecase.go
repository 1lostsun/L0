package usecase

import (
	"OrderService/internal/kafka"
	"OrderService/internal/repository/postgres"
)

type Usecase interface {
}

type usecase struct {
	kafka        kafka.Kafka
	pgRepository postgres.OrderRepository
}

func New(kafka kafka.Kafka) Usecase {
	return &usecase{
		kafka: kafka,
	}
}

func (uc *usecase) SaveOrderAndSendKafkaMsg() error {

}
