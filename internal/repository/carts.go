package repository

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	BaseRepository[models.Cart]
	WithTx(tx *gorm.DB) CartRepository
}

type cartRepository struct {
	BaseRepository[models.Cart]
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		BaseRepository: NewBaseRepository[models.Cart](db),
		db:             db,
	}
}

func (r *cartRepository) WithTx(tx *gorm.DB) CartRepository {
	return NewCartRepository(tx)
}
