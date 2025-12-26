package models

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	UserID            uuid.UUID          `gorm:"type:uuid;not null" json:"user_id"`
	Total             float64            `gorm:"type:decimal(10,2);not null" json:"total"`
	Status            *enums.OrderStatus `gorm:"type:varchar(50)" json:"status"`
	ShippingAddressID uuid.UUID          `gorm:"type:uuid;not null;index" json:"shipping_address_id"`
	BillingAddressID  uuid.UUID          `gorm:"type:uuid;not null;index" json:"billing_address_id"`

	ShippingAddress Address     `gorm:"foreignKey:ShippingAddressID;references:ID" json:"shipping_address"`
	BillingAddress  Address     `gorm:"foreignKey:BillingAddressID;references:ID" json:"billing_address"`
	Items           []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	Payment         Payment     `gorm:"foreignKey:OrderID" json:"payment"`
	User            User        `gorm:"foreignKey:UserID;references:ID" json:"user"`

	Base
}

func (o *Order) BeforeSave(tx *gorm.DB) (err error) {
	if o.Status == nil {
		return nil
	}

	switch *o.Status {
	case enums.OrderStatusPending, enums.OrderStatusPaid, enums.OrderStatusShipped, enums.OrderStatusDelivered, enums.OrderStatusCancelled:

	default:
		return gorm.ErrInvalidData
	}
	return nil
}
