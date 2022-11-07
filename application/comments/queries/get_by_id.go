package comment_queries

import (
	"context"
	"rocky.my.id/git/mygram/application/comments/dto"
	"rocky.my.id/git/mygram/application/comments/payloads"
)

func (q CommentQueries) GetByID(ctx context.Context, payload comment_payloads.CommentGetByIDPayload) (*comment_dto.CommentWithRelationsDTO, error) {
	var comment *comment_dto.CommentWithRelationsDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	*comment = comment_dto.MapFromEntityWithRelations(*data)

	return comment, nil
}
