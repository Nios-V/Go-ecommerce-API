package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type CartService struct {
	*BaseService[models.Cart]
}

func NewCartService(db *gorm.DB, cartRepo repository.CartRepository) *CartService {
	return &CartService{
		BaseService: NewBaseService(cartRepo),
	}
}
