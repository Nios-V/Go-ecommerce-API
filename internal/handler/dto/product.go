package dto

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"category_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
}

func ToProductResponse(p *models.Product) *ProductResponse {
	return &ProductResponse{
		ID:          p.ID,
		CategoryID:  p.CategoryID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
	}
}

func ToProductListResponse(products []models.Product) []*ProductResponse {
	res := make([]*ProductResponse, len(products))
	for i, p := range products {
		res[i] = ToProductResponse(&p)
	}
	return res
}

type CreateProductRequest struct {
	CategoryID  uuid.UUID `json:"category_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description *string   `json:"description" validate:"max=255"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Stock       int       `json:"stock" validate:"required,gte=0"`
}

func (r *CreateProductRequest) CreateProductToModel() *models.Product {
	return &models.Product{
		CategoryID:  r.CategoryID,
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Stock:       r.Stock,
	}
}

type UpdateProductRequest struct {
	CategoryID  *uuid.UUID `json:"category_id" validate:"omitempty"`
	Name        *string    `json:"name" validate:"omitempty,min=3,max=100"`
	Description *string    `json:"description" validate:"omitempty,max=255"`
	Price       *float64   `json:"price" validate:"omitempty,gt=0"`
}

func (r *UpdateProductRequest) UpdateProductToModel(product *models.Product) {
	if r.CategoryID != nil {
		product.CategoryID = *r.CategoryID
	}
	if r.Name != nil {
		product.Name = *r.Name
	}
	if r.Description != nil {
		product.Description = r.Description
	}
	if r.Price != nil {
		product.Price = *r.Price
	}
}

type UpdateProductStockRequest struct {
	Stock int `json:"stock" validate:"required,gte=0"`
}

func (r *UpdateProductStockRequest) UpdateProductStockToModel(product *models.Product) {
	product.Stock = r.Stock
}
