package user_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/users/dto"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
)

type UserCommandsContract interface {
	RegisterUser(ctx context.Context, payload payloads.UserRegisterPayload) (*dto.UserDTO, error)
	UpdateUser(ctx context.Context, payload payloads.UserProfileUpdatePayload) (*dto.UserDTO, error)
	DeleteUser(ctx context.Context, payload payloads.UserDeletePayload) (bool, error)
}
