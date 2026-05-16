package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {

	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID      uuid.UUID `gorm:"type:char(36);index"`
	ShortCode   string    `gorm:"type:varchar(100);unique;not null"`
	OriginalURL string    `gorm:"type:varchar(255);not null"`
	Title       string    `gorm:"type:varchar(255)"`
	IsActive    bool      `gorm:"default:true"`
	
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Clicks 		[]Click   `gorm:"foreignKey:LinkID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	BaseModel
}

func (l *Link) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}
