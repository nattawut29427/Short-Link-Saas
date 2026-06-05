package services

import (
	"context"
	"net"
	"strings"
	"time"

	"go-links/internal/api/url/models"
	"go-links/internal/api/url/repositories"
	"go-links/internal/entities"

	"github.com/google/uuid"
)

type LinkService interface {
	CreateLink(ctx context.Context, req *models.UrlRequest, baseURL string) (*models.UrlResponse, error)
	GetOriginalURL(ctx context.Context, shortCode string) (string, error)
	GetOriginalURLAndRecordClick(ctx context.Context, shortCode, ip, userAgent, referrer string) (string, error)
	GetLinkByID(ctx context.Context, id uuid.UUID) (*entities.Link, error)
	GetUserLinks(ctx context.Context, userID uuid.UUID, page, limit int) (*models.LinkListResponse, error)
	GetLinkStats(ctx context.Context, linkID uuid.UUID, page, limit int) (*models.LinkStatsResponse, error)
}

type service struct {
	repo repositories.LinkRepository
}

func NewLinkService(repo repositories.LinkRepository) LinkService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOriginalURLAndRecordClick(ctx context.Context, shortCode, ip, userAgent, referrer string) (string, error) {
	link, err := s.repo.FindLinkByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		cleanIP := strings.Split(ip, ",")[0]
		host, _, err := net.SplitHostPort(cleanIP)
		if err == nil {
			cleanIP = host
		}

		device, browser := parseUserAgent(userAgent)

		click := &entities.Click{
			LinkID:   link.ID,
			Country:  "", // TODO: GeoIP lookup
			Device:   device,
			Browser:  browser,
			Referrer: referrer,
		}
		s.repo.RecordClick(ctx, click)
	}()

	return link.OriginalURL, nil
}

func (s *service) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	link, err := s.repo.FindLinkByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}
	return link.OriginalURL, nil
}

func parseUserAgent(ua string) (device, browser string) {
	ua = strings.ToLower(ua)
	
	switch {
	case strings.Contains(ua, "mobile"):
		device = "Mobile"
	case strings.Contains(ua, "tablet"):
		device = "Tablet"
	default:
		device = "Desktop"
	}

	switch {
	case strings.Contains(ua, "chrome"):
		browser = "Chrome"
	case strings.Contains(ua, "firefox"):
		browser = "Firefox"
	case strings.Contains(ua, "safari"):
		browser = "Safari"
	case strings.Contains(ua, "edge"):
		browser = "Edge"
	default:
		browser = "Other"
	}

	return
}

func (s *service) GetLinkByID(ctx context.Context, id uuid.UUID) (*entities.Link, error) {
	return s.repo.FindLinkByID(ctx, id)
}

func (s *service) GetUserLinks(ctx context.Context, userID uuid.UUID, page, limit int) (*models.LinkListResponse, error) {
	links, total, err := s.repo.FindLinksByUserID(ctx, userID, page, limit)
	if err != nil {
		return nil, err
	}

	items := make([]models.LinkItem, len(links))
	for i, link := range links {
		clicks, _ := s.repo.CountClicksByLinkID(ctx, link.ID)
		items[i] = models.LinkItem{
			ID:          link.ID,
			UserID:      link.UserID,
			ShortCode:   link.ShortCode,
			OriginalURL: link.OriginalURL,
			Title:       link.Title,
			IsActive:    link.IsActive,
			Clicks:      clicks,
			CreatedAt:   link.CreatedAt,
			UpdatedAt:   link.UpdatedAt,
		}
	}

	return &models.LinkListResponse{
		Links: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *service) GetLinkStats(ctx context.Context, linkID uuid.UUID, page, limit int) (*models.LinkStatsResponse, error) {
	link, err := s.repo.FindLinkByID(ctx, linkID)
	if err != nil {
		return nil, err
	}

	totalClicks, err := s.repo.CountClicksByLinkID(ctx, linkID)
	if err != nil {
		return nil, err
	}

	clicks, err := s.repo.FindClicksByLinkID(ctx, linkID, page, limit)
	if err != nil {
		return nil, err
	}

	clickItems := make([]models.ClickItem, len(clicks))
	for i, click := range clicks {
		clickItems[i] = models.ClickItem{
			ID:        click.ID,
			Country:   click.Country,
			Device:    click.Device,
			Browser:   click.Browser,
			Referrer:  click.Referrer,
			ClickTime: click.ClickTime,
		}
	}

	return &models.LinkStatsResponse{
		Link:       link,
		TotalClicks: totalClicks,
		Clicks:     clickItems,
		Page:       page,
		Limit:      limit,
	}, nil
}
