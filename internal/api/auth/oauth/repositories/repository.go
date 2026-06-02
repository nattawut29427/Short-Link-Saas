package repositories

import (
	"context"

	"go-links/internal/entities"

	"gorm.io/gorm"
)

type OAuthRepository interface {
	CreateOrUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetOrCreateUser(ctx context.Context, email, name string) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewOAuthRepository(db *gorm.DB) OAuthRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateOrUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	existingUser := &entities.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", user.Email).First(existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
				return nil, err
			}
			return user, nil
		} else {
			return nil, err
		}
	}

	// Update existing user with new info if needed
	existingUser.Title = user.Title
	if err := r.db.WithContext(ctx).Save(existingUser).Error; err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (r *userRepository) GetOrCreateUser(ctx context.Context, email, name string) (*entities.User, error) {
	user := &entities.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			newUser := &entities.User{
				Email: email,
				Title: name, // Using name as title
			}
			if err := r.db.WithContext(ctx).Create(newUser).Error; err != nil {
				return nil, err
			}
			return newUser, nil
		} else {
			return nil, err
		}
	}

	return user, nil
}