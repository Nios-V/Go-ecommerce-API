package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var ErrCategoryExists = errors.New("category with this name already exists")

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
	caser := cases.Title(language.Spanish)
	category.Name = caser.String(strings.ToLower(category.Name))

	existing, err := s.repo.GetByName(ctx, category.Name)
	if err != nil && existing != nil {
		return ErrCategoryExists
	}

	return s.repo.Create(ctx, category)
}
