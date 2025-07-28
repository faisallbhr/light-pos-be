package mapper

import (
	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func ToCategoriesResponse(categories []*entities.Category) []*dto.CategoryResponse {
	res := make([]*dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		res = append(res, &dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return res
}
