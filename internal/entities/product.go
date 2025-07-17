package entities

import "time"

type Product struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	SKU        string      `json:"sku"`
	Image      *string     `json:"image,omitempty"`
	Categories []*Category `json:"categories" gorm:"many2many:product_categories"`
	BuyPrice   int         `json:"buy_price"`
	SellPrice  int         `json:"sell_price"`
	Stock      int         `json:"stock"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at,omitempty"`
}
