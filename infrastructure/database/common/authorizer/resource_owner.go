package db_authorize

import (
	"context"
	"gorm.io/gorm"
)

type GormResourceOwnerAuthorizer struct {
	db         *gorm.DB
	model      any
	ownerID    any
	resourceID any
}

func NewGormResourceOwnerAuthorizer(db *gorm.DB, model, ownerID, resourceID any) *GormResourceOwnerAuthorizer {
	return &GormResourceOwnerAuthorizer{db: db, model: model, ownerID: ownerID, resourceID: resourceID}
}

func (authorizer GormResourceOwnerAuthorizer) Execute(ctx context.Context) bool {
	result := authorizer.db.WithContext(ctx).
		Select("id").
		Where("id = ? AND user_id = ?", authorizer.resourceID, authorizer.ownerID).
		First(authorizer.model)
	if result.Error != nil {
		return false
	}
	return result.RowsAffected == 1
}
