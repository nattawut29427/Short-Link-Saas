package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"go-links/configs"
	oauthHandlers "go-links/internal/api/auth/oauth/handlers"
	oauthRepositories "go-links/internal/api/auth/oauth/repositories"
	oauthServices "go-links/internal/api/auth/oauth/services"
)

func OAuthRoutes(e *echo.Echo, rateLimitGroup *echo.Group, db *gorm.DB, rdb *redis.Client, jwtSecret string) {
	secret := configs.GetSecret()
	repo := oauthRepositories.NewOAuthRepository(db)
	service := oauthServices.NewOAuthService(repo, jwtSecret, secret.GoogleClientID, secret.GoogleClientSecret, secret.GoogleRedirectURL)
	handler := oauthHandlers.NewOAuthHandler(service)

	group := rateLimitGroup.Group("/auth/oauth")

	group.GET("/google/url", handler.GoogleAuthURL)
	group.POST("/google/callback", handler.GoogleAuthCallback)
}