package repository

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	BaseRepository[models.OrderItem]
	WithTx(tx *gorm.DB) OrderItemRepository
}

type orderItemRepository struct {
	BaseRepository[models.OrderItem]
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{
		BaseRepository: NewBaseRepository[models.OrderItem](db),
		db:             db,
	}
}

func (r *orderItemRepository) WithTx(tx *gorm.DB) OrderItemRepository {
	return NewOrderItemRepository(tx)
}
