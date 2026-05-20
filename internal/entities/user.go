package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	Title    string    `gorm:"type:text"`

	Links []Link `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
