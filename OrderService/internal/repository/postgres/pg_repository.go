package postgres

import (
	"OrderService/internal/model"
	"context"
	"gorm.io/gorm"
)

type PGRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *PGRepository {
	return &PGRepository{
		db: db,
	}
}

func (repository *PGRepository) SaveOrder(ctx context.Context, order *model.Order) error {
	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(order).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Save(&order.Delivery).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Save(&order.Payment).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Save(&order.Items).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *PGRepository) GetOrderByUID(ctx context.Context, id string) (*model.Order, error) {
	var order model.Order
	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("order_uid = ?", id).
			Preload("Delivery").
			Preload("Payment").
			Preload("Items").
			First(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}
