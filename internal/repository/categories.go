package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	BaseRepository[models.Category]
}

type categoryRepository struct {
	BaseRepository[models.Category]
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		BaseRepository: NewBaseRepository[models.Category](db),
		db:             db,
	}
}

func (r *categoryRepository) GetByName(ctx context.Context, name string) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).
		Where("name = ?", name).
		First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
