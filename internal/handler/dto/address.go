package dto

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
)

type AddressResponse struct {
	ID                uuid.UUID `json:"id"`
	UserID            uuid.UUID `json:"user_id"`
	Street            string    `json:"street"`
	City              string    `json:"city"`
	State             string    `json:"state"`
	ZipCode           *string   `json:"zip_code,omitempty"`
	Country           string    `json:"country"`
	IsDefaultShipping bool      `json:"is_default_shipping"`
	IsDefaultBilling  bool      `json:"is_default_billing"`
}

func ToAddressResponse(address *models.Address) AddressResponse {
	return AddressResponse{
		ID:                address.ID,
		UserID:            address.UserID,
		Street:            address.Street,
		City:              address.City,
		State:             address.State,
		ZipCode:           address.ZipCode,
		Country:           address.Country,
		IsDefaultShipping: address.IsDefaultShipping,
		IsDefaultBilling:  address.IsDefaultBilling,
	}
}

func ToAddressListResponse(address []models.Address) []AddressResponse {
	res := make([]AddressResponse, len(address))
	for i, a := range address {
		res[i] = ToAddressResponse(&a)
	}
	return res
}

type CreateAddressRequest struct {
	UserID            uuid.UUID `json:"user_id" validate:"required"`
	Street            string    `json:"street" validate:"required,min=5,max=120"`
	City              string    `json:"city" validate:"required,min=2,max=50"`
	State             string    `json:"state" validate:"required,min=2,max=50"`
	ZipCode           *string   `json:"zip_code" validate:"omitempty,max=20"`
	Country           string    `json:"country" validate:"required,min=2,max=50"`
	IsDefaultShipping bool      `json:"is_default_shipping"`
	IsDefaultBilling  bool      `json:"is_default_billing"`
}

func (r *CreateAddressRequest) CreateAddressToModel() *models.Address {
	return &models.Address{
		UserID:            r.UserID,
		Street:            r.Street,
		City:              r.City,
		State:             r.State,
		ZipCode:           r.ZipCode,
		Country:           r.Country,
		IsDefaultShipping: r.IsDefaultShipping,
		IsDefaultBilling:  r.IsDefaultBilling,
	}
}

type UpdateAddressRequest struct {
	Street            *string `json:"street" validate:"omitempty,min=5,max=120"`
	City              *string `json:"city" validate:"omitempty,min=2,max=50"`
	State             *string `json:"state" validate:"omitempty,min=2,max=50"`
	ZipCode           *string `json:"zip_code" validate:"omitempty,max=20"`
	Country           *string `json:"country" validate:"omitempty,min=2,max=50"`
	IsDefaultShipping *bool   `json:"is_default_shipping"`
	IsDefaultBilling  *bool   `json:"is_default_billing"`
}

func (r *UpdateAddressRequest) UpdateAddressToModel(address *models.Address) {
	if r.Street != nil {
		address.Street = *r.Street
	}
	if r.City != nil {
		address.City = *r.City
	}
	if r.State != nil {
		address.State = *r.State
	}
	if r.ZipCode != nil {
		address.ZipCode = r.ZipCode
	}
	if r.Country != nil {
		address.Country = *r.Country
	}
	if r.IsDefaultShipping != nil {
		address.IsDefaultShipping = *r.IsDefaultShipping
	}
	if r.IsDefaultBilling != nil {
		address.IsDefaultBilling = *r.IsDefaultBilling
	}
}
