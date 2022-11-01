package user_contracts

import (
	"context"
	"rocky.my.id/git/mygram/application/common/authentication/dto"
	dto "rocky.my.id/git/mygram/application/users/dto"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
)

type UserQueriesContract interface {
	GetUserProfile(ctx context.Context, payload payloads.UserGetPayload) (*dto.UserDTO, error)
	AuthenticateUser(ctx context.Context, payload payloads.UserLoginPayload) (*dto.UserDTO, *authentication_dto.AuthTokenDTO, error)
}
