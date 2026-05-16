package services

import (
	"context"
	"fmt"

	"go-links/internal/api/Url/models"
	"go-links/internal/entities"
)

func (s *service) CreateLink(ctx context.Context, req *models.UrlRequest, baseURL string) (*models.UrlResponse, error) {
	shortCode := ""
	shortCode = generateShortCode(6)

	link := &entities.Link{
		UserID:      req.UserID,
		OriginalURL: req.OriginalURL,
		ShortCode:   shortCode,
		Title:       req.Title,
		IsActive:    true,
	}

	if err := s.repo.CreateLink(ctx, link); err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("%s/%s", baseURL, link.ShortCode)

	return &models.UrlResponse{
		Url: fullURL,
	}, nil
}
