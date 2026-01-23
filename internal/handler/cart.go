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

// AddItem godoc
// @Summary Add item to cart
// @Description Add an item to the user's cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param user_id path string true "User ID" format(uuid)
// @Param id path string true "Cart ID" format(uuid)
// @Param item body dto.AddCartItemRequest true "Item to add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{user_id}/cart/{id}/items [post]
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

// RemoveItem godoc
// @Summary Remove item from cart
// @Description Remove an item from the user's cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param user_id path string true "User ID" format(uuid)
// @Param id path string true "Cart ID" format(uuid)
// @Param item body dto.RemoveCartItemRequest true "Item to remove"
// @Success 200 {object} map[string]string
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{user_id}/cart/{id}/items [delete]
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

// ViewCart godoc
// @Summary View cart
// @Description View the user's cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param user_id path string true "User ID" format(uuid)
// @Success 200 {object} dto.CartResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{user_id}/cart [get]
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

// UpdateItemQuantity godoc
// @Summary Update item quantity in cart
// @Description Update the quantity of an item in the user's cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param user_id path string true "User ID" format(uuid)
// @Param id path string true "Cart ID" format(uuid)
// @Param item body dto.UpdateCartItemRequest true "Item quantity to update"
// @Success 200 {object} map[string]string
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{user_id}/cart/{id}/items [put]
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

// ClearCart godoc
// @Summary Clear cart
// @Description Clear all items from the user's cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param user_id path string true "User ID" format(uuid)
// @Param id path string true "Cart ID" format(uuid)
// @Success 200 {object} map[string]string
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{user_id}/cart/{id}/ [delete]
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
