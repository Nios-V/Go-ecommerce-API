package models

import "github.com/google/uuid"

type Product struct {
	Base
	CategoryID  uuid.UUID `gorm:"type:uuid;not null;index" json:"category_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description *string   `gorm:"type:varchar(255)" json:"description,omitempty"`
	Price       float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`

	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}
