package dto

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models/enums"
	"github.com/google/uuid"
)

type CheckoutRequest struct {
	ShippingAddressID uuid.UUID `json:"shipping_address_id" validate:"required,uuid4"`
	BillingAddressID  uuid.UUID `json:"billing_address_id" validate:"required,uuid4"`
}

func (r *CheckoutRequest) ToModel() *CheckoutRequest {
	return &CheckoutRequest{
		ShippingAddressID: r.ShippingAddressID,
		BillingAddressID:  r.BillingAddressID,
	}
}

type ConfirmCheckoutRequest struct {
	PaymentMethod enums.PaymentMethod `json:"payment_method" validate:"required,oneof=credit_card paypal bank_transfer"`
}

func (r *ConfirmCheckoutRequest) ToModel() *ConfirmCheckoutRequest {
	return &ConfirmCheckoutRequest{
		PaymentMethod: r.PaymentMethod,
	}
}
