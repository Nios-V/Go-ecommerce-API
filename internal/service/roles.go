package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type RoleService struct {
	*BaseService[models.Role]
}

func NewRoleService(db *gorm.DB, roleRepo repository.RoleRepository) *RoleService {
	return &RoleService{
		BaseService: NewBaseService(roleRepo),
	}
}
