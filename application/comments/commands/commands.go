package comment_commands

import (
	"context"
	contracts "rocky.my.id/git/mygram/application/comments/contracts"
	dto "rocky.my.id/git/mygram/application/comments/dto"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
)

type CommentCommands struct {
	Repository contracts.CommentRepositoryContract
}

func NewCommentCommands(repository contracts.CommentRepositoryContract) *CommentCommands {
	return &CommentCommands{Repository: repository}
}

func (c CommentCommands) Save(ctx context.Context, payload payloads.CommentInsertPayload) (*dto.CommentDTO, error) {
	var comment dto.CommentDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := c.Repository.Save(ctx, payload)
	if err != nil {
		return nil, err
	}
	comment = dto.MapFromEntity(*data)

	return &comment, nil
}

func (c CommentCommands) Update(ctx context.Context, payload payloads.CommentUpdatePayload) (*dto.CommentDTO, error) {
	var comment dto.CommentDTO

	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return nil, err
	}

	data, err := c.Repository.Update(ctx, payload)
	if err != nil {
		return nil, err
	}
	comment = dto.MapFromEntity(*data)

	return &comment, nil
}

func (c CommentCommands) Delete(ctx context.Context, payload payloads.CommentDeletePayload) (bool, error) {
	if err := payload.ValidateAndAuthorizeWith(ctx, c.Repository.Authorize); err != nil {
		return false, err
	}

	deleted, err := c.Repository.Delete(ctx, payload)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}
