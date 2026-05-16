package repositories

import (
	"context"
	"go-links/internal/entities"
	"gorm.io/gorm"
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link *entities.Link) error
	FindLinkByShortCode(ctx context.Context, shortCode string) (*entities.Link, error)
}

type repository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateLink(ctx context.Context, link *entities.Link) error {
	return r.db.WithContext(ctx).Create(link).Error
}

func (r *repository) FindLinkByShortCode(ctx context.Context, shortCode string) (*entities.Link, error) {
	var link entities.Link
	err := r.db.WithContext(ctx).Where("short_code = ? AND is_active = ?", shortCode, true).First(&link).Error
	return &link, err
}
