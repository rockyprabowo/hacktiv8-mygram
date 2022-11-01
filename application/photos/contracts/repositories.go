package photo_contracts

import (
	"context"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/pagination"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
	"rocky.my.id/git/mygram/domain/entities"
)

type PhotoRepositoryContract interface {
	authorization.ResourceOwnerAuthorization
	GetByID(ctx context.Context, payload payloads.PhotoGetByIDPayload) (*entities.Photo, error)
	GetAll(ctx context.Context, payload payloads.PhotoGetAllPayload) ([]entities.Photo, pagination.State, error)
	GetOwnedPhotos(ctx context.Context, payload payloads.PhotosGetByOwnerPayload) ([]entities.Photo, pagination.State, error)
	Save(ctx context.Context, payload payloads.PhotoInsertPayload) (*entities.Photo, error)
	Update(ctx context.Context, payload payloads.PhotoUpdatePayload) (*entities.Photo, error)
	Delete(ctx context.Context, payload payloads.PhotoDeletePayload) (bool, error)
}
