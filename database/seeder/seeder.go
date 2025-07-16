package seeder

import (
	"log"
	"strings"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/database/seeder/data"
)

var registry = map[string]func(*database.DB){
	"users":            data.SeedUsers,
	"roles":            data.SeedRoles,
	"permissions":      data.SeedPermissions,
	"role_permissions": data.SeedRolePermissions,
	"user_roles":       data.SeedUserRoles,
}

func SeedAll(db *database.DB) {
	tx := db.DB.Begin()
	if tx.Error != nil {
		log.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	txDB := &database.DB{DB: tx}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatalf("seeding failed and rolled back due to panic: %v", r)
		}
	}()

	for name, fn := range registry {
		log.Printf("Seeding %s...", name)
		fn(txDB)
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("failed to commit transaction: %v", err)
	}

	log.Println("✅ all seeders committed successfully")
}

func SeedByName(name string, db *database.DB) {
	fn, ok := registry[strings.ToLower(name)]
	if !ok {
		log.Fatalf("unknown seeder: %s", name)
	}
	fn(db)
}
