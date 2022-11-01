package social_media_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/social_medias/dto"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
)

type SocialMediaCommandsContract interface {
	Save(ctx context.Context, payload payloads.SocialMediaInsertPayload) (*dto.SocialMediaDTO, error)
	Update(ctx context.Context, payload payloads.SocialMediaUpdatePayload) (*dto.SocialMediaDTO, error)
	Delete(ctx context.Context, payload payloads.SocialMediaDeletePayload) (bool, error)
}
