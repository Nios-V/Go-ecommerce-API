package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type ProductService struct {
	*BaseService[models.Product]
}

func NewProductService(db *gorm.DB, productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		BaseService: NewBaseService(productRepo),
	}
}
