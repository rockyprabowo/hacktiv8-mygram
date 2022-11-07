package photo_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/photos/dto"
	"rocky.my.id/git/mygram/application/photos/payloads"
)

func (c PhotoCommands) Save(ctx context.Context, payload photo_payloads.PhotoInsertPayload) (*photo_dto.PhotoDTO, error) {
	var photo photo_dto.PhotoDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.Save(ctx, payload)
	if err != nil {
		return nil, err
	}
	photo = photo_dto.MapFromEntity(*data)

	return &photo, nil
}
