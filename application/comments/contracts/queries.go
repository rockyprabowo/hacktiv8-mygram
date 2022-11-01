package comment_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/comments/dto"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
)

type CommentQueriesContract interface {
	GetByID(ctx context.Context, payload payloads.CommentGetByIDPayload) (*dto.CommentWithRelationsDTO, error)
	GetAll(ctx context.Context, payload payloads.CommentGetAllPayload) (*dto.PaginatedCommentWithRelationsDTO, error)
	GetOwnedPhotosComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) (*dto.PaginatedCommentWithRelationsDTO, error)
	GetOwnedComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) (*dto.PaginatedCommentWithRelationsDTO, error)
}
