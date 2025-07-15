package repository

import (
	"context"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
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

func (r *BaseRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
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

func (r *BaseRepository[T]) FindAll(ctx context.Context, params *httpx.QueryParams, searchFields []string) ([]*T, int64, error) {
	var entities []*T
	query, total, err := httpx.ApplyMetaQuery(r.db.WithContext(ctx), new(T), params, searchFields)
	if err != nil {
		return nil, 0, err
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return entities, total, nil
}
