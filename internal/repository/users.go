package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User]
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[models.User](db),
		db:             db,
	}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
