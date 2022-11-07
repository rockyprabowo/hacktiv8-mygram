package comment_contracts

import (
	"context"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
	"rocky.my.id/git/mygram/application/common/authorization"
	"rocky.my.id/git/mygram/application/common/pagination"
	"rocky.my.id/git/mygram/domain/entities"
)

type CommentRepositoryContract interface {
	authorization.ResourceOwnerAuthorization
	GetByID(ctx context.Context, payload payloads.CommentGetByIDPayload) (*entities.Comment, error)
	GetAll(ctx context.Context, payload payloads.CommentGetAllPayload) ([]entities.Comment, pagination.State, error)
	GetOwnedPhotosComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) ([]entities.Comment, pagination.State, error)
	GetOwnedComments(ctx context.Context, payload payloads.CommentGetByOwnerPayload) ([]entities.Comment, pagination.State, error)
	Save(ctx context.Context, payload payloads.CommentInsertPayload) (*entities.Comment, error)
	Update(ctx context.Context, payload payloads.CommentUpdatePayload) (*entities.Comment, error)
	Delete(ctx context.Context, payload payloads.CommentDeletePayload) (bool, error)
}
