package repositories

import (
	"context"
	"go-api/internal/entities"
	"gorm.io/gorm"
)

type AddressRepository interface {
	FindAll(ctx context.Context) ([]entities.AddressTh, error)
	FindByID(ctx context.Context, id uint) (*entities.AddressTh, error)
}

type repository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &repository{db: db}
}

func (r *repository) FindAll(ctx context.Context) ([]entities.AddressTh, error) {
	var addresses []entities.AddressTh
	err := r.db.WithContext(ctx).Find(&addresses).Error
	return addresses, err
}

func (r *repository) FindByID(ctx context.Context, id uint) (*entities.AddressTh, error) {
	var address entities.AddressTh
	err := r.db.WithContext(ctx).First(&address, id).Error
	return &address, err
}
