package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID               uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Username         string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password         string    `gorm:"type:varchar(255);not null" json:"-"`
	Title            string    `gorm:"type:text" json:"title"`
	Email            string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	TwoFactorEnabled bool      `gorm:"default:false" json:"two_factor_enabled"`
	EmailVerified    bool      `gorm:"default:false" json:"email_verified"`

	Links []Link `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"links,omitempty"`

	BaseModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
