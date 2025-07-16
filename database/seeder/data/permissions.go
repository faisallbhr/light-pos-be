package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedPermissions(db *database.DB) {
	log.Println("seeding permissions table...")

	var count int64
	db.Model(&entities.Permission{}).Count(&count)
	if count > 0 {
		log.Println("permissions already seeded")
		return
	}

	permissions := []entities.Permission{
		{
			Name: "manage_users",
		},
		{
			Name: "manage_roles_and_permissions",
		},
		{
			Name: "manage_products",
		},
		{
			Name: "manage_inventory",
		},
		{
			Name: "manage_purchases",
		},
		{
			Name: "manage_sales",
		},
		{
			Name: "view_sales_reports",
		},
		{
			Name: "view_purchases_reports",
		},
		{
			Name: "view_financial_reports",
		},
		{
			Name: "view_inventory_reports",
		},
	}

	if err := db.Create(&permissions).Error; err != nil {
		log.Fatalf("failed to seed permissions: %v", err)
	}
}
