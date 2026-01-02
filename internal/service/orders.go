package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type OrderService struct {
	*BaseService[models.Order]
}

func NewOrderService(db *gorm.DB, orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{
		BaseService: NewBaseService(orderRepo),
	}
}
