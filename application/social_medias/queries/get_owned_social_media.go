package social_media_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	"rocky.my.id/git/mygram/application/social_medias/dto"
	"rocky.my.id/git/mygram/application/social_medias/payloads"
)

func (q SocialMediaQueries) GetOwnedSocialMedia(ctx context.Context, payload social_media_payloads.SocialMediaGetAllByOwnerPayload) ([]social_media_dto.SocialMediaWithUserDTO, error) {
	var socialMedias []social_media_dto.SocialMediaWithUserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetAll(ctx, payload)
	if err != nil {
		return nil, err
	}
	socialMedias = slices.Map(data, social_media_dto.MapFromEntityWithUser)

	return socialMedias, nil
}
