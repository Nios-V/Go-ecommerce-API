package router

import (
	"net/http"

	"github.com/Nios-V/Go-ecommerce-API/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(r *chi.Mux, h *handler.Registry) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		registerCategory(r, h.Category)
	})
}

func registerCategory(r chi.Router, h *handler.CategoryHandler) {
	r.Route("/categories", func(r chi.Router) {
		r.Get("/{id}", h.GetByID)
	})
}
