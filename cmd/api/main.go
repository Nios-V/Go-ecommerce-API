package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nios-V/Go-ecommerce-API/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.LoadConfig()
	err := cfg.DB.AutoMigrate()
	if err != nil {
		log.Fatal("Migration Error: ", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running in port %s\n", port)
	http.ListenAndServe(":"+port, r)
}
