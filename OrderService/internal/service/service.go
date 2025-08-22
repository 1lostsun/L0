package service

import (
	"OrderService/internal/model"
	"OrderService/internal/repository/postgres"
	"OrderService/internal/repository/redis"
	"encoding/json"
	"errors"
	r "github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

const defaultTTL = 10 * time.Minute

type Service struct {
	pgRepository *postgres.PGRepository
	redis        *redis.Redis
}

func New(pgRepository *postgres.PGRepository, redis *redis.Redis) *Service {
	return &Service{
		pgRepository: pgRepository,
		redis:        redis,
	}
}

func (orderService *Service) ProcessOrder(ctx context.Context, ordRequest model.OrderRequest) error {
	order := model.Order{
		OrderUID:        ordRequest.OrderUID,
		TrackNumber:     ordRequest.TrackNumber,
		Entry:           ordRequest.Entry,
		Locale:          ordRequest.Locale,
		CustomerID:      ordRequest.CustomerID,
		DeliveryService: ordRequest.DeliveryService,
		DateCreated:     ordRequest.DateCreated,
		ShardKey:        ordRequest.ShardKey,
		SmID:            ordRequest.SmID,
		OOFShard:        ordRequest.OOFShard,
		Delivery: model.Delivery{
			Name:    ordRequest.Delivery.Name,
			Phone:   ordRequest.Delivery.Phone,
			Zip:     ordRequest.Delivery.Zip,
			City:    ordRequest.Delivery.City,
			Address: ordRequest.Delivery.Address,
			Region:  ordRequest.Delivery.Region,
			Email:   ordRequest.Delivery.Email,
		},
		Payment: model.Payment{
			Transaction:  ordRequest.Payment.Transaction,
			RequestID:    ordRequest.Payment.RequestID,
			Currency:     ordRequest.Payment.Currency,
			Provider:     ordRequest.Payment.Provider,
			Amount:       ordRequest.Payment.Amount,
			PaymentDT:    ordRequest.Payment.PaymentDT,
			Bank:         ordRequest.Payment.Bank,
			DeliveryCost: ordRequest.Payment.DeliveryCost,
			GoodsTotal:   ordRequest.Payment.GoodsTotal,
			CustomFee:    ordRequest.Payment.CustomFee,
		},
	}

	for _, i := range ordRequest.Items {
		order.Items = append(order.Items, model.Item{
			ChrtID:      i.ChrtID,
			TrackNumber: i.TrackNumber,
			Price:       i.Price,
			RID:         i.RID,
			Name:        i.Name,
			Sale:        i.Sale,
			Size:        i.Size,
			TotalPrice:  i.TotalPrice,
			NM_ID:       i.NM_ID,
			Brand:       i.Brand,
			Status:      i.Status,
		})
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := orderService.redis.RedisDB.Set(ctx, order.OrderUID, orderJson, defaultTTL).Err(); err != nil {
		return err
	}

	if err := orderService.pgRepository.SaveOrder(ctx, &order); err != nil {
		return err
	}

	return nil
}

func (orderService *Service) GetOrderByUID(ctx context.Context, orderUID string) (*model.Order, error) {
	redisOrder, err := orderService.getOrderFromRedis(ctx, orderUID)
	if errors.Is(err, r.Nil) {
		pgOrder, err := orderService.pgRepository.GetOrderByUID(ctx, orderUID)
		if err != nil {
			return nil, err
		}

		return pgOrder, nil
	}

	return redisOrder, nil
}

func (orderService *Service) getOrderFromRedis(ctx context.Context, orderUID string) (*model.Order, error) {
	val, err := orderService.redis.RedisDB.Get(ctx, orderUID).Result()

	if err != nil {
		return nil, err
	}

	var order model.Order
	if err := json.Unmarshal([]byte(val), &order); err != nil {
		return nil, err
	}

	return &order, nil
}
