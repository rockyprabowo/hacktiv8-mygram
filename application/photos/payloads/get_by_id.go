package photo_payloads

import (
	"github.com/jellydator/validation"
)

type PhotoGetByIDPayload struct {
	UserID int `json:"user_id"`
	ID     int `json:"id" param:"id"`
}

func (p *PhotoGetByIDPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserID, validation.Required),
		validation.Field(&p.ID, validation.Required),
	)
}
