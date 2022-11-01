package photo_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/photos/dto"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
)

type PhotoQueriesContract interface {
	GetAll(ctx context.Context, payload payloads.PhotoGetAllPayload) (*dto.PaginatedPhotoWithUserDTO, error)
	GetOwnedPhotos(ctx context.Context, payload payloads.PhotosGetByOwnerPayload) (*dto.PaginatedPhotoWithUserDTO, error)
	GetByID(ctx context.Context, payload payloads.PhotoGetByIDPayload) (*dto.PhotoWithUserDTO, error)
}
