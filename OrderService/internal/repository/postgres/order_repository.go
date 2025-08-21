package postgres

import (
	"OrderService/OrderService/internal/model"
	"context"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repository *OrderRepository) SaveOrder(ctx context.Context, order *model.Order) error {
	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) GetOrderByUID(ctx context.Context, id string) (*model.Order, error) {
	var order model.Order
	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("order_uid = ?", id).First(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}
