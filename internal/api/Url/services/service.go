package services

import (
	"context"
	"go-links/internal/api/url/models"
	"go-links/internal/api/url/repositories"
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