package user_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/users/dto"
	"rocky.my.id/git/mygram/application/users/payloads"
)

func (c UserCommands) RegisterUser(ctx context.Context, payload user_payloads.UserRegisterPayload) (*user_dto.UserDTO, error) {
	var user user_dto.UserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.CreateUser(ctx, payload)
	if err != nil {
		return nil, err
	}
	user = user_dto.MapFromEntity(*data)

	return &user, nil
}
