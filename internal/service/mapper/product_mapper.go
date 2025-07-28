package mapper

import (
	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func ToProductResponse(product *entities.Product) *dto.ProductResponse {
	var categories []string
	for _, c := range product.Categories {
		categories = append(categories, c.Name)
	}

	return &dto.ProductResponse{
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

func ToProductsResponse(products []*entities.Product) []*dto.ProductResponse {
	res := make([]*dto.ProductResponse, 0, len(products))
	for _, product := range products {
		res = append(res, ToProductResponse(product))
	}
	return res
}
