package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateOpeningStock(ctx context.Context, product *entities.Product, openingStock *dto.CreateOpeningStockRequest) error
	FindAll(ctx context.Context, params *httpx.QueryParams) ([]*entities.Product, int64, error)
	FindByID(ctx context.Context, id uint) (*entities.Product, error)
	ExistsBySKU(ctx context.Context, sku string) (bool, error)
	UpdateWithCategories(ctx context.Context, product *entities.Product, categories []string) error
	Delete(ctx context.Context, id uint) error
}

type productRepository struct {
	BaseRepository *BaseRepository[entities.Product]
	db             *database.DB
}

func NewProductRepository(db *database.DB) ProductRepository {
	return &productRepository{
		BaseRepository: NewBaseRepository[entities.Product](db),
		db:             db,
	}
}

func (r *productRepository) Create(ctx context.Context, product *entities.Product) error {
	return r.BaseRepository.Create(ctx, product)
}

func (r *productRepository) CreateOpeningStock(ctx context.Context, product *entities.Product, openingStock *dto.CreateOpeningStockRequest) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			return err
		}
		for _, category := range openingStock.Categories {
			var existingCategory entities.Category
			if err := tx.Where("name = ?", category).FirstOrCreate(&existingCategory, entities.Category{
				Name: category,
			}).Error; err != nil {
				return err
			}

			if err := tx.Model(product).Association("Categories").Append(&existingCategory); err != nil {
				return err
			}
		}

		today := time.Now().Format("20060102")
		prefix := "OS-" + today

		var count int64
		if err := tx.Model(&entities.Purchase{}).Where("invoice_number LIKE ?", prefix+"%").Count(&count).Error; err != nil {
			return err
		}

		runningNumber := fmt.Sprintf("%04d", count+1)
		invoiceNumber := prefix + "-" + runningNumber

		purchase := &entities.Purchase{
			InvoiceNumber: invoiceNumber,
			Type:          "opening_stock",
			SupplierID:    1,
			PurchaseDate:  time.Now(),
		}

		if err := tx.Create(purchase).Error; err != nil {
			return err
		}

		purchaseItem := &entities.PurchaseItem{
			PurchaseID:   purchase.ID,
			ProductID:    product.ID,
			Quantity:     openingStock.Stock,
			BuyPrice:     openingStock.BuyPrice,
			TotalPrice:   openingStock.Stock * openingStock.BuyPrice,
			RemainingQty: openingStock.Stock,
		}

		if err := tx.Create(purchaseItem).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *productRepository) FindAll(ctx context.Context, params *httpx.QueryParams) ([]*entities.Product, int64, error) {
	var products []*entities.Product
	var total int64

	query := r.db.WithContext(ctx).Preload("Categories").Model(&entities.Product{})

	search := params.GetSearch()
	if search != "" {
		query = query.
			Joins("JOIN product_categories ON product_categories.product_id = products.id").
			Joins("JOIN categories ON categories.id = product_categories.category_id").
			Where("products.name LIKE ? OR products.sku LIKE ? OR categories.name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order(params.GetOrderBy() + " " + params.GetSort()).
		Offset(params.Offset()).
		Limit(params.GetLimit())

	if err := query.Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.WithContext(ctx).Preload("Categories").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entities.Product{}).Where("sku = ?", sku).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *productRepository) UpdateWithCategories(ctx context.Context, product *entities.Product, categories []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(product).Error; err != nil {
			return err
		}

		if err := tx.Model(product).Association("Categories").Clear(); err != nil {
			return err
		}

		for _, category := range categories {
			var existingCategory entities.Category

			if err := tx.Where("name = ?", category).FirstOrCreate(&existingCategory, entities.Category{
				Name: category,
			}).Error; err != nil {
				return err
			}

			if err := tx.Model(product).Association("Categories").Append(&existingCategory); err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
	return r.BaseRepository.Delete(ctx, id)
}
