package dto

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
)

type CartResponse struct {
	ID     uuid.UUID          `json:"id"`
	UserID uuid.UUID          `json:"user_id"`
	Items  []CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	ProductID   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity"`
	IsPurchased bool      `json:"is_purchased"`
}

func ToCartResponse(c *models.Cart) *CartResponse {
	cartItems := make([]CartItemResponse, len(c.Items))
	for i, item := range c.Items {
		cartItems[i] = CartItemResponse{
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			IsPurchased: item.IsPurchased,
		}
	}
	return &CartResponse{
		ID:     c.ID,
		UserID: c.UserID,
		Items:  cartItems,
	}
}

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

type UpdateCartItemRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required,uuid4"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

func (r *UpdateCartItemRequest) UpdateCartItemToModel() *models.CartItem {
	return &models.CartItem{
		ProductID: r.ProductID,
		Quantity:  r.Quantity,
	}
}
