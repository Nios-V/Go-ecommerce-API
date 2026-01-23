package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductService struct {
	*BaseService[models.Product]
	rdb *redis.Client
}

func NewProductService(db *gorm.DB, rdb *redis.Client, productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		BaseService: NewBaseService(productRepo),
	}
}

func (s *ProductService) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	cacheKey := fmt.Sprintf("product:%s", id.String())

	cachedProduct, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var product models.Product
		if err := json.Unmarshal([]byte(cachedProduct), &product); err == nil {
			return &product, nil
		}
	}

	product, err := s.BaseService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(product)
	if err == nil {
		s.rdb.Set(ctx, cacheKey, productJSON, 10*time.Minute)
	}

	return product, nil
}

func (s *ProductService) Update(ctx context.Context, product *models.Product) error {
	err := s.BaseService.Update(ctx, product)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("product:%s", product.ID.String())
	s.rdb.Del(ctx, cacheKey)
	return nil
}
