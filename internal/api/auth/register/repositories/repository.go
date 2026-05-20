package repositories

import (
	"context"
	"go-links/internal/entities"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
}

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewAuthRepository(db *gorm.DB, rdb *redis.Client) AuthRepository {
	return &repository{
		db: db,
		rdb: rdb,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
