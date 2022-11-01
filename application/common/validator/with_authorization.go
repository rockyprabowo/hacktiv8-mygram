package validator

import (
	"context"
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/domain/exceptions"
)

type WithResourceOwnerAuthorization struct {
	payload    validation.Validatable
	authorizer *authorization.ResourceOwnerAuthorizer
}

func NewWithResourceOwnerAuthorization(
	payload validation.Validatable,
	authorizer *authorization.ResourceOwnerAuthorizer,
) *WithResourceOwnerAuthorization {
	return &WithResourceOwnerAuthorization{payload: payload, authorizer: authorizer}
}

func (v WithResourceOwnerAuthorization) Execute(ctx context.Context) error {
	if err := v.payload.Validate(); err != nil {
		return err
	}

	ok := v.authorizer.AuthorizerFunction(ctx, v.authorizer.OwnerID, v.authorizer.ResourceID)
	if !ok {
		return exceptions.Unauthorized
	}

	return nil
}
