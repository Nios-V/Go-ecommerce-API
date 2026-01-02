package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type OrderItemService struct {
	*BaseService[models.OrderItem]
}

func NewOrderItemService(db *gorm.DB, orderItemRepo repository.OrderItemRepository) *OrderItemService {
	return &OrderItemService{
		BaseService: NewBaseService(orderItemRepo),
	}
}
