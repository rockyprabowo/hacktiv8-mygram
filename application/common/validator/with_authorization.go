package validator

import (
	"context"
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/authorization"
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

	err := v.authorizer.AuthorizerFunction(ctx, v.authorizer.OwnerID, v.authorizer.ResourceID)
	if err != nil {
		return err
	}

	return nil
}
