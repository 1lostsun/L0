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

func (orderService *OrderService) ProcessOrder(ctx context.Context, msg []byte) error {
	var order model.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		return err
	}

	if err := orderService.redis.Set(order.OrderUID, msg); err != nil {
		return err
	}

	if err := orderService.pgRepository.SaveOrder(ctx, &order); err != nil {
		return err
	}

	return nil
}

func (orderService *OrderService) GetOrderByUID(ctx context.Context, orderUID string) ([]byte, error) {
	redisOrder, err := orderService.redis.Get(orderUID)
	if err != nil {
		pgOrder, err := orderService.pgRepository.GetOrderByUID(ctx, orderUID)
		if err != nil {
			return nil, err
		}

		return json.Marshal(pgOrder)
	}

	return redisOrder, nil
}
