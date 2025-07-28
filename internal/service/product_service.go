package service

import (
	"context"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/internal/service/mapper"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"github.com/faisallbhr/light-pos-be/pkg/utils"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateOpeningStock(ctx context.Context, req *dto.CreateOpeningStockRequest, imageFile *multipart.FileHeader) (*dto.ProductResponse, error)
	GetProducts(ctx context.Context, params *httpx.QueryParams) ([]*dto.ProductResponse, int64, error)
	GetProductByID(ctx context.Context, id uint) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error)
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

func (s *productService) CreateOpeningStock(ctx context.Context, req *dto.CreateOpeningStockRequest, imageFile *multipart.FileHeader) (*dto.ProductResponse, error) {
	var imageFileURL string

	if imageFile != nil {
		ext := strings.ToLower(filepath.Ext(imageFile.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return nil, errorsx.NewError(errorsx.ErrBadRequest, "invalid image format, only jpg, jpeg, and png are allowed", nil)
		}

		dstFolder := filepath.Join("storage", "products")
		fileURL, err := utils.SaveUploadedFile(imageFile, dstFolder)
		if err != nil {
			return nil, errorsx.NewError(errorsx.ErrInternal, "failed to save image", err)
		}
		imageFileURL = fileURL
	}

	product := &entities.Product{
		Name:      req.Name,
		SKU:       req.SKU,
		BuyPrice:  req.BuyPrice,
		SellPrice: req.SellPrice,
		Stock:     req.Stock,
	}

	if imageFileURL != "" {
		product.Image = &imageFileURL
	}

	exists, err := s.productRepo.ExistsBySKU(ctx, req.SKU)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	if exists {
		return nil, errorsx.NewError(errorsx.ErrConflict, "sku already exists", err)
	}

	if err := s.productRepo.CreateOpeningStock(ctx, product, req); err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	return mapper.ToProductResponse(product), nil
}

func (s *productService) GetProducts(ctx context.Context, params *httpx.QueryParams) ([]*dto.ProductResponse, int64, error) {
	products, total, err := s.productRepo.FindAll(ctx, params)
	if err != nil {
		return nil, 0, errorsx.NewError(errorsx.ErrInternal, "failed to fetch products", err)
	}
	res := mapper.ToProductsResponse(products)
	return res, total, nil
}

func (s *productService) GetProductByID(ctx context.Context, id uint) (*dto.ProductResponse, error) {
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

func (s *productService) UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrNotFound, "product not found", nil)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "failed to check product existence", err)
	}

	if req.SKU != "" {
		skuExists, err := s.productRepo.ExistsBySKU(ctx, req.SKU)
		if err != nil {
			return nil, errorsx.NewError(errorsx.ErrInternal, "failed to check sku existence", err)
		}
		if skuExists && product.SKU != req.SKU {
			return nil, errorsx.NewError(errorsx.ErrConflict, "sku already exists", nil)
		}
	}

	product.Name = req.Name
	product.SKU = req.SKU
	product.SellPrice = req.SellPrice

	if req.Image != nil {
		ext := strings.ToLower(filepath.Ext(req.Image.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return nil, errorsx.NewError(errorsx.ErrBadRequest, "invalid image format, only jpg, jpeg, and png are allowed", nil)
		}

		dstFolder := filepath.Join("storage", "products")
		fileURL, err := utils.SaveUploadedFile(req.Image, dstFolder)
		if err != nil {
			return nil, errorsx.NewError(errorsx.ErrInternal, "failed to save image", err)
		}

		if product.Image != nil {
			_ = os.Remove(*product.Image)
		}

		product.Image = &fileURL
	}

	if err := s.productRepo.UpdateWithCategories(ctx, product, req.Categories); err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "failed to update product", err)
	}

	res := mapper.ToProductResponse(product)
	return res, nil
}

func (s *productService) DeleteProduct(ctx context.Context, id uint) error {
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsx.NewError(errorsx.ErrNotFound, "product not found", nil)
		}
		return errorsx.NewError(errorsx.ErrInternal, "failed to check product existence", err)
	}

	if product.Image != nil {
		if err := os.Remove(*product.Image); err != nil && !os.IsNotExist(err) {
			return errorsx.NewError(errorsx.ErrInternal, "failed to delete product image", err)
		}
	}

	if err := s.productRepo.Delete(ctx, id); err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "failed to delete product", err)
	}
	return nil
}
