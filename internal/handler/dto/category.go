package dto

import "github.com/Nios-V/Go-ecommerce-API/internal/models"

type CreateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=50"`
	Description *string `json:"description" validate:"max=255"`
}

func (r *CreateCategoryRequest) ToModel() *models.Category {
	return &models.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}
