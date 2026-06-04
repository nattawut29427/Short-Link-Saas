package services

import (
	"context"
	"crypto/rand"
	"fmt"

	"go-links/internal/api/url/models"
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

const base62Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	for code := 0; code < n; code++ {
		bytes[code] = base62Alphabet[bytes[code]%62]
	}
	return string(bytes)
}

