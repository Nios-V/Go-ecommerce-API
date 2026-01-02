package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type CartItemsService struct {
	*BaseService[models.CartItem]
}

func NewCartItemService(db *gorm.DB, cartItemRepo repository.CartItemRepository) *CartItemsService {
	return &CartItemsService{
		BaseService: NewBaseService(cartItemRepo),
	}
}
