package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository interface {
	BaseRepository[models.Address]
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]models.Address, error)
	GetDefaultShippingByUserID(ctx context.Context, userID uuid.UUID) (*models.Address, error)
	GetDefaultBillingByUserID(ctx context.Context, userID uuid.UUID) (*models.Address, error)
}

type addressRepository struct {
	BaseRepository[models.Address]
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		BaseRepository: NewBaseRepository[models.Address](db),
		db:             db,
	}
}

func (r *addressRepository) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&addresses).Error

	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *addressRepository) GetDefaultShippingByUserID(ctx context.Context, userID uuid.UUID) (*models.Address, error) {
	var address models.Address
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND is_default_shipping = ?", userID, true).
		First(&address).Error

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *addressRepository) GetDefaultBillingByUserID(ctx context.Context, userID uuid.UUID) (*models.Address, error) {
	var address models.Address
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND is_default_billing = ?", userID, true).
		First(&address).Error

	if err != nil {
		return nil, err
	}

	return &address, nil
}
