package social_media_contracts

import (
	"context"
	"rocky.my.id/git/mygram/application/common/authorization"
	payloads "rocky.my.id/git/mygram/application/social_medias/payloads"
	"rocky.my.id/git/mygram/domain/entities"
)

type SocialMediaRepositoryContract interface {
	authorization.ResourceOwnerAuthorization
	GetAll(ctx context.Context, owner payloads.SocialMediaGetAllByOwnerPayload) ([]entities.SocialMedia, error)
	Save(ctx context.Context, payload payloads.SocialMediaInsertPayload) (*entities.SocialMedia, error)
	Update(ctx context.Context, payload payloads.SocialMediaUpdatePayload) (*entities.SocialMedia, error)
	Delete(ctx context.Context, payload payloads.SocialMediaDeletePayload) (bool, error)
}
