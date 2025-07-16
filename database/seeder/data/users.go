package data

import (
	"log"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(db *database.DB) {
	log.Println("seeding users table...")

	var count int64
	db.Model(&entities.User{}).Count(&count)
	if count > 0 {
		log.Println("users already seeded")
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	users := []entities.User{
		{
			Name:     "Super Admin",
			Email:    "superadmin@mail.com",
			Password: string(hashed),
		},
		{
			Name:     "Admin",
			Email:    "admin@mail.com",
			Password: string(hashed),
		},
		{
			Name:     "Cashier",
			Email:    "cashier@mail.com",
			Password: string(hashed),
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}

	log.Println("seeded users table")
}
