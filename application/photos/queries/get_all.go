package photo_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	"rocky.my.id/git/mygram/application/photos/dto"
	"rocky.my.id/git/mygram/application/photos/payloads"
)

func (q PhotoQueries) GetAll(ctx context.Context, payload photo_payloads.PhotoGetAllPayload) (*photo_dto.PaginatedPhotoWithUserDTO, error) {
	var (
		photos          []photo_dto.PhotoWithUserDTO
		paginatedPhotos *photo_dto.PaginatedPhotoWithUserDTO
	)

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, paginationState, err := q.Repository.GetAll(ctx, payload)
	if err != nil {
		return nil, err
	}
	photos = slices.Map(data, photo_dto.MapFromEntityWithUser)
	paginatedPhotos = &photo_dto.PaginatedPhotoWithUserDTO{
		State: paginationState,
		Data:  photos,
	}

	return paginatedPhotos, nil
}
