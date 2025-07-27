package mapper

import (
	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func ToProductResponse(product *entities.Product) *dto.UpdateProductResponse {
	var categories []string
	for _, c := range product.Categories {
		categories = append(categories, c.Name)
	}

	return &dto.UpdateProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		SKU:        product.SKU,
		Image:      product.Image,
		Categories: categories,
		BuyPrice:   product.BuyPrice,
		SellPrice:  product.SellPrice,
		Stock:      product.Stock,
	}
}

func ToProductResponses(products []*entities.Product) []*dto.UpdateProductResponse {
	res := make([]*dto.UpdateProductResponse, 0, len(products))
	for _, product := range products {
		res = append(res, ToProductResponse(product))
	}
	return res
}
