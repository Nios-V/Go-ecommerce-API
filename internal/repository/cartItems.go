package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItemRepository interface {
	BaseRepository[models.CartItem]
	WithTx(tx *gorm.DB) CartItemRepository
	GetByCartAndProduct(ctx context.Context, cartID, productID uuid.UUID) (*models.CartItem, error)
	DeleteByCartID(ctx context.Context, cartID uuid.UUID) error
}

type cartItemRepository struct {
	BaseRepository[models.CartItem]
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{
		BaseRepository: NewBaseRepository[models.CartItem](db),
		db:             db,
	}
}

func (r *cartItemRepository) WithTx(tx *gorm.DB) CartItemRepository {
	return NewCartItemRepository(tx)
}

func (r *cartItemRepository) GetByCartAndProduct(ctx context.Context, cartID, productID uuid.UUID) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ? AND is_purchased = ?", cartID, productID, false).
		First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func (r *cartItemRepository) DeleteByCartID(ctx context.Context, cartID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("cart_id = ? AND is_purchased = ?", cartID, false).
		Delete(&models.CartItem{}).Error
}
