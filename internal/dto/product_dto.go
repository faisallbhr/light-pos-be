package dto

type CreateOpeningStockRequest struct {
	Name       string   `json:"name" binding:"required"`
	SKU        string   `json:"sku" binding:"required"`
	Image      *string  `json:"image,omitempty"`
	Categories []string `json:"categories" binding:"required"`
	BuyPrice   int      `json:"buy_price" binding:"required,min=1"`
	SellPrice  int      `json:"sell_price" binding:"required,min=1"`
	Stock      int      `json:"stock" binding:"required"`
}
