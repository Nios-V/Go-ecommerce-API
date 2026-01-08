package service

import (
	"context"
	"fmt"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	*BaseService[models.User]
	repo     repository.UserRepository
	cartRepo repository.CartRepository
	roleRepo repository.RoleRepository
	db       *gorm.DB
}

func NewUserService(db *gorm.DB, userRepo repository.UserRepository, cartRepo repository.CartRepository, roleRepo repository.RoleRepository) *UserService {
	return &UserService{
		BaseService: NewBaseService(userRepo),
		repo:        userRepo,
		cartRepo:    cartRepo,
		roleRepo:    roleRepo,
		db:          db,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txUserRepo := s.repo.WithTx(tx)
		txCartRepo := s.cartRepo.WithTx(tx)
		txRoleRepo := s.roleRepo.WithTx(tx)

		role, err := txRoleRepo.GetByName(ctx, "customer")
		if err != nil {
			return err
		}

		user.Roles = append(user.Roles, *role)

		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPwd)

		if err := txUserRepo.Create(ctx, user); err != nil {
			return err
		}

		newCart := &models.Cart{
			UserID: user.ID,
		}
		if err := txCartRepo.Create(ctx, newCart); err != nil {
			return fmt.Errorf("Failed to create cart: %w", err)
		}

		return nil
	})
}
