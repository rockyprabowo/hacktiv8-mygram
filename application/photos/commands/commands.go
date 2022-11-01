package photo_commands

import (
	"context"
	contracts "rocky.my.id/git/mygram/application/photos/contracts"
	dto "rocky.my.id/git/mygram/application/photos/dto"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
)

type PhotoCommands struct {
	Repository contracts.PhotoRepositoryContract
}

func NewPhotoCommands(repository contracts.PhotoRepositoryContract) *PhotoCommands {
	return &PhotoCommands{Repository: repository}
}

func (c PhotoCommands) Save(ctx context.Context, payload payloads.PhotoInsertPayload) (*dto.PhotoDTO, error) {
	var photo dto.PhotoDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.Save(ctx, payload)
	if err != nil {
		return nil, err
	}
	photo = dto.MapFromEntity(*data)

	return &photo, nil
}

func (c PhotoCommands) Update(ctx context.Context, payload payloads.PhotoUpdatePayload) (*dto.PhotoDTO, error) {
	var photo dto.PhotoDTO

	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return nil, err
	}

	data, err := c.Repository.Update(ctx, payload)
	if err != nil {
		return nil, err
	}
	photo = dto.MapFromEntity(*data)

	return &photo, nil
}

func (c PhotoCommands) Delete(ctx context.Context, payload payloads.PhotoDeletePayload) (bool, error) {
	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return false, err
	}

	deleted, err := c.Repository.Delete(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
