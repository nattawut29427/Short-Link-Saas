package database

import (
	"fmt"
	"go-links/configs"
	"go-links/internal/migrations"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(secret *configs.Secrets) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		secret.MySQL.User,
		secret.MySQL.Password,
		secret.MySQL.Host,
		secret.MySQL.Port,
		secret.MySQL.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	migrations.RunMigrations(db)

	return db
}
