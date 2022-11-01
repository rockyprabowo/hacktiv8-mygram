package comment_payloads

import (
	"context"
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/validator"
)

type CommentDeletePayload struct {
	UserID int `json:"user_id"`
	ID     int `json:"id" param:"id"`
}

func (p *CommentDeletePayload) ValidateAndAuthorizeWith(
	ctx context.Context,
	authorizerFunc authorization.ResourceOwnerAuthorizerFunc,
) error {
	authorizer := authorization.NewResourceOwnerAuthorizer(authorizerFunc, p.UserID, p.ID)
	return validator.NewWithResourceOwnerAuthorization(p, authorizer).Execute(ctx)
}

func (p *CommentDeletePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.UserID, validation.Required),
	)
}
