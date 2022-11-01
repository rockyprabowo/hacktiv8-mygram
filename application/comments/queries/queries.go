package comment_queries

import (
	"context"
	"github.com/rockyprabowo/h8-helpers/slices"
	contracts "rocky.my.id/git/mygram/application/comments/contracts"
	dto "rocky.my.id/git/mygram/application/comments/dto"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
)

type CommentQueries struct {
	Repository contracts.CommentRepositoryContract
}

func NewCommentQueries(repository contracts.CommentRepositoryContract) *CommentQueries {
	return &CommentQueries{Repository: repository}
}

func (q CommentQueries) GetAll(ctx context.Context, payload payloads.CommentGetAllPayload) (*dto.PaginatedCommentWithRelationsDTO, error) {
	var (
		paginatedCommentVM *dto.PaginatedCommentWithRelationsDTO
		comments           []dto.CommentWithRelationsDTO
	)

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, paginationState, err := q.Repository.GetAll(ctx, payload)
	if err != nil {
		return nil, err
	}
	comments = slices.Map(data, dto.MapFromEntityWithRelations)
	paginatedCommentVM = &dto.PaginatedCommentWithRelationsDTO{
		State: paginationState,
		Data:  comments,
	}

	return paginatedCommentVM, nil
}

func (q CommentQueries) GetOwnedPhotosComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) (*dto.PaginatedCommentWithRelationsDTO, error) {
	var (
		paginatedCommentVM *dto.PaginatedCommentWithRelationsDTO
		comments           []dto.CommentWithRelationsDTO
	)

	data, paginationState, err := q.Repository.GetOwnedPhotosComments(ctx, payload)
	if err != nil {
		return nil, err
	}
	comments = slices.Map(data, dto.MapFromEntityWithRelations)
	paginatedCommentVM = &dto.PaginatedCommentWithRelationsDTO{
		State: paginationState,
		Data:  comments,
	}

	return paginatedCommentVM, nil
}

func (q CommentQueries) GetOwnedComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) (*dto.PaginatedCommentWithRelationsDTO, error) {
	var (
		paginatedCommentVM *dto.PaginatedCommentWithRelationsDTO
		comments           []dto.CommentWithRelationsDTO
	)

	data, paginationState, err := q.Repository.GetOwnedComments(ctx, payload)
	if err != nil {
		return nil, err
	}
	comments = slices.Map(data, dto.MapFromEntityWithRelations)
	paginatedCommentVM = &dto.PaginatedCommentWithRelationsDTO{
		State: paginationState,
		Data:  comments,
	}

	return paginatedCommentVM, nil
}

func (q CommentQueries) GetByID(ctx context.Context, payload payloads.CommentGetByIDPayload) (*dto.CommentWithRelationsDTO, error) {
	var comment *dto.CommentWithRelationsDTO

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	data, err := q.Repository.GetByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	*comment = dto.MapFromEntityWithRelations(*data)

	return comment, nil
}
