package routes

import (
	"go-links/internal/api/url/handlers"
	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterLinkRoutes(e *echo.Echo, v1 *echo.Group, db *gorm.DB, rdb *redis.Client) {
	handler := handlers.NewLinkHandler(db, rdb)

	// API Routes
	v1.POST("/links", handler.CreateUrl)
	v1.GET("/links", handler.GetMyLinks)
	v1.GET("/links/:id", handler.GetLinkByID)
	v1.GET("/links/:id/stats", handler.GetLinkStats)

	// Redirect Route
	e.GET("/:shortCode", handler.Redirect)
}
