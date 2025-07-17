package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateOpeningStock(ctx context.Context, product *entities.Product, openingStock *dto.CreateOpeningStockRequest) error
	ExistsBySKU(ctx context.Context, sku string) (bool, error)
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
			if err := tx.Where("name = ?", category).FirstOrCreate(&existingCategory, category).Error; err != nil {
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

func (r *productRepository) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entities.Product{}).Where("sku = ?", sku).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
