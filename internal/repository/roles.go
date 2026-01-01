package repository

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[models.Role]
	GetByName(ctx context.Context, name string) (*models.Role, error)
}

type roleRepository struct {
	BaseRepository[models.Role]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		BaseRepository: NewBaseRepository[models.Role](db),
		db:             db,
	}
}

func (r *roleRepository) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	err := r.db.WithContext(ctx).
		Where("name = ?", name).
		First(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}
