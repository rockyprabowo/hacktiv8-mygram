package photo_repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"rocky.my.id/git/mygram/application/common/pagination"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/domain/exceptions"
	"rocky.my.id/git/mygram/infrastructure/database/common/authorizer"
	"rocky.my.id/git/mygram/infrastructure/database/common/scopes"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(DB *gorm.DB) *PhotoRepository {
	return &PhotoRepository{DB: DB}
}

func (r PhotoRepository) Authorize(ctx context.Context, ownerID, resourceID any) error {
	return db_authorize.NewGormResourceOwnerAuthorizer(r.DB, &entities.Photo{}, ownerID, resourceID).Execute(ctx)

}

func (r PhotoRepository) GetByID(ctx context.Context, payload payloads.PhotoGetByIDPayload) (*entities.Photo, error) {
	photo := entities.Photo{}
	err := r.DB.WithContext(ctx).First(&photo, payload.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exceptions.PhotoNotFoundError
	}
	return &photo, err
}

func (r PhotoRepository) GetAll(ctx context.Context, payload payloads.PhotoGetAllPayload) ([]entities.Photo, pagination.State, error) {
	var photos []entities.Photo
	var totalPhotoCount int64
	paginationState := payload.PaginationState

	results := r.DB.WithContext(ctx).
		Preload("User").
		Scopes(scopes.Paginate(&paginationState), scopes.SortDescByID()).
		Find(&photos)
	if results.Error != nil {
		return nil, paginationState, results.Error
	}

	r.DB.WithContext(ctx).Model(&entities.Photo{}).Count(&totalPhotoCount)
	paginationState.SetPaginateTotalCount(totalPhotoCount)

	return photos, paginationState, nil
}

func (r PhotoRepository) GetOwnedPhotos(ctx context.Context, payload payloads.PhotosGetByOwnerPayload) ([]entities.Photo, pagination.State, error) {
	var photos []entities.Photo
	paginationState := payload.PaginationState
	results := r.DB.WithContext(ctx).
		Scopes(
			scopes.Paginate(&paginationState),
			scopes.OwnedBy(payload.UserID),
			scopes.SortDescByID(),
		).
		//Where("user_id = ?", payload.UserID).
		Find(&photos)
	if results.Error != nil {
		return nil, paginationState, results.Error
	}
	paginationState.SetPaginateTotalCount(results.RowsAffected)
	return photos, paginationState, nil
}

func (r PhotoRepository) Save(ctx context.Context, payload payloads.PhotoInsertPayload) (*entities.Photo, error) {
	photo := &entities.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoURL: payload.PhotoURL,
		UserID:   payload.UserID,
	}
	err := r.DB.WithContext(ctx).Create(&photo).Error
	return photo, err
}

func (r PhotoRepository) BatchSave(ctx context.Context, payloads []entities.Photo) (int64, error) {
	results := r.DB.WithContext(ctx).CreateInBatches(payloads, 100)
	return results.RowsAffected, results.Error
}

func (r PhotoRepository) Update(ctx context.Context, payload payloads.PhotoUpdatePayload) (*entities.Photo, error) {
	photo := entities.Photo{}
	err := r.DB.WithContext(ctx).First(&photo, payload.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = exceptions.PhotoNotFoundError
		return nil, err
	}

	photo.Title = payload.Title
	photo.Caption = payload.Caption
	photo.PhotoURL = payload.PhotoURL

	err = r.DB.WithContext(ctx).Save(photo).Error
	return &photo, err
}

func (r PhotoRepository) Delete(ctx context.Context, payload payloads.PhotoDeletePayload) (bool, error) {
	result := r.DB.WithContext(ctx).Delete(&entities.Photo{}, payload.ID)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
