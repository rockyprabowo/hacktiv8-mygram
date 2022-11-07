package comment_payloads

import (
	"context"
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/validator"
)

type CommentUpdatePayload struct {
	UserID  int    `json:"user_id"`
	ID      int    `json:"id" param:"id"`
	PhotoID int    `json:"photo_id"`
	Message string `json:"message"`
}

func (p *CommentUpdatePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Message, validation.Required),
	)
}

func (p *CommentUpdatePayload) ValidateAndAuthorizeWith(
	ctx context.Context,
	authorizerFunc authorization.ResourceOwnerAuthorizerFunc,
) error {
	authorizer := authorization.NewResourceOwnerAuthorizer(authorizerFunc, p.UserID, p.ID)
	return validator.NewWithResourceOwnerAuthorization(p, authorizer).Execute(ctx)
}
