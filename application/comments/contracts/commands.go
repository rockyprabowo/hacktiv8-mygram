package comment_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/comments/dto"
	payloads "rocky.my.id/git/mygram/application/comments/payloads"
)

type CommentCommandsContract interface {
	Save(ctx context.Context, payload payloads.CommentInsertPayload) (*dto.CommentDTO, error)
	Update(ctx context.Context, payload payloads.CommentUpdatePayload) (*dto.CommentDTO, error)
	Delete(ctx context.Context, payload payloads.CommentDeletePayload) (bool, error)
}
