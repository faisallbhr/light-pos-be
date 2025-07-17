package entities

import "time"

type Category struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Products  []*Product `gorm:"many2many:product_categories"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
