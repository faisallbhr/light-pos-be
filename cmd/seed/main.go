package main

import (
	"flag"
	"log"

	"github.com/faisallbhr/light-pos-be/config"
	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/database/seeder"
)

func main() {
	only := flag.String("only", "", "run seeder only (ex: users)")
	flag.Parse()

	cfg := config.LoadConfig()
	database.InitDB(cfg)
	db := database.GetDB()

	if *only == "" {
		log.Println("run all seeders...")
		seeder.SeedAll(db)
	} else {
		log.Printf("run seeder: %s\n", *only)
		seeder.SeedByName(*only, db)
	}

	log.Println("seeding completed")
}
