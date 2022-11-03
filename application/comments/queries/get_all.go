package comment_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	"rocky.my.id/git/mygram/application/comments/dto"
	"rocky.my.id/git/mygram/application/comments/payloads"
)

func (q CommentQueries) GetAll(ctx context.Context, payload comment_payloads.CommentGetAllPayload) (*comment_dto.PaginatedCommentWithRelationsDTO, error) {
	var (
		paginatedCommentVM *comment_dto.PaginatedCommentWithRelationsDTO
		comments           []comment_dto.CommentWithRelationsDTO
	)

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, paginationState, err := q.Repository.GetAll(ctx, payload)
	if err != nil {
		return nil, err
	}
	comments = slices.Map(data, comment_dto.MapFromEntityWithRelations)
	paginatedCommentVM = &comment_dto.PaginatedCommentWithRelationsDTO{
		State: paginationState,
		Data:  comments,
	}

	return paginatedCommentVM, nil
}
