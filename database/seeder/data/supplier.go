package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedSuppliers(db *database.DB) {
	log.Println("seeding suppliers table...")

	var count int64
	db.Model(&entities.Supplier{}).Count(&count)
	if count > 0 {
		log.Println("suppliers already seeded")
		return
	}

	suppliers := []entities.Supplier{
		{
			Name:    "OPENING_STOCK",
			Phone:   nil,
			Address: nil,
		},
	}

	if err := db.Create(&suppliers).Error; err != nil {
		log.Fatalf("failed to seed suppliers: %v", err)
	}

	log.Println("seeded suppliers table")
}
