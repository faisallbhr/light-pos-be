package service

import (
	"context"
	"errors"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/internal/service/mapper"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateOpeningStock(ctx context.Context, req *dto.CreateOpeningStockRequest) error
	GetProducts(ctx context.Context, params *httpx.QueryParams) ([]*dto.UpdateProductResponse, int64, error)
	GetProductByID(ctx context.Context, id uint) (*dto.UpdateProductResponse, error)
	UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, id uint) error
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

func (s *productService) GetProducts(ctx context.Context, params *httpx.QueryParams) ([]*dto.UpdateProductResponse, int64, error) {
	products, total, err := s.productRepo.FindAll(ctx, params)
	if err != nil {
		return nil, 0, errorsx.NewError(errorsx.ErrInternal, "failed to fetch products", err)
	}
	res := mapper.ToProductResponses(products)
	return res, total, nil
}

func (s *productService) GetProductByID(ctx context.Context, id uint) (*dto.UpdateProductResponse, error) {
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrNotFound, "product not found", err)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "failed to fetch product", err)
	}
	res := mapper.ToProductResponse(product)
	return res, nil
}

func (s *productService) UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.UpdateProductResponse, error) {
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrNotFound, "product not found", nil)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "failed to check product existence", err)
	}

	if req.SKU != "" {
		skuExists, err := s.productRepo.ExistsBySKU(ctx, product.SKU)
		if err != nil {
			return nil, errorsx.NewError(errorsx.ErrInternal, "failed to check sku existence", err)
		}
		if skuExists && req.SKU != product.SKU {
			return nil, errorsx.NewError(errorsx.ErrConflict, "sku already exists", nil)
		}
	}
	product.Name = req.Name
	product.SKU = req.SKU
	product.Image = req.Image
	product.SellPrice = req.SellPrice

	if err := s.productRepo.UpdateWithCategories(ctx, product, req.Categories); err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "failed to update product", err)
	}

	res := mapper.ToProductResponse(product)
	return res, nil
}

func (s *productService) DeleteProduct(ctx context.Context, id uint) error {
	if err := s.productRepo.Delete(ctx, id); err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "failed to delete product", err)
	}
	return nil
}
