package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	BaseRepository[models.Order]
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]models.Order, error)
}

type orderRepository struct {
	BaseRepository[models.Order]
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		BaseRepository: NewBaseRepository[models.Order](db),
		db:             db,
	}
}

func (r *orderRepository) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&orders).
		Order("created_at DESC").Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
