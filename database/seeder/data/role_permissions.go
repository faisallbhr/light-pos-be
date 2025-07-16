package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedRolePermissions(db *database.DB) {
	log.Println("seeding users table...")

	var count int64
	db.Model(&entities.RolePermission{}).Count(&count)
	if count > 0 {
		log.Println("role permissions already seeded")
	}

	var roles []entities.Role
	if err := db.Find(&roles).Error; err != nil {
		log.Fatalf("failed to fetch roles: %v", err)
	}

	var permissions []entities.Permission
	if err := db.Find(&permissions).Error; err != nil {
		log.Fatalf("failed to fetch permissions: %v", err)
	}

	roleMap := map[string]uint{}
	for _, r := range roles {
		roleMap[r.Name] = r.ID
	}

	permMap := map[string]uint{}
	for _, p := range permissions {
		permMap[p.Name] = p.ID
	}

	var rolePerms []entities.RolePermission

	for _, perm := range permissions {
		rolePerms = append(rolePerms, entities.RolePermission{
			RoleID:       roleMap["super admin"],
			PermissionID: perm.ID,
		})
	}

	adminPerms := []string{
		"manage_products",
		"manage_inventory",
		"manage_purchases",
		"manage_sales",
		"view_sales_reports",
		"view_inventory_reports",
	}

	for _, name := range adminPerms {
		rolePerms = append(rolePerms, entities.RolePermission{
			RoleID:       roleMap["admin"],
			PermissionID: permMap[name],
		})
	}

	cashierPerms := []string{
		"manage_sales",
		"view_sales_reports",
	}
	for _, name := range cashierPerms {
		rolePerms = append(rolePerms, entities.RolePermission{
			RoleID:       roleMap["cashier"],
			PermissionID: permMap[name],
		})
	}

	if err := db.Create(&rolePerms).Error; err != nil {
		log.Fatalf("failed to seed role permissions: %v", err)
	}

	log.Println("seeded role_permissions table")
}
