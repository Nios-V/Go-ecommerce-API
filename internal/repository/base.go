package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	List(ctx context.Context, limit, offset int) ([]T, error)
	GetByID(ctx context.Context, id uuid.UUID) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *baseRepository[T]) List(ctx context.Context, limit, offset int) ([]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&entities).Error
	return entities, err
}

func (r *baseRepository[T]) GetByID(ctx context.Context, id uuid.UUID) (*T, error) {
	var entity T

	err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *baseRepository[T]) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(new(T), "id = ?", id).Error
}
