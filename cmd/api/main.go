package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nios-V/Go-ecommerce-API/internal/config"
	"github.com/Nios-V/Go-ecommerce-API/internal/handler"
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/Nios-V/Go-ecommerce-API/internal/router"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()
	err := cfg.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Address{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
	)
	if err != nil {
		log.Fatal("Migration Error: ", err)
	}

	catRepo := repository.NewCategoryRepository(cfg.DB)
	// TODO: Initialize other repositories

	catService := service.NewCategoryService(cfg.DB, catRepo)
	// TODO: Initialize other services

	r := chi.NewRouter()

	h := &handler.Container{
		Category: handler.NewCategoryHandler(catService),
	}

	router.SetupRoutes(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running in port %s\n", port)
	http.ListenAndServe(":"+port, r)
}
