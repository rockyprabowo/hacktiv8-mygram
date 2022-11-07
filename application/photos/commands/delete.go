package photo_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/photos/payloads"
)

func (c PhotoCommands) Delete(ctx context.Context, payload photo_payloads.PhotoDeletePayload) (bool, error) {
	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return false, err
	}

	deleted, err := c.Repository.Delete(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
