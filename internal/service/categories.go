package service

import (
	"context"
	"errors"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type CategoryService struct {
	*BaseService[models.Category]
	repo repository.CategoryRepository
}

func NewCategoryService(db *gorm.DB, categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		BaseService: NewBaseService(categoryRepo),
		repo:        categoryRepo,
	}
}

func (s *CategoryService) Create(ctx context.Context, category *models.Category) error {
	existing, err := s.repo.GetByName(ctx, category.Name)
	if err != nil && existing != nil {
		return errors.New("category with this name already exists")
	}

	return s.repo.Create(ctx, category)
}
