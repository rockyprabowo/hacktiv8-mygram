package photo_payloads

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/application/common/pagination"
)

type PhotoGetAllPayload struct {
	PaginationState pagination.State
}

func (p *PhotoGetAllPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PaginationState),
	)
}
