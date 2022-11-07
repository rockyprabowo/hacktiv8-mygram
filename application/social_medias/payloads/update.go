package social_media_payloads

import (
	"context"

	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/validator"
)

type SocialMediaUpdatePayload struct {
	UserID         int    `json:"user_id"`
	ID             int    `json:"id" param:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (p *SocialMediaUpdatePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.SocialMediaURL, validation.Required, is.URL),
	)
}

func (p *SocialMediaUpdatePayload) ValidateAndAuthorizeWith(
	ctx context.Context,
	authorizerFunc authorization.ResourceOwnerAuthorizerFunc,
) error {
	authorizer := authorization.NewResourceOwnerAuthorizer(authorizerFunc, p.UserID, p.ID)
	return validator.NewWithResourceOwnerAuthorization(p, authorizer).Execute(ctx)
}
