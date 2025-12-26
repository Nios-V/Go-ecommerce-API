package models

import "github.com/google/uuid"

type CartItem struct {
	CartID    uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"cart_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`

	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
}
