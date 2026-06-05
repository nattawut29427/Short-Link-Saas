package handlers

import (
	"go-links/internal/api/url/repositories"
	"go-links/internal/api/url/services"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LinkHandler struct {
	service services.LinkService
}

func NewLinkHandler(db *gorm.DB, rdb *redis.Client) *LinkHandler {
	repo := repositories.NewLinkRepository(db, rdb)
	service := services.NewLinkService(repo)
	return &LinkHandler{
		service: service,
	}
}
