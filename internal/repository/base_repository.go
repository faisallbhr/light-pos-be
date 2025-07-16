package repository

import (
	"context"

	"github.com/faisallbhr/light-pos-be/database"
)

type BaseRepository[T any] struct {
	db *database.DB
}

func NewBaseRepository[T any](db *database.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	if err := r.db.WithContext(ctx).Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	if err := r.db.WithContext(ctx).Save(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	if err := r.db.WithContext(ctx).Delete(&entity, id).Error; err != nil {
		return err
	}
	return nil
}
