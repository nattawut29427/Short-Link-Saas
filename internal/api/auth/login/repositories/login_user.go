package repositories

import (
	"context"
	"gorm.io/gorm"

	"go-links/internal/entities"

	"github.com/redis/go-redis/v9"
)

type LoginRepository interface {
    FindUserByEmail(ctx context.Context ,email string) (*entities.User, error)
}

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewLoginRepository(db *gorm.DB, rdb *redis.Client) LoginRepository {
	return &repository{
		db:  db,
		rdb: rdb,
	}
}


func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}