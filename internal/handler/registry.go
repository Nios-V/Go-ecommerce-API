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
	addressRepo := repository.NewAddressRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	cartRepo := repository.NewCartRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	orderItemRepo := repository.NewOrderItemRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	// Initialize services
	categoryService := service.NewCategoryService(db, categoryRepo)

	// Initialize handlers
	return &Registry{
		Category: NewCategoryHandler(categoryService),
	}
}
