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

type CategoryHandler struct {
	categoryService *service.CategoryService
	validate        *validator.Validate
}

func NewCategoryHandler(s *service.CategoryService, v *validator.Validate) *CategoryHandler {
	return &CategoryHandler{
		categoryService: s,
		validate:        v,
	}
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get category details by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID" format(uuid)
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	category, err := h.categoryService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve category")
		return
	}

	res := dto.ToCategoryResponse(category)

	response.JSON(w, http.StatusOK, res)
}

// ListCategories godoc
// @Summary List categories
// @Description Get a paginated list of categories
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} dto.CategoryResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /categories [get]
func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
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

	categories, err := h.categoryService.List(r.Context(), offset, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve categories")
		return
	}

	res := dto.ToCategoryListResponse(categories)

	response.JSON(w, http.StatusOK, res)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Category data"
// @Success 201 {object} dto.CategoryResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /categories [post]
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest

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

	category := *req.CreateCategoryToModel()

	err = h.categoryService.Create(r.Context(), &category)
	if err != nil {
		if errors.Is(err, service.ErrCategoryExists) {
			response.Error(w, http.StatusConflict, "Category with this name already exists")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	response.JSON(w, http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID" format(uuid)
// @Param category body dto.UpdateCategoryRequest true "Updated category data"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /categories/{id} [put]
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	category, err := h.categoryService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Category not found")
		return
	}

	var req dto.UpdateCategoryRequest
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

	req.UpdateCategoryToModel(category)

	err = h.categoryService.Update(r.Context(), category)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update category")
		return
	}

	response.JSON(w, http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID" format(uuid)
// @Success 204
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	IdStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(IdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.categoryService.Delete(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
