package service

import (
	"context"

	"github.com/Nios-V/Go-ecommerce-API/internal/repository"
	"github.com/google/uuid"
)

type BaseService[T any] struct {
	repo repository.BaseRepository[T]
}

func NewBaseService[T any](repo repository.BaseRepository[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) Create(ctx context.Context, entity *T) error {
	return s.repo.Create(ctx, entity)
}

func (s *BaseService[T]) GetByID(ctx context.Context, id uuid.UUID) (*T, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BaseService[T]) Update(ctx context.Context, entity *T) error {
	return s.repo.Update(ctx, entity)
}

func (s *BaseService[T]) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *BaseService[T]) List(ctx context.Context, offset, limit int) ([]T, error) {
	return s.repo.List(ctx, offset, limit)
}
