package service

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type UserService struct {
	*BaseService[models.User]
}

func NewUserService(db *gorm.DB, userRepo repository.UserRepository) *UserService {
	return &UserService{
		BaseService: NewBaseService(userRepo),
	}
}
