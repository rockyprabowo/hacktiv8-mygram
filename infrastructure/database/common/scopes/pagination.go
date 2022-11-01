package scopes

import (
	"gorm.io/gorm"
	"rocky.my.id/git/mygram/application/common/pagination"
)

func Paginate(state *pagination.State) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(state.GetOffset()).Limit(state.GetLimit())
	}
}
