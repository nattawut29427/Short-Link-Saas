package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"go-links/internal/api/Url/models"
	"go-links/internal/api/Url/repositories"
)

type LinkService interface {
	CreateLink(ctx context.Context, req *models.UrlRequest, baseURL string) (*models.UrlResponse, error)
	GetOriginalURL(ctx context.Context, shortCode string) (string, error)
}

type service struct {
	repo repositories.LinkRepository
}

func NewLinkService(repo repositories.LinkRepository) LinkService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	link, err := s.repo.FindLinkByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}
	return link.OriginalURL, nil
}

func generateShortCode(n int) string {
	bytes := make([]byte, n/2)
	
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
