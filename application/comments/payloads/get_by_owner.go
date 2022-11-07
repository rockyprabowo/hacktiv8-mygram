package comment_payloads

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/pagination"
)

type CommentGetByOwnerPayload struct {
	UserID          int `json:"user_id"`
	PaginationState pagination.State
}

func (p *CommentGetByOwnerPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.PaginationState),
	)
}
