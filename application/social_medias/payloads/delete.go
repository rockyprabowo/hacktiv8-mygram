package social_media_payloads

import (
	"context"
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/validator"
)

type SocialMediaDeletePayload struct {
	UserID int `json:"user_id"`
	ID     int `json:"id" param:"id"`
}

func (p *SocialMediaDeletePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.ID, validation.Required),
	)
}

func (p *SocialMediaDeletePayload) ValidateAndAuthorizeWith(
	ctx context.Context,
	authorizerFunc authorization.ResourceOwnerAuthorizerFunc,
) error {
	authorizer := authorization.NewResourceOwnerAuthorizer(authorizerFunc, p.UserID, p.ID)
	return validator.NewWithResourceOwnerAuthorization(p, authorizer).Execute(ctx)
}
