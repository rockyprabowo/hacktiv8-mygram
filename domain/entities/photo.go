package entities

import (
	"rocky.my.id/git/mygram/domain/entities/embedded"
)

type Photo struct {
	ID       int    `gorm:"primaryKey"`
	Title    string `gorm:"type:TEXT NOT NULL"`
	Caption  string `gorm:"type:TEXT NOT NULL"`
	PhotoURL string `gorm:"type:TEXT NOT NULL"`
	UserID   int
	User     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	embedded.DateTime
}
