package routes

import (
	"go-api/internal/api/Address/handlers"
	"go-api/internal/api/Address/repositories"
	"go-api/internal/api/Address/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAddressRoutes(e *echo.Group, db *gorm.DB) {
	// Initialize layers
	repo := repositories.NewAddressRepository(db)
	service := services.NewAddressService(repo)
	handler := handlers.NewAddressHandler(service)

	// Routes
	addressGroup := e.Group("/addresses")
	addressGroup.GET("", handler.GetAddresses)
	addressGroup.GET("/:id", handler.GetAddressByID)
}
