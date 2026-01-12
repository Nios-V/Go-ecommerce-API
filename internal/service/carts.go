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
	productRepo  repository.ProductRepository
}

func NewCartService(db *gorm.DB, cartRepo repository.CartRepository, cartItemRepo repository.CartItemRepository, productRepo repository.ProductRepository) *CartService {
	return &CartService{
		BaseService:  NewBaseService(cartRepo),
		repo:         cartRepo,
		cartItemRepo: cartItemRepo,
		productRepo:  productRepo,
	}
}

func (s *CartService) GetCartByUserID(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	return s.repo.GetCurrentCart(ctx, userID)
}

func (s *CartService) AddItemToCart(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error {
	existingItem, err := s.cartItemRepo.GetByCartAndProduct(ctx, cartID, item.ProductID)

	if err == nil {
		existingItem.Quantity += item.Quantity
		return s.cartItemRepo.Update(ctx, existingItem)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newItem := &models.CartItem{
			CartID:      cartID,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			IsPurchased: false,
		}

		return s.cartItemRepo.Create(ctx, newItem)
	}

	return err
}

func (s *CartService) RemoveItemFromCart(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error {
	existingItem, err := s.cartItemRepo.GetByCartAndProduct(ctx, cartID, item.ProductID)
	if err != nil {
		return err
	}

	err = s.cartItemRepo.Delete(ctx, existingItem.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) UpdateItemQuantity(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error {
	existingItem, err := s.cartItemRepo.GetByCartAndProduct(ctx, cartID, item.ProductID)
	if err != nil {
		return err
	}

	product, err := s.productRepo.GetByID(ctx, item.ProductID)
	if err != nil {
		return err
	}

	if item.Quantity > product.Stock {
		return errors.New("insufficient stock for the requested quantity")
	}

	existingItem.Quantity = item.Quantity
	return s.cartItemRepo.Update(ctx, existingItem)
}
