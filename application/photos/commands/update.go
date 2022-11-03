package photo_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/photos/dto"
	"rocky.my.id/git/mygram/application/photos/payloads"
)

func (c PhotoCommands) Update(ctx context.Context, payload photo_payloads.PhotoUpdatePayload) (*photo_dto.PhotoDTO, error) {
	var photo photo_dto.PhotoDTO

	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return nil, err
	}

	data, err := c.Repository.Update(ctx, payload)
	if err != nil {
		return nil, err
	}
	photo = photo_dto.MapFromEntity(*data)

	return &photo, nil
}
