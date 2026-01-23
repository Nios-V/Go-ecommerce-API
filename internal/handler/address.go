package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Nios-V/Go-ecommerce-API/internal/handler/dto"
	"github.com/Nios-V/Go-ecommerce-API/internal/response"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AddressHandler struct {
	addressService *service.AddressService
	validate       *validator.Validate
}

func NewAddressHandler(s *service.AddressService, v *validator.Validate) *AddressHandler {
	return &AddressHandler{
		addressService: s,
		validate:       v,
	}
}

// GetAddressByID godoc
// @Summary Get address by ID
// @Description Get address details by ID
// @Tags Addresses
// @Accept json
// @Produce json
// @Param id path string true "Address ID" format(uuid)
// @Success 200 {object} dto.AddressResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /addresses/{id} [get]
func (h *AddressHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	address, err := h.addressService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve address")
		return
	}

	res := dto.ToAddressResponse(address)
	response.JSON(w, http.StatusOK, res)
}

// ListAddresses godoc
// @Summary List addresses
// @Description Get a paginated list of addresses
// @Tags Addresses
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} dto.AddressResponse
// @Failure 500 {object} response.StandardResponse
// @Router /addresses [get]
func (h *AddressHandler) List(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	addresses, err := h.addressService.List(r.Context(), offset, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve addresses")
		return
	}

	res := dto.ToAddressListResponse(addresses)
	response.JSON(w, http.StatusOK, res)
}

// CreateAddress godoc
// @Summary Create a new address
// @Description Create a new address
// @Tags Addresses
// @Accept json
// @Produce json
// @Param address body dto.CreateAddressRequest true "Address data"
// @Success 201 {object} dto.AddressResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /addresses [post]
func (h *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAddressRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	if err := h.validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation Error: "+err.Error())
		return
	}

	address := *req.CreateAddressToModel()

	err = h.addressService.Create(r.Context(), &address)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	response.JSON(w, http.StatusCreated, address)
}

// UpdateAddress godoc
// @Summary Update an address
// @Description Update an existing address by ID
// @Tags Addresses
// @Accept json
// @Produce json
// @Param id path string true "Address ID" format(uuid)
// @Param address body dto.UpdateAddressRequest true "Updated address data"
// @Success 200 {object} dto.AddressResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /addresses/{id} [put]
func (h *AddressHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	address, err := h.addressService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Address not found")
		return
	}

	var req dto.UpdateAddressRequest
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
	if err := h.validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation Error: "+err.Error())
		return
	}

	req.UpdateAddressToModel(address)

	err = h.addressService.Update(r.Context(), address)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update address")
		return
	}

	response.JSON(w, http.StatusOK, address)
}

// DeleteAddress godoc
// @Summary Delete an address
// @Description Delete an address by ID
// @Tags Addresses
// @Accept json
// @Produce json
// @Param id path string true "Address ID" format(uuid)
// @Success 204 "No Content"
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /addresses/{id} [delete]
func (h *AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	IdStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(IdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.addressService.Delete(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete address")
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
