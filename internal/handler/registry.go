package handler

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Registry struct {
	Address  *AddressHandler
	Category *CategoryHandler
	Product  *ProductHandler
	User     *UserHandler
	// TODO: Add other handlers
	Validate *validator.Validate
}

func NewRegistry(db *gorm.DB) *Registry {
	v := validator.New()

	// Initialize repositories
	addressRepo := repository.NewAddressRepository(db)
	// cartItemRepo := repository.NewCartItemRepository(db)
	cartRepo := repository.NewCartRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	// orderItemRepo := repository.NewOrderItemRepository(db)
	// orderRepo := repository.NewOrderRepository(db)
	// paymentRepo := repository.NewPaymentRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	// Initialize services
	categoryService := service.NewCategoryService(db, categoryRepo)
	productService := service.NewProductService(db, productRepo)
	userService := service.NewUserService(db, userRepo, cartRepo, roleRepo)
	addressService := service.NewAddressService(db, addressRepo)

	// Initialize handlers
	return &Registry{
		Address:  NewAddressHandler(addressService, v),
		Category: NewCategoryHandler(categoryService, v),
		Product:  NewProductHandler(productService, v),
		User:     NewUserHandler(userService, v),
		Validate: v,
	}
}
