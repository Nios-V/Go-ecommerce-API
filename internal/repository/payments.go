package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	BaseRepository[models.Payment]
	GetByOrderID(ctx context.Context, orderID uuid.UUID) (*models.Payment, error)
}

type paymentRepository struct {
	BaseRepository[models.Payment]
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		BaseRepository: NewBaseRepository[models.Payment](db),
		db:             db,
	}
}

func (r *paymentRepository) GetByOrderID(ctx context.Context, orderID uuid.UUID) (*models.Payment, error) {
	var payment models.Payment

	err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
