package photo_payloads

import (
	"context"
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/validator"
)

type PhotoUpdatePayload struct {
	UserID   int    `json:"user_id"`
	ID       int    `json:"id" param:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

func (p *PhotoUpdatePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Caption, validation.Required),
		validation.Field(&p.PhotoURL, validation.Required, is.URL),
	)
}

func (p *PhotoUpdatePayload) ValidateAndAuthorizeWith(
	ctx context.Context,
	authorizerFunc authorization.ResourceOwnerAuthorizerFunc,
) error {
	authorizer := authorization.NewResourceOwnerAuthorizer(authorizerFunc, p.UserID, p.ID)
	return validator.NewWithResourceOwnerAuthorization(p, authorizer).Execute(ctx)
}
