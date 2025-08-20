package usecase

import (
	"OrderService/ProducerService/internal/kafka"
)

type Usecase struct {
	kafka kafka.Kafka
}

func New(kafka kafka.Kafka) *Usecase {
	return &Usecase{
		kafka: kafka,
	}
}

func (uc *Usecase) SendKafkaMsg() error {
	return nil
}
