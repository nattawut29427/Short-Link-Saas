package models

import "github.com/google/uuid"

type UrlRequest struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	OriginalURL string    `json:"original_url" validate:"required"`
	Title       string    `json:"title"`
	IsActive    bool      `json:"is_active"`
}

type UrlResponse struct {
	Url         string    `json:"url"`
}