package photo_contracts

import (
	"context"
	dto "rocky.my.id/git/mygram/application/photos/dto"
	payloads "rocky.my.id/git/mygram/application/photos/payloads"
)

type PhotoCommandsContract interface {
	Save(ctx context.Context, payload payloads.PhotoInsertPayload) (*dto.PhotoDTO, error)
	Update(ctx context.Context, payload payloads.PhotoUpdatePayload) (*dto.PhotoDTO, error)
	Delete(ctx context.Context, payload payloads.PhotoDeletePayload) (bool, error)
}
