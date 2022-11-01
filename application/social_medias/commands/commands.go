package social_media_commands

import (
	"context"
	contracts "rocky.my.id/git/mygram/application/social_medias/contracts"
	dto "rocky.my.id/git/mygram/application/social_medias/dto"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
)

type SocialMediaCommands struct {
	Repository contracts.SocialMediaRepositoryContract
}

func NewSocialMediaCommands(repository contracts.SocialMediaRepositoryContract) *SocialMediaCommands {
	return &SocialMediaCommands{Repository: repository}
}

func (c SocialMediaCommands) Save(ctx context.Context, payload payloads.SocialMediaInsertPayload) (*dto.SocialMediaDTO, error) {
	var socialMedia dto.SocialMediaDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.Save(ctx, payload)
	if err != nil {
		return nil, err
	}
	socialMedia = dto.MapFromEntity(*data)

	return &socialMedia, nil
}

func (c SocialMediaCommands) Update(ctx context.Context, payload payloads.SocialMediaUpdatePayload) (*dto.SocialMediaDTO, error) {
	var socialMedia dto.SocialMediaDTO

	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return nil, err
	}

	data, err := c.Repository.Update(ctx, payload)
	if err != nil {
		return nil, err
	}
	socialMedia = dto.MapFromEntity(*data)

	return &socialMedia, nil
}

func (c SocialMediaCommands) Delete(ctx context.Context, payload payloads.SocialMediaDeletePayload) (bool, error) {
	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return false, err
	}

	deleted, err := c.Repository.Delete(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
