package comment_payloads

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/pagination"
)

type CommentGetAllPayload struct {
	PaginationState pagination.State
}

func (p *CommentGetAllPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PaginationState),
	)
}
