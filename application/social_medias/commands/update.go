package social_media_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/social_medias/dto"
	"rocky.my.id/git/mygram/application/social_medias/payloads"
)

func (c SocialMediaCommands) Update(ctx context.Context, payload social_media_payloads.SocialMediaUpdatePayload) (*social_media_dto.SocialMediaDTO, error) {
	var socialMedia social_media_dto.SocialMediaDTO

	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return nil, err
	}

	data, err := c.Repository.Update(ctx, payload)
	if err != nil {
		return nil, err
	}
	socialMedia = social_media_dto.MapFromEntity(*data)

	return &socialMedia, nil
}
