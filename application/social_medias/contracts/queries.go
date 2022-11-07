package social_media_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/social_medias/dto"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
)

type SocialMediaQueriesContract interface {
	GetOwnedSocialMedia(ctx context.Context, payload payloads.SocialMediaGetAllByOwnerPayload) ([]dto.SocialMediaWithUserDTO, error)
}
