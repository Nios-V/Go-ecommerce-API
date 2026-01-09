package router

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Nios-V/Go-ecommerce-API/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(r *chi.Mux, h *handler.Registry) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(MaxBodySize)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		registerCategory(r, h.Category)
		registerProduct(r, h.Product)
		registerUser(r, h.User)
		registerAddress(r, h.Address)
	})
}

func MaxBodySize(next http.Handler) http.Handler {
	maxBytesStr := os.Getenv("MAX_BODY_SIZE")
	maxBytes, err := strconv.ParseInt(maxBytesStr, 10, 64)
	if err != nil {
		maxBytes = 1048576
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
		next.ServeHTTP(w, r)
	})
}

func registerCategory(r chi.Router, h *handler.CategoryHandler) {
	r.Route("/categories", func(r chi.Router) {
		r.Get("/{id}", h.GetByID)
		r.Get("/", h.List)
		r.Post("/", h.Create)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}

func registerProduct(r chi.Router, h *handler.ProductHandler) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/{id}", h.GetByID)
		r.Get("/", h.List)
		r.Post("/", h.Create)
		r.Put("/{id}", h.Update)
		r.Put("/{id}/stock", h.UpdateStock)
		r.Delete("/{id}", h.Delete)
	})
}

func registerUser(r chi.Router, h *handler.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", h.GetByID)
		r.Post("/", h.Create)
	})
}

func registerAddress(r chi.Router, h *handler.AddressHandler) {
	r.Route("/addresses", func(r chi.Router) {
		r.Get("/{id}", h.GetByID)
		r.Get("/", h.List)
		r.Post("/", h.Create)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}
