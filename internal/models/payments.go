package models

import (
	"errors"
	"time"

	"github.com/Nios-V/Go-ecommerce-API/internal/models/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	OrderID uuid.UUID            `gorm:"type:uuid;not null" json:"order_id"`
	Amount  float64              `gorm:"not null" json:"amount"`
	Method  *enums.PaymentMethod `gorm:"type:varchar(50);not null" json:"method"`
	Status  *enums.PaymentStatus `gorm:"type:varchar(50);not null" json:"status"`
	PaidAt  *time.Time           `json:"paid_at,omitempty"`

	Base
}

func (p *Payment) BeforeSave(tx *gorm.DB) (err error) {

	if p.Status != nil {
		switch *p.Status {
		case enums.PaymentStatusCompleted, enums.PaymentStatusFailed, enums.PaymentStatusRefunded, enums.PaymentStatusPending:

		default:
			return errors.New("invalid payment status")
		}
	}

	if p.Method != nil {
		switch *p.Method {
		case enums.PaymentMethodCreditCard, enums.PaymentMethodPayPal, enums.PaymentMethodBankTransfer, enums.PaymentMethodCashOnDelivery:

		default:
			return errors.New("invalid payment method")
		}
	}

	return nil
}
