package models

import "github.com/google/uuid"

type OrderItem struct {
	Base
	ProductID       uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	OrderID         uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	Quantity        int       `gorm:"not null" json:"quantity"`
	PriceAtPurchase float64   `gorm:"not null" json:"price_at_purchase"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
