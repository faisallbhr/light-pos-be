package repository

import (
	"context"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
)

type CategoryRepository interface {
	FindByName(ctx context.Context, params *httpx.QueryParams) ([]*entities.Category, int64, error)
}

type categoryRepository struct {
	BaseRepository *BaseRepository[entities.Category]
	db             *database.DB
}

func NewCategoryRepository(db *database.DB) CategoryRepository {
	return &categoryRepository{
		BaseRepository: NewBaseRepository[entities.Category](db),
		db:             db,
	}
}

func (r *categoryRepository) FindByName(ctx context.Context, params *httpx.QueryParams) ([]*entities.Category, int64, error) {
	var categories []*entities.Category
	var total int64

	query := r.db.WithContext(ctx).Model(&entities.Category{})

	search := params.GetSearch()
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order("name ASC").
		Offset(params.Offset()).
		Limit(params.GetLimit())

	if err := query.Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}
