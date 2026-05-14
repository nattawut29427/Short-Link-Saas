package entities

import "gorm.io/gorm"

type AddressTh struct {
	gorm.Model
	ProvinceName    MultiLang `json:"provinceName" gorm:"type:json"`
	DistrictName    MultiLang `json:"districtName" gorm:"type:json"`
	SubdistrictName MultiLang `json:"subdistrictName" gorm:"type:json"`
	PostalCode      int       `json:"postalCode"`
}