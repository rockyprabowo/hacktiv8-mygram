package entities

import (
	"rocky.my.id/git/mygram/domain/entities/embedded"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type User struct {
	ID           int                    `gorm:"primaryKey"`
	Username     value_objects.Username `gorm:"type:VARCHAR(32) NOT NULL UNIQUE"`
	Email        value_objects.Email    `gorm:"type:VARCHAR(192) NOT NULL UNIQUE"`
	Password     []byte                 `gorm:"type:VARCHAR(72) NOT NULL"`
	Age          int                    `gorm:"not null"`
	Photos       []Photo                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocialMedias []SocialMedia          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	embedded.DateTime
}
