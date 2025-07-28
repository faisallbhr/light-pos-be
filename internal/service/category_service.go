package service

import (
	"context"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/internal/service/mapper"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
)

type CategoryService interface {
	GetCategories(ctx context.Context, params *httpx.QueryParams) ([]*dto.CategoryResponse, int64, error)
}
type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) GetCategories(ctx context.Context, params *httpx.QueryParams) ([]*dto.CategoryResponse, int64, error) {
	categories, total, err := s.categoryRepo.FindByName(ctx, params)
	if err != nil {
		return nil, 0, errorsx.NewError(errorsx.ErrInternal, "failed to find category", err)
	}

	res := mapper.ToCategoriesResponse(categories)
	return res, total, nil
}
