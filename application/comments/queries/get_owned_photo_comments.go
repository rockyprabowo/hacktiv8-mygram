package comment_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	"rocky.my.id/git/mygram/application/comments/dto"
	"rocky.my.id/git/mygram/application/comments/payloads"
)

func (q CommentQueries) GetOwnedPhotoComments(ctx context.Context, payload comment_payloads.CommentGetByOwnerPayload) (*comment_dto.PaginatedCommentWithRelationsDTO, error) {
	var (
		paginatedCommentVM *comment_dto.PaginatedCommentWithRelationsDTO
		comments           []comment_dto.CommentWithRelationsDTO
	)

	data, paginationState, err := q.Repository.GetOwnedPhotosComments(ctx, payload)
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
