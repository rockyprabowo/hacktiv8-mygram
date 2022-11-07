package user_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/users/dto"
	"rocky.my.id/git/mygram/application/users/payloads"
)

func (c UserCommands) UpdateUser(ctx context.Context, payload user_payloads.UserProfileUpdatePayload) (*user_dto.UserDTO, error) {
	var user user_dto.UserDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.UpdateUser(ctx, payload)
	if err != nil {
		return nil, err
	}
	user = user_dto.MapFromEntity(*data)

	return &user, nil
}
