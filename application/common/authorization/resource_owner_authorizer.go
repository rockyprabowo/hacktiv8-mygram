package authorization

type ResourceOwnerAuthorizer struct {
	AuthorizerFunction ResourceOwnerAuthorizerFunc
	OwnerID            any
	ResourceID         any
}

func NewResourceOwnerAuthorizer(
	authorizerFunction ResourceOwnerAuthorizerFunc,
	ownerID any,
	resourceID any,
) *ResourceOwnerAuthorizer {
	return &ResourceOwnerAuthorizer{
		AuthorizerFunction: authorizerFunction,
		OwnerID:            ownerID,
		ResourceID:         resourceID,
	}
}
