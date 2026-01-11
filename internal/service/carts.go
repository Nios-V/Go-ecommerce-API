package service

import (
	"context"
	"errors"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartService struct {
	*BaseService[models.Cart]
	repo         repository.CartRepository
	cartItemRepo repository.CartItemRepository
}

func NewCartService(db *gorm.DB, cartRepo repository.CartRepository, cartItemRepo repository.CartItemRepository) *CartService {
	return &CartService{
		BaseService:  NewBaseService(cartRepo),
		repo:         cartRepo,
		cartItemRepo: cartItemRepo,
	}
}

func (s *CartService) GetCartByUserID(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	return s.repo.GetCurrentCart(ctx, userID)
}

func (s *CartService) AddItemToCart(ctx context.Context, cartID, productID uuid.UUID, quantity int) error {
	cart, err := s.repo.GetByID(ctx, cartID)
	if err != nil {
		return err
	}

	existingItem, err := s.cartItemRepo.GetByCartAndProduct(ctx, cartID, productID)

	if err == nil {
		existingItem.Quantity += quantity
		return s.cartItemRepo.Update(ctx, existingItem)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newItem := &models.CartItem{
			CartID:      cart.ID,
			ProductID:   productID,
			Quantity:    quantity,
			IsPurchased: false,
		}

		return s.cartItemRepo.Create(ctx, newItem)
	}

	return err
}
