package service

import (
	"OrderService/OrderService/internal/model"
	"OrderService/OrderService/internal/repository/postgres"
	"OrderService/OrderService/internal/repository/redis"
	"encoding/json"
	"golang.org/x/net/context"
)

type OrderService struct {
	pgRepository *postgres.OrderRepository
	redis        *redis.Redis
}

func New(pgRepository *postgres.OrderRepository, redis *redis.Redis) *OrderService {
	return &OrderService{
		pgRepository: pgRepository,
		redis:        redis,
	}
}

func (os *OrderService) ProcessOrder(ctx context.Context, msg []byte) (*model.Order, error) {
	var order model.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		return nil, err
	}

	if err := os.redis.Set(order.OrderUID, string(msg)); err != nil {
		return nil, err
	}

	if err := os.pgRepository.SaveOrder(ctx, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
