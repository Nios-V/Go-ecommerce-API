package repository

import (
	"context"
	"errors"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	BaseRepository[models.Product]
	UpdateStock(ctx context.Context, productID uuid.UUID, quantity int) error
}

type productRepository struct {
	BaseRepository[models.Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		BaseRepository: NewBaseRepository[models.Product](db),
		db:             db,
	}
}

func (r *productRepository) UpdateStock(ctx context.Context, productID uuid.UUID, quantity int) error {
	result := r.db.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", productID).
		Update("stock", gorm.Expr("stock - ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Insufficient stock or product not found")
	}

	return nil
}
