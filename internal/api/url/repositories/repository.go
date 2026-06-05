package repositories

import (
	"context"
	"encoding/json"
	"time"

	"go-links/internal/entities"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link *entities.Link) error
	FindLinkByShortCode(ctx context.Context, shortCode string) (*entities.Link, error)
	FindLinkByID(ctx context.Context, id uuid.UUID) (*entities.Link, error)
	FindLinksByUserID(ctx context.Context, userID uuid.UUID, page, limit int) ([]entities.Link, int64, error)
	CountClicksByLinkID(ctx context.Context, linkID uuid.UUID) (int64, error)
	FindClicksByLinkID(ctx context.Context, linkID uuid.UUID, page, limit int) ([]entities.Click, error)
	RecordClick(ctx context.Context, click *entities.Click) error
}

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewLinkRepository(db *gorm.DB, rdb *redis.Client) LinkRepository {
	return &repository{
		db:  db,
		rdb: rdb,
	}
}

func (r *repository) CreateLink(ctx context.Context, link *entities.Link) error {
	return r.db.WithContext(ctx).Create(link).Error
}

func (r *repository) FindLinkByShortCode(ctx context.Context, shortCode string) (*entities.Link, error) {
	cacheKey := "link:" + shortCode

	val, err := r.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var link entities.Link
		if err := json.Unmarshal([]byte(val), &link); err == nil {
			return &link, nil
		}
	}

	var link entities.Link
	err = r.db.WithContext(ctx).Where("short_code = ? AND is_active = ?", shortCode, true).First(&link).Error
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(link); err == nil {
		r.rdb.Set(ctx, cacheKey, data, 24*time.Hour)
	}

	return &link, nil
}

func (r *repository) FindLinkByID(ctx context.Context, id uuid.UUID) (*entities.Link, error) {
	var link entities.Link
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&link).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *repository) FindLinksByUserID(ctx context.Context, userID uuid.UUID, page, limit int) ([]entities.Link, int64, error) {
	var links []entities.Link
	var total int64

	offset := (page - 1) * limit

	err := r.db.WithContext(ctx).Model(&entities.Link{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(limit).Find(&links).Error
	if err != nil {
		return nil, 0, err
	}

	return links, total, nil
}

func (r *repository) CountClicksByLinkID(ctx context.Context, linkID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entities.Click{}).Where("link_id = ?", linkID).Count(&count).Error
	return count, err
}

func (r *repository) FindClicksByLinkID(ctx context.Context, linkID uuid.UUID, page, limit int) ([]entities.Click, error) {
	var clicks []entities.Click
	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Where("link_id = ?", linkID).Order("click_time DESC").Offset(offset).Limit(limit).Find(&clicks).Error
	return clicks, err
}

func (r *repository) RecordClick(ctx context.Context, click *entities.Click) error {
	return r.db.WithContext(ctx).Create(click).Error
}
