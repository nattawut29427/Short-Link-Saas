package database

import (
	"fmt"
	"go-links/configs"
	"go-links/internal/migrations"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(secret *configs.Secrets) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		secret.Postgres.Host,
		secret.Postgres.User,
		secret.Postgres.Password,
		secret.Postgres.Database,
		secret.Postgres.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	migrations.RunMigrations(db)

	return db
}
