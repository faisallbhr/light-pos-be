package service

import (
	"context"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
)

type ProductService interface {
	CreateOpeningStock(ctx context.Context, req *dto.CreateOpeningStockRequest) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateOpeningStock(ctx context.Context, req *dto.CreateOpeningStockRequest) error {
	product := entities.Product{
		Name:      req.Name,
		SKU:       req.SKU,
		Image:     req.Image,
		BuyPrice:  req.BuyPrice,
		SellPrice: req.SellPrice,
		Stock:     req.Stock,
	}

	exists, err := s.productRepo.ExistsBySKU(ctx, req.SKU)
	if err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	if exists {
		return errorsx.NewError(errorsx.ErrConflict, "sku already exists", err)
	}

	if err := s.productRepo.CreateOpeningStock(ctx, &product, req); err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	return nil
}
