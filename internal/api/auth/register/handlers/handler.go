package handlers

import (
	"go-links/internal/api/auth/register/repositories"
	"go-links/internal/api/auth/register/services"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(db *gorm.DB, rdb *redis.Client, jwtSecret string) *AuthHandler {
	repo := repositories.NewAuthRepository(db, rdb)
	service := services.NewAuthService(repo, jwtSecret)
	return &AuthHandler{
		service: service,
	}
}
