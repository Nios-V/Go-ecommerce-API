package models

import "github.com/google/uuid"

type Cart struct {
	Base
	UserID uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	Items  []CartItem `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}
