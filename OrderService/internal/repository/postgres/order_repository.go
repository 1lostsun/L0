package postgres

import (
	"OrderService/internal/model"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *model.OrderRequest) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (repository *orderRepository) SaveOrder(ctx context.Context, order *model.OrderRequest) error {
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

func (repository *orderRepository) GetOrderByID(ctx context.Context, id int64) (*model.OrderRequest, error) {
	var order model.OrderRequest
	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("id = ?", id).First(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}
