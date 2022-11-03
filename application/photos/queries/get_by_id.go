package photo_queries

import (
	"context"
	"rocky.my.id/git/mygram/application/photos/dto"
	"rocky.my.id/git/mygram/application/photos/payloads"
)

func (q PhotoQueries) GetByID(ctx context.Context, payload photo_payloads.PhotoGetByIDPayload) (*photo_dto.PhotoWithUserDTO, error) {
	var photo photo_dto.PhotoWithUserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	photo = photo_dto.MapFromEntityWithUser(*data)

	return &photo, nil
}
