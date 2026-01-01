package repository

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	BaseRepository[models.OrderItem]
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
