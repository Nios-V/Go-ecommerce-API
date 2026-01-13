package dto

import "github.com/google/uuid"

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
