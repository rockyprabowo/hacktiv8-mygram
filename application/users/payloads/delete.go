package user_payloads

import "github.com/jellydator/validation"

type UserDeletePayload struct {
	ID int `json:"id"`
}

func (p *UserDeletePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.ID, validation.Required),
	)
}
