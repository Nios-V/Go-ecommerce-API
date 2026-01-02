package handler

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"gorm.io/gorm"
)

type Registry struct {
	Category *CategoryHandler
	// TODO: Add other handlers
}

func NewRegistry(db *gorm.DB) *Registry {
	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(db)

	// Initialize services
	categoryService := service.NewCategoryService(db, categoryRepo)

	// Initialize handlers
	return &Registry{
		Category: NewCategoryHandler(categoryService),
	}
}
