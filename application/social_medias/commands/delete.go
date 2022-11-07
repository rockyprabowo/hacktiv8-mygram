package social_media_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/social_medias/payloads"
)

func (c SocialMediaCommands) Delete(ctx context.Context, payload social_media_payloads.SocialMediaDeletePayload) (bool, error) {
	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return false, err
	}

	deleted, err := c.Repository.Delete(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
