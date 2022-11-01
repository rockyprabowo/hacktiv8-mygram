package entities

import (
	"rocky.my.id/git/mygram/domain/entities/embedded"
)

type Comment struct {
	ID      int `gorm:"primaryKey"`
	UserID  int
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhotoID int
	Photo   Photo  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message string `gorm:"type:TEXT NOT NULL"`
	embedded.DateTime
}
