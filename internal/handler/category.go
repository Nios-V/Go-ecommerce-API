package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/response"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: s,
	}
}

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

	response.JSON(w, http.StatusOK, category)
}

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

	response.JSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	if category.Name == "" {
		response.Error(w, http.StatusBadRequest, "Category name needed")
		return
	}

	if len(category.Name) < 3 {
		response.Error(w, http.StatusBadRequest, "Category name must be longer than 3 characters")
		return
	}

	err = h.categoryService.Create(r.Context(), &category)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	response.JSON(w, http.StatusCreated, category)
}
