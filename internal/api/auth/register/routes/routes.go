package routes

import (
	"go-links/internal/api/auth/register/handlers"
	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(e *echo.Echo, v1 *echo.Group, db *gorm.DB, rdb *redis.Client, jwtSecret string) {
	handler := handlers.NewAuthHandler(db, rdb, jwtSecret)
	
	// API Routes
	v1.POST("/auth/register", handler.CreateUser)
}
