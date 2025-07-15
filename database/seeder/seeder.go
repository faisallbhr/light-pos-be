package seeder

import (
	"log"
	"strings"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/database/seeder/data"
)

var registry = map[string]func(*database.DB){
	"users": data.SeedUsers,
	// Tambah seed lain di sini
}

func SeedAll(db *database.DB) {
	for _, fn := range registry {
		fn(db)
	}
}

func SeedByName(name string, db *database.DB) {
	fn, ok := registry[strings.ToLower(name)]
	if !ok {
		log.Fatalf("unknown seeder: %s", name)
	}
	fn(db)
}
