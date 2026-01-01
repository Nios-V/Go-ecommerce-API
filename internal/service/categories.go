package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type CategoryService struct {
	*BaseService[models.Category]
}

func NewCategoryService(db *gorm.DB, categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		BaseService: NewBaseService(categoryRepo),
	}
}
