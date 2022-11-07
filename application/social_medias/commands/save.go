package social_media_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/social_medias/dto"
	"rocky.my.id/git/mygram/application/social_medias/payloads"
)

func (c SocialMediaCommands) Save(ctx context.Context, payload social_media_payloads.SocialMediaInsertPayload) (*social_media_dto.SocialMediaDTO, error) {
	var socialMedia social_media_dto.SocialMediaDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.Save(ctx, payload)
	if err != nil {
		return nil, err
	}
	socialMedia = social_media_dto.MapFromEntity(*data)

	return &socialMedia, nil
}
