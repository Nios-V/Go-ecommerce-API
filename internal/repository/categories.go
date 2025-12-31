package repository

import (
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
