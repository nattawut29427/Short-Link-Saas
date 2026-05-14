package services

import (
	"context"
	"go-api/internal/api/Address/repositories"
	"go-api/internal/entities"
)

type AddressService interface {
	GetAllAddresses(ctx context.Context) ([]entities.AddressTh, error)
	GetAddressByID(ctx context.Context, id uint) (*entities.AddressTh, error)
}

type service struct {
	repo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) AddressService {
	return &service{repo: repo}
}

func (s *service) GetAllAddresses(ctx context.Context) ([]entities.AddressTh, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) GetAddressByID(ctx context.Context, id uint) (*entities.AddressTh, error) {
	return s.repo.FindByID(ctx, id)
}
