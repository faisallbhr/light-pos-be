package main

import (
	"log"

	"github.com/faisallbhr/light-pos-be/config"
	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/delivery/http/routes"
)

func main() {
	cfg := config.LoadConfig()

	database.InitDB(cfg)
	db := database.GetDB()

	router := routes.SetupRouter(db)
	if err := router.Run(":" + cfg.App.Port); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
