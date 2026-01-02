package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type AddressService struct {
	*BaseService[models.Address]
}

func NewAddressRepository(db *gorm.DB, addressRepo repository.AddressRepository) *AddressService {
	return &AddressService{
		BaseService: NewBaseService(addressRepo),
	}
}
