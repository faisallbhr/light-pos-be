package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func SeedUserRoles(db *database.DB) {
	log.Println("seeding user-role associations...")

	var users []entities.User
	var roles []entities.Role
	db.Find(&users)
	db.Find(&roles)

	roleMap := map[string]entities.Role{}
	for _, r := range roles {
		roleMap[r.Name] = r
	}

	for _, u := range users {
		var role entities.Role
		switch u.Email {
		case "superadmin@mail.com":
			role = roleMap["super admin"]
		case "admin@mail.com":
			role = roleMap["admin"]
		case "cashier@mail.com":
			role = roleMap["cashier"]
		default:
			continue
		}

		if err := db.Model(&u).Association("Roles").Replace([]entities.Role{role}); err != nil {
			log.Fatalf("failed to assign role to %s: %v", u.Email, err)
		}
	}

	log.Println("user-role associations seeded")
}
