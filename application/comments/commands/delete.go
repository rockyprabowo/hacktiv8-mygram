package comment_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/comments/payloads"
)

func (c CommentCommands) Delete(ctx context.Context, payload comment_payloads.CommentDeletePayload) (bool, error) {
	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return false, err
	}

	deleted, err := c.Repository.Delete(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
