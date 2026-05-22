package handlers

import (
	"go-links/internal/api/auth/login/repositories"
	"go-links/internal/api/auth/login/services"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service services.LoginService
}

func NewAuthHandler(db *gorm.DB, rdb *redis.Client, jwtSecret string) *AuthHandler {
	repo := repositories.NewLoginRepository(db, rdb)
	service := services.NewAuthService(repo, jwtSecret)
	return &AuthHandler{
		service: service,
	}
}
