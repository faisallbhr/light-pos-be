package dto

import "mime/multipart"

type CreateOpeningStockRequest struct {
	Name       string                `form:"name" binding:"required"`
	SKU        string                `form:"sku" binding:"required"`
	Image      *multipart.FileHeader `form:"image"`
	Categories []string              `form:"categories[]" binding:"required"`
	BuyPrice   int                   `form:"buy_price" binding:"required,min=1"`
	SellPrice  int                   `form:"sell_price" binding:"required,min=1"`
	Stock      int                   `form:"stock" binding:"required"`
}

type UpdateProductRequest struct {
	Name       string                `form:"name" binding:"required"`
	SKU        string                `form:"sku" binding:"required"`
	Image      *multipart.FileHeader `form:"image"`
	Categories []string              `form:"categories[]" binding:"required"`
	SellPrice  int                   `form:"sell_price" binding:"required,min=1"`
}

type ProductResponse struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	SKU        string   `json:"sku"`
	Image      *string  `json:"image"`
	Categories []string `json:"categories"`
	BuyPrice   int      `json:"buy_price"`
	SellPrice  int      `json:"sell_price"`
	Stock      int      `json:"stock"`
}
