package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Click struct {
	
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	LinkID    uuid.UUID `gorm:"type:char(36);index"`
	Country   string    `gorm:"type:varchar(255)"`
	Device    string    `gorm:"type:varchar(255)"`
	Browser   string    `gorm:"type:varchar(255)"`
	Referrer  string    `gorm:"type:varchar(255)"`
	ClickTime time.Time `gorm:"type:timestamp;autoCreateTime"`
	
	// Relationship
	Link Link `gorm:"foreignKey:LinkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	
	BaseModel
}

func (c *Click) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}


