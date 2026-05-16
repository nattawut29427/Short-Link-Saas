package migrations

import (
	"go-links/internal/entities"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running migrations...")

	err := db.AutoMigrate(
		&entities.User{},
		&entities.Link{},
		&entities.Click{},
	)

	if err != nil {
		log.Fatalf("could not run migrations: %v", err)
	}

	log.Println("Migrations completed successfully.")
}
