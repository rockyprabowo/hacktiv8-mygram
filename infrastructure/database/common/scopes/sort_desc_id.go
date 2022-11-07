package scopes

import (
	"gorm.io/gorm"
)

func SortDescByID() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}
}
