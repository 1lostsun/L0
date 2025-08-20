package usecase

import "OrderService/internal/kafka"

type Usecase struct {
	kafka kafka.Kafka
}

func New(kafka kafka.Kafka) *Usecase {
	return &Usecase{
		kafka: kafka,
	}
}
