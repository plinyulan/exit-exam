package main

import (
	"log"

	"github.com/plinyulan/exit-exam/internal/database"
)

func main() {
	db := database.New().GetClient()

	if err := database.SeedIfEmpty(db); err != nil {
		log.Fatal("❌ seed failed:", err)
	}

	log.Println("🌱 seed completed")
}
