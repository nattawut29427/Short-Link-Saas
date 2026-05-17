package repositories

import (
	"context"
	"encoding/json"
	"time"

	"go-links/internal/entities"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link *entities.Link) error
	FindLinkByShortCode(ctx context.Context, shortCode string) (*entities.Link, error)
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
