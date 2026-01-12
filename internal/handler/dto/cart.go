package dto

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
)

type AddCartItemRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required,uuid4"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

func (r *AddCartItemRequest) AddCartItemToModel() *models.CartItem {
	return &models.CartItem{
		ProductID: r.ProductID,
		Quantity:  r.Quantity,
	}
}

type RemoveCartItemRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required,uuid4"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

func (r *RemoveCartItemRequest) RemoveCartItemToModel() *models.CartItem {
	return &models.CartItem{
		ProductID: r.ProductID,
		Quantity:  r.Quantity,
	}
}
