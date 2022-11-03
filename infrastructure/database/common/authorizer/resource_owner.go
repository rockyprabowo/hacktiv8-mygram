package db_authorize

import (
	"context"
	"gorm.io/gorm"
	"rocky.my.id/git/mygram/domain/exceptions"
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

func (authorizer GormResourceOwnerAuthorizer) CheckResource(ctx context.Context) error {
	var rowCount int64
	authorizer.db.WithContext(ctx).
		Model(authorizer.model).
		Where("id = ?", authorizer.resourceID).
		Count(&rowCount)
	if rowCount == 0 {
		return exceptions.GenericNotFound(authorizer.model)
	}
	return nil
}

func (authorizer GormResourceOwnerAuthorizer) Execute(ctx context.Context) error {
	err := authorizer.CheckResource(ctx)
	if err != nil {
		return err
	}

	result := authorizer.db.WithContext(ctx).
		Select("id").
		Where("id = ? AND user_id = ?", authorizer.resourceID, authorizer.ownerID).
		First(authorizer.model)
	if result.RowsAffected != 1 {
		return exceptions.Unauthorized
	}

	return nil
}
