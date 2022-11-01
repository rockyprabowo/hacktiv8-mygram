package user_contracts

import (
	"context"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
	"rocky.my.id/git/mygram/domain/entities"
)

type UserRepositoryContracts interface {
	GetUser(ctx context.Context, payload payloads.UserGetPayload) (*entities.User, error)
	CreateUser(ctx context.Context, payload payloads.UserRegisterPayload) (*entities.User, error)
	AuthenticateUser(ctx context.Context, payload payloads.UserLoginPayload) (*entities.User, error)
	CreateAuthToken(payload payloads.AuthTokenPayload) (string, error)
	UpdateUser(ctx context.Context, payload payloads.UserProfileUpdatePayload) (*entities.User, error)
	DeleteUser(ctx context.Context, payload payloads.UserDeletePayload) (bool, error)
}
