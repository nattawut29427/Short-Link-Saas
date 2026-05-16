package handlers

import (
	"go-links/internal/api/Url/repositories"
	"go-links/internal/api/Url/services"

	"gorm.io/gorm"
)

type LinkHandler struct {
	service services.LinkService
}

func NewLinkHandler(db *gorm.DB) *LinkHandler {
	repo := repositories.NewLinkRepository(db)
	service := services.NewLinkService(repo)
	return &LinkHandler{
		service: service,
	}
}
