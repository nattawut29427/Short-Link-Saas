package routes

import (
	"go-links/internal/api/Url/handlers"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterLinkRoutes(e *echo.Echo, v1 *echo.Group, db *gorm.DB) {
	handler := handlers.NewLinkHandler(db)
	
	// API Routes
	v1.POST("/links", handler.CreateUrl)

	// Redirect Route
	e.GET("/:shortCode", handler.Redirect)
}
