package user_commands

import (
	"context"
	contracts "rocky.my.id/git/mygram/application/users/contracts"
	dto "rocky.my.id/git/mygram/application/users/dto"
	payloads "rocky.my.id/git/mygram/application/users/payloads"
)

type UserCommands struct {
	Repository contracts.UserRepositoryContracts
}

func NewUserCommands(repository contracts.UserRepositoryContracts) *UserCommands {
	return &UserCommands{Repository: repository}
}

func (c UserCommands) RegisterUser(ctx context.Context, payload payloads.UserRegisterPayload) (*dto.UserDTO, error) {
	var user dto.UserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.CreateUser(ctx, payload)
	if err != nil {
		return nil, err
	}
	user = dto.MapFromEntity(*data)

	return &user, nil
}

func (c UserCommands) UpdateUser(ctx context.Context, payload payloads.UserProfileUpdatePayload) (*dto.UserDTO, error) {
	var user dto.UserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.UpdateUser(ctx, payload)
	if err != nil {
		return nil, err
	}
	user = dto.MapFromEntity(*data)

	return &user, nil
}

func (c UserCommands) DeleteUser(ctx context.Context, payload payloads.UserDeletePayload) (bool, error) {
	if err := payload.Validate(); err != nil {
		return false, err
	}

	deleted, err := c.Repository.DeleteUser(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
