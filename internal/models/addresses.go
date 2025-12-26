package models

import "github.com/google/uuid"

type Address struct {
	Base
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Street            string    `gorm:"type:varchar(120);not null" json:"street"`
	City              string    `gorm:"type:varchar(50);not null" json:"city"`
	State             string    `gorm:"type:varchar(50);not null" json:"state"`
	ZipCode           *string   `gorm:"type:varchar(20)" json:"zip_code"`
	Country           string    `gorm:"type:varchar(50);not null" json:"country"`
	IsDefaultShipping bool      `gorm:"not null;default:false" json:"is_default_shipping"`
	IsDefaultBilling  bool      `gorm:"not null;default:false" json:"is_default_billing"`

	User *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}
