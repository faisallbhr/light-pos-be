package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedRoles(db *database.DB) {
	log.Println("seeding roles table...")

	var count int64
	db.Model(&entities.Role{}).Count(&count)
	if count > 0 {
		log.Println("roles already seeded")
		return
	}

	roles := []entities.Role{
		{
			Name: "super admin",
		},
		{
			Name: "admin",
		},
		{
			Name: "cashier",
		},
	}

	if err := db.Create(&roles).Error; err != nil {
		log.Fatalf("failed to seed roles: %v", err)
	}

	log.Println("seeded roles table")
}
