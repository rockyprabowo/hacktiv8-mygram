package entities

import (
	"rocky.my.id/git/mygram/domain/entities/embedded"
)

type SocialMedia struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"type:VARCHAR(255) NOT NULL"`
	SocialMediaURL string `gorm:"type:TEXT NOT NULL"`
	UserID         int
	User           User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	embedded.DateTime
}
