package comment_commands

import (
	"context"
	"rocky.my.id/git/mygram/application/comments/dto"
	"rocky.my.id/git/mygram/application/comments/payloads"
)

func (c CommentCommands) Save(ctx context.Context, payload comment_payloads.CommentInsertPayload) (*comment_dto.CommentDTO, error) {
	var comment comment_dto.CommentDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.Save(ctx, payload)
	if err != nil {
		return nil, err
	}
	comment = comment_dto.MapFromEntity(*data)

	return &comment, nil
}
