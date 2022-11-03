package user_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/users/payloads"
)

func (c UserCommands) DeleteUser(ctx context.Context, payload user_payloads.UserDeletePayload) (bool, error) {
	if err := payload.Validate(); err != nil {
		return false, err
	}

	deleted, err := c.Repository.DeleteUser(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
