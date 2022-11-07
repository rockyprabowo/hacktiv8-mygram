package photo_payloads

import (
	"github.com/jellydator/validation"
)

type PhotoGetByIDPayload struct {
	ID int `json:"id" param:"id"`
}

func (p *PhotoGetByIDPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.ID, validation.Required),
	)
}
