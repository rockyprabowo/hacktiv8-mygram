package scopes

import (
	"gorm.io/gorm"
)

func OwnedBy(ownerID any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", ownerID)
	}
}
