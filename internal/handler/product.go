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

type ProductHandler struct {
	productService *service.ProductService
	validate       *validator.Validate
}

func NewProductHandler(s *service.ProductService, v *validator.Validate) *ProductHandler {
	return &ProductHandler{
		productService: s,
		validate:       v,
	}
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get product details by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	product, err := h.productService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve product")
		return
	}

	res := dto.ToProductResponse(product)
	response.JSON(w, http.StatusOK, res)
}

// ListProducts godoc
// @Summary List products
// @Description Get a paginated list of products
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} dto.ProductResponse
// @Failure 500 {object} response.StandardResponse
// @Router /products [get]
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
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

	products, err := h.productService.List(r.Context(), offset, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	res := dto.ToProductListResponse(products)
	response.JSON(w, http.StatusOK, res)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product data"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest

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

	product := *req.CreateProductToModel()

	err = h.productService.Create(r.Context(), &product)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	response.JSON(w, http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update an existing product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Param product body dto.UpdateProductRequest true "Product data"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	product, err := h.productService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Product not found")
		return
	}

	var req dto.UpdateProductRequest
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

	req.UpdateProductToModel(product)

	err = h.productService.Update(r.Context(), product)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	response.JSON(w, http.StatusOK, product)
}

// UpdateProductStock godoc
// @Summary Update product stock
// @Description Update the stock quantity of an existing product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Param stock body dto.UpdateProductStockRequest true "Updated stock data"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /products/{id}/stock [put]
func (h *ProductHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	product, err := h.productService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Product not found")
		return
	}

	var req dto.UpdateProductStockRequest
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

	req.UpdateProductStockToModel(product)

	err = h.productService.Update(r.Context(), product)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	response.JSON(w, http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete an existing product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Success 204 {object} nil
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	IdStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(IdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.productService.Delete(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
