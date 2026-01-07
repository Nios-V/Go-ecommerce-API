package service

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"gorm.io/gorm"
)

type UserService struct {
	*BaseService[models.User]
	repo     repository.UserRepository
	cartRepo repository.CartRepository
}

func NewUserService(db *gorm.DB, userRepo repository.UserRepository, cartRepo repository.CartRepository) *UserService {
	return &UserService{
		BaseService: NewBaseService(userRepo),
		repo:        userRepo,
		cartRepo:    cartRepo,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	// TODO: Asign role
	// TODO: Create cart for user
	err := s.repo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
