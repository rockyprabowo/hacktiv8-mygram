package photo_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	contracts "rocky.my.id/git/mygram/application/photos/contracts"
	dto "rocky.my.id/git/mygram/application/photos/dto"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
)

type PhotoQueries struct {
	Repository contracts.PhotoRepositoryContract
}

func NewPhotoQueries(repository contracts.PhotoRepositoryContract) *PhotoQueries {
	return &PhotoQueries{Repository: repository}
}

func (q PhotoQueries) GetAll(ctx context.Context, payload payloads.PhotoGetAllPayload) (*dto.PaginatedPhotoWithUserDTO, error) {
	var (
		photos          []dto.PhotoWithUserDTO
		paginatedPhotos *dto.PaginatedPhotoWithUserDTO
	)

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, paginationState, err := q.Repository.GetAll(ctx, payload)
	if err != nil {
		return nil, err
	}
	photos = slices.Map(data, dto.MapFromEntityWithUser)
	paginatedPhotos = &dto.PaginatedPhotoWithUserDTO{
		State: paginationState,
		Data:  photos,
	}

	return paginatedPhotos, nil
}

func (q PhotoQueries) GetOwnedPhotos(ctx context.Context, payload payloads.PhotosGetByOwnerPayload) (*dto.PaginatedPhotoWithUserDTO, error) {
	var (
		photos          []dto.PhotoWithUserDTO
		paginatedPhotos *dto.PaginatedPhotoWithUserDTO
	)

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, paginationState, err := q.Repository.GetOwnedPhotos(ctx, payload)
	if err != nil {
		return nil, err
	}

	photos = slices.Map(data, dto.MapFromEntityWithUser)
	paginatedPhotos = &dto.PaginatedPhotoWithUserDTO{
		State: paginationState,
		Data:  photos,
	}

	return paginatedPhotos, nil
}

func (q PhotoQueries) GetByID(ctx context.Context, payload payloads.PhotoGetByIDPayload) (*dto.PhotoWithUserDTO, error) {
	var photo dto.PhotoWithUserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	photo = dto.MapFromEntityWithUser(*data)

	return &photo, nil
}
