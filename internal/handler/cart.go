package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Nios-V/Go-ecommerce-API/internal/handler/dto"
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
	var req dto.AddCartItemRequest

	userIdStr := chi.URLParam(r, "user_id")
	user_id, err := uuid.Parse(userIdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if user_id != r.Context().Value("user_id").(uuid.UUID) {
		response.Error(w, http.StatusForbidden, "Access denied")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	err = h.validate.Struct(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	item := req.AddCartItemToModel()

	err = h.cartService.AddItemToCart(r.Context(), id, item)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to add item to cart")
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "Item added to cart"})
}

func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	var req dto.RemoveCartItemRequest

	userIdStr := chi.URLParam(r, "user_id")
	user_id, err := uuid.Parse(userIdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if user_id != r.Context().Value("user_id").(uuid.UUID) {
		response.Error(w, http.StatusForbidden, "Access denied")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	err = h.validate.Struct(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	item := req.RemoveCartItemToModel()
	err = h.cartService.RemoveItemFromCart(r.Context(), id, item)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to remove item from cart")
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "Item removed from cart"})
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

	res := dto.ToCartResponse(cart)

	response.JSON(w, http.StatusOK, res)
}

func (h *CartHandler) UpdateItemQuantity(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCartItemRequest

	userIdStr := chi.URLParam(r, "user_id")
	user_id, err := uuid.Parse(userIdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if user_id != r.Context().Value("user_id").(uuid.UUID) {
		response.Error(w, http.StatusForbidden, "Access denied")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	err = h.validate.Struct(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	item := req.UpdateCartItemToModel()

	err = h.cartService.UpdateItemQuantity(r.Context(), id, item)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update item quantity")
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "Item quantity updated"})
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "user_id")
	user_id, err := uuid.Parse(userIdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if user_id != r.Context().Value("user_id").(uuid.UUID) {
		response.Error(w, http.StatusForbidden, "Access denied")
		return
	}

	err = h.cartService.ClearCart(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to clear cart")
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "Cart cleared"})
}

func (h *CartHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	// Implementation for checking out the cart
}
