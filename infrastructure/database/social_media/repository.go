package social_media_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/exceptions"
	"rocky.my.id/git/mygram/infrastructure/database/common/authorizer"
	"rocky.my.id/git/mygram/infrastructure/database/common/scopes"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(DB *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{DB: DB}
}

func (r SocialMediaRepository) Authorize(ctx context.Context, ownerID, resourceID any) error {
	return db_authorize.NewGormResourceOwnerAuthorizer(r.DB, &entities.SocialMedia{}, ownerID, resourceID).Execute(ctx)
}

func (r SocialMediaRepository) GetAll(ctx context.Context, payload payloads.SocialMediaGetAllByOwnerPayload) ([]entities.SocialMedia, error) {
	var socialMedias []entities.SocialMedia
	err := r.DB.WithContext(ctx).
		Preload("User").
		Scopes(scopes.OwnedBy(payload.UserID), scopes.SortDescByID()).
		Find(&socialMedias).Error
	return socialMedias, err
}

func (r SocialMediaRepository) Save(ctx context.Context, payload payloads.SocialMediaInsertPayload) (*entities.SocialMedia, error) {
	socialMedia := &entities.SocialMedia{
		UserID:         payload.UserID,
		Name:           payload.Name,
		SocialMediaURL: payload.SocialMediaURL,
	}
	err := r.DB.WithContext(ctx).Save(socialMedia).Error
	return socialMedia, err
}

func (r SocialMediaRepository) BatchSave(ctx context.Context, payloads []entities.SocialMedia) (int64, error) {
	results := r.DB.WithContext(ctx).CreateInBatches(payloads, 100)
	return results.RowsAffected, results.Error
}

func (r SocialMediaRepository) Update(ctx context.Context, payload payloads.SocialMediaUpdatePayload) (*entities.SocialMedia, error) {
	var socialMedia entities.SocialMedia
	err := r.DB.WithContext(ctx).First(&socialMedia, payload.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.SocialMediaNotFoundError
		}
		return nil, err
	}

	socialMedia.Name = payload.Name
	socialMedia.SocialMediaURL = payload.SocialMediaURL

	err = r.DB.WithContext(ctx).Save(&socialMedia).Error
	return &socialMedia, err
}

func (r SocialMediaRepository) Delete(ctx context.Context, payload payloads.SocialMediaDeletePayload) (bool, error) {
	result := r.DB.WithContext(ctx).Delete(&entities.SocialMedia{}, payload.ID)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
