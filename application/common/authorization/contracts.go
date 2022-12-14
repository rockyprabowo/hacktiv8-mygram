package authorization

import "context"

type ResourceOwnerAuthorizerFunc = func(ctx context.Context, ownerID, resourceID any) error

type ResourceOwnerAuthorization interface {
	Authorize(ctx context.Context, ownerID, resourceID any) error
}

type NeedsResourceOwnerAuthorization interface {
	ValidateAndAuthorizeWith(ctx context.Context, authorizerFunc ResourceOwnerAuthorizerFunc) error
}
