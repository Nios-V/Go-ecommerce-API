package service

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckoutService struct {
	cartRepo    repository.CartRepository
	orderRepo   repository.OrderRepository
	paymentRepo repository.PaymentRepository

	db *gorm.DB
}

func NewCheckoutService(cartRepo repository.CartRepository, orderRepo repository.OrderRepository, paymentRepo repository.PaymentRepository, addressRepo repository.AddressRepository) *CheckoutService {
	return &CheckoutService{
		cartRepo:    cartRepo,
		orderRepo:   orderRepo,
		paymentRepo: paymentRepo,
	}
}

func (s *CheckoutService) StartCheckout(ctx context.Context, userID uuid.UUID, shippingAddressID, billingAddressID uuid.UUID) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCartRepo := s.cartRepo.WithTx(tx)
		txOrderRepo := s.orderRepo.WithTx(tx)

		cart, err := txCartRepo.GetCurrentCart(ctx, userID)
		if err != nil {
			return err
		}

		items := cart.Items
		if len(items) == 0 {
			return nil
		}

		var total float64
		orderItems := make([]models.OrderItem, len(items))
		for i, item := range items {
			subtotal := float64(item.Quantity) * item.Product.Price
			total += subtotal

			orderItems[i] = models.OrderItem{
				ProductID:       item.ProductID,
				Quantity:        item.Quantity,
				PriceAtPurchase: item.Product.Price,
			}
		}

		order := &models.Order{
			UserID:            userID,
			Total:             total,
			Items:             orderItems,
			ShippingAddressID: shippingAddressID,
			BillingAddressID:  billingAddressID,
		}

		err = txOrderRepo.Create(ctx, order)
		if err != nil {
			return err
		}

		err = txCartRepo.PurchaseCartItems(ctx, cart.ID)
		if err != nil {
			return err
		}

		return nil
	})
}

// func (s *CheckoutService) ConfirmCheckout(ctx context.Context, userID uuid.UUID, paymentInfo map[string]interface{}) error {

// }
