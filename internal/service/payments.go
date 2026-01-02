package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type PaymentService struct {
	*BaseService[models.Payment]
}

func NewPaymentService(db *gorm.DB, paymentRepo repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		BaseService: NewBaseService(paymentRepo),
	}
}
