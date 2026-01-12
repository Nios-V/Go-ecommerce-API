package handler

import (
	"net/http"

	"github.com/Nios-V/Go-ecommerce-API/internal/response"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CartHandler struct {
	cartService *service.CartService
	validate    *validator.Validate
}

func NewCartHandler(s *service.CartService, v *validator.Validate) *CartHandler {
	return &CartHandler{
		cartService: s,
		validate:    v,
	}
}

func (h *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {

}

func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	// Implementation for removing item from cart
}

func (h *CartHandler) ViewCart(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "user_id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if id != r.Context().Value("user_id").(uuid.UUID) {
		response.Error(w, http.StatusForbidden, "Access denied")
		return
	}

	cart, err := h.cartService.GetCartByUserID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve cart")
		return
	}

	response.JSON(w, http.StatusOK, cart)
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	// Implementation for clearing the cart
}

func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	// Implementation for checking out the cart
}
