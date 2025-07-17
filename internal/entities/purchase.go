package entities

import "time"

type Purchase struct {
	ID            uint      `json:"id"`
	InvoiceNumber string    `json:"invoice_number"`
	SupplierID    uint      `json:"supplier_id"`
	Supplier      Supplier  `gorm:"foreignKey:SupplierID"`
	Type          string    `json:"type"`
	PurchaseDate  time.Time `json:"purchase_date"`
	TotalPrice    int       `json:"total_price"`
}

type PurchaseItem struct {
	ID           uint     `json:"id"`
	PurchaseID   uint     `json:"purchase_id"`
	Purchase     Purchase `gorm:"foreignKey:PurchaseID"`
	ProductID    uint     `json:"product_id"`
	Product      Product  `gorm:"foreignKey:ProductID"`
	Quantity     int      `json:"quantity"`
	BuyPrice     int      `json:"buy_price"`
	TotalPrice   int      `json:"total_price"`
	RemainingQty int      `json:"remaining_quantity" gorm:"column:remaining_quantity"`
}
