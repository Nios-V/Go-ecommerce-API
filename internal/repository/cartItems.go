package repository

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type CartItemRepository interface {
	BaseRepository[models.CartItem]
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
