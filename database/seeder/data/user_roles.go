package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedUserRoles(db *database.DB) {
	log.Println("seeding user_roles table...")

	var count int64
	db.Model(&entities.UserRole{}).Count(&count)
	if count > 0 {
		log.Println("user_roles already seeded")
		return
	}

	var users []entities.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("failed to fetch users: %v", err)
	}

	var roles []entities.Role
	if err := db.Find(&roles).Error; err != nil {
		log.Fatalf("failed to fetch roles: %v", err)
	}

	userMap := map[string]uint{}
	for _, u := range users {
		userMap[u.Email] = u.ID
	}

	roleMap := map[string]uint{}
	for _, r := range roles {
		roleMap[r.Name] = r.ID
	}

	userRoles := []entities.UserRole{
		{
			UserID: userMap["superadmin@mail.com"],
			RoleID: roleMap["super admin"],
		},
		{
			UserID: userMap["admin@mail.com"],
			RoleID: roleMap["admin"],
		},
		{
			UserID: userMap["cashier@mail.com"],
			RoleID: roleMap["cashier"],
		},
	}

	if err := db.Create(&userRoles).Error; err != nil {
		log.Fatalf("failed to seed user_roles: %v", err)
	}

	log.Println("seeded user_roles table")
}
