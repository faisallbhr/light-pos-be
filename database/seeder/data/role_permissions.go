package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedRolePermissions(db *database.DB) {
	log.Println("seeding role-permission associations...")

	var roles []entities.Role
	var permissions []entities.Permission
	db.Find(&roles)
	db.Find(&permissions)

	permMap := map[string]entities.Permission{}
	for _, p := range permissions {
		permMap[p.Name] = p
	}

	for _, role := range roles {
		var assignPerms []entities.Permission

		switch role.Name {
		case "super admin":
			assignPerms = permissions
		case "admin":
			assignPerms = []entities.Permission{
				permMap["manage_products"],
				permMap["manage_inventory"],
				permMap["manage_purchases"],
				permMap["manage_sales"],
				permMap["view_sales_reports"],
				permMap["view_inventory_reports"],
			}
		case "cashier":
			assignPerms = []entities.Permission{
				permMap["manage_sales"],
				permMap["view_sales_reports"],
			}
		}

		if err := db.Model(&role).Association("Permissions").Replace(assignPerms); err != nil {
			log.Fatalf("failed to assign permissions to %s: %v", role.Name, err)
		}
	}

	log.Println("role-permission associations seeded")
}
