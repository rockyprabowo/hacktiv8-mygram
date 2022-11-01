package social_media_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	contracts "rocky.my.id/git/mygram/application/social_medias/contracts"
	dto "rocky.my.id/git/mygram/application/social_medias/dto"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
)

type SocialMediaQueries struct {
	Repository contracts.SocialMediaRepositoryContract
}

func NewSocialMediaQueries(repository contracts.SocialMediaRepositoryContract) *SocialMediaQueries {
	return &SocialMediaQueries{Repository: repository}
}

func (q SocialMediaQueries) GetAll(ctx context.Context, payload payloads.SocialMediaGetAllByOwnerPayload) ([]dto.SocialMediaWithUserDTO, error) {
	var socialMedias []dto.SocialMediaWithUserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetAll(ctx, payload)
	if err != nil {
		return nil, err
	}
	socialMedias = slices.Map(data, dto.MapFromEntityWithUser)

	return socialMedias, nil
}
