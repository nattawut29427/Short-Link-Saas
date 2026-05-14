package handlers

import (
	"go-api/internal/api/Address/services"
)

type handler struct {
	service services.AddressService
}

func NewAddressHandler(service services.AddressService) *handler {
	return &handler{service: service}
}
