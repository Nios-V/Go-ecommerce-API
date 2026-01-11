package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	BaseRepository[models.Cart]
	WithTx(tx *gorm.DB) CartRepository
	GetCurrentCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error)
}

type cartRepository struct {
	BaseRepository[models.Cart]
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		BaseRepository: NewBaseRepository[models.Cart](db),
		db:             db,
	}
}

func (r *cartRepository) WithTx(tx *gorm.DB) CartRepository {
	return NewCartRepository(tx)
}

func (r *cartRepository) GetCurrentCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.WithContext(ctx).
		Preload("Items", "is_purchased = ?", false).
		Preload("Items.Product").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err != nil {
		return nil, err
	}
	return &cart, nil
}
