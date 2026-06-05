package models

import (
	"time"

	"github.com/google/uuid"
)

type UrlRequest struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	OriginalURL string    `json:"original_url" validate:"required"`
	Title       string    `json:"title"`
	IsActive    bool      `json:"is_active"`
}

type UrlResponse struct {
	Url string `json:"url"`
}

type LinkItem struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	Title       string    `json:"title"`
	IsActive    bool      `json:"is_active"`
	Clicks      int64     `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LinkListResponse struct {
	Links []LinkItem `json:"links"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
}

type ClickItem struct {
	ID        uuid.UUID `json:"id"`
	Country   string    `json:"country"`
	Device    string    `json:"device"`
	Browser   string    `json:"browser"`
	Referrer  string    `json:"referrer"`
	ClickTime time.Time `json:"click_time"`
}

type LinkStatsResponse struct {
	Link        interface{}  `json:"link"`
	TotalClicks int64        `json:"total_clicks"`
	Clicks      []ClickItem  `json:"clicks"`
	Page        int          `json:"page"`
	Limit       int          `json:"limit"`
}
