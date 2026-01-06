package dto

import "github.com/Nios-V/Go-ecommerce-API/internal/models"

type CreateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=50"`
	Description *string `json:"description" validate:"max=255"`
}

func (r *CreateCategoryRequest) CreateCategoryToModel() *models.Category {
	return &models.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=50"`
	Description *string `json:"description" validate:"omitempty,max=255"`
}

func (r *UpdateCategoryRequest) UpdateCategoryToModel(category *models.Category) {
	if r.Name != nil {
		category.Name = *r.Name
	}
	if r.Description != nil {
		category.Description = r.Description
	}
}
