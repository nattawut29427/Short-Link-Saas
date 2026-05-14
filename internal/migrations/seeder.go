package migrations

import (
	"encoding/json"
	"go-api/internal/entities"
	"log"
	"os"

	"gorm.io/gorm"
)

type GeoData struct {
	ID                uint   `json:"id"`
	ProvinceNameEn    string `json:"provinceNameEn"`
	ProvinceNameTh    string `json:"provinceNameTh"`
	DistrictNameEn    string `json:"districtNameEn"`
	DistrictNameTh    string `json:"districtNameTh"`
	SubdistrictNameEn string `json:"subdistrictNameEn"`
	SubdistrictNameTh string `json:"subdistrictNameTh"`
	PostalCode        int    `json:"postalCode"`
}

func SeedGeographyData(db *gorm.DB) {
	var count int64
	db.Model(&entities.AddressTh{}).Count(&count)

	if count > 0 {
		log.Println("Geography data already seeded, skipping...")
		return
	}

	log.Println("Seeding geography data from geography.json...")

	// อ่านไฟล์ JSON
	file, err := os.ReadFile("geography.json")
	if err != nil {
		log.Printf("could not read geography.json: %v", err)
		return
	}

	var rawData []GeoData
	if err := json.Unmarshal(file, &rawData); err != nil {
		log.Printf("could not unmarshal geography.json: %v", err)
		return
	}

	var addresses []entities.AddressTh
	for _, item := range rawData {
		addr := entities.AddressTh{
			ProvinceName: entities.MultiLang{
				Th: item.ProvinceNameTh,
				En: item.ProvinceNameEn,
			},
			DistrictName: entities.MultiLang{
				Th: item.DistrictNameTh,
				En: item.DistrictNameEn,
			},
			SubdistrictName: entities.MultiLang{
				Th: item.SubdistrictNameTh,
				En: item.SubdistrictNameEn,
			},
			PostalCode: item.PostalCode,
		}
		// ใช้ ID เดิมจาก JSON เพื่อให้ข้อมูลตรงกัน
		addr.ID = item.ID
		addresses = append(addresses, addr)
	}

	// ใช้ CreateInBatches เพื่อประสิทธิภาพ (ใส่ทีละ 1,000 แถว)
	if err := db.CreateInBatches(addresses, 1000).Error; err != nil {
		log.Printf("could not seed geography data: %v", err)
	} else {
		log.Println("Geography data seeded successfully!")
	}
}
