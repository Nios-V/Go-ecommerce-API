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

type CheckoutHandler struct {
	checkoutService *service.CheckoutService
	validate        *validator.Validate
}

func NewCheckoutHandler(s *service.CheckoutService, v *validator.Validate) *CheckoutHandler {
	return &CheckoutHandler{
		checkoutService: s,
		validate:        v,
	}
}

func (h *CheckoutHandler) StartCheckout(w http.ResponseWriter, r *http.Request) {
	var req dto.CheckoutRequest

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if id != r.Context().Value("user_id").(uuid.UUID) {
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

	checkout := req.ToModel()

	err = h.checkoutService.StartCheckout(r.Context(), id, checkout.ShippingAddressID, checkout.BillingAddressID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to start checkout: "+err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "Checkout started successfully"})
}

func (h *CheckoutHandler) ConfirmCheckout(w http.ResponseWriter, r *http.Request) {
	var req dto.ConfirmCheckoutRequest

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	orderIdStr := chi.URLParam(r, "order_id")
	orderId, err := uuid.Parse(orderIdStr)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if id != r.Context().Value("user_id").(uuid.UUID) {
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
	paymentMethod := req.PaymentMethod

	err = h.checkoutService.ConfirmCheckout(r.Context(), id, orderId, &paymentMethod)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to confirm checkout: "+err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "Checkout confirmed successfully"})
}
