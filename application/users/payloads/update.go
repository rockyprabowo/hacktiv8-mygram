package user_payloads

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type UserProfileUpdatePayload struct {
	ID       int
	Username value_objects.Username
	Email    value_objects.Email
}

func (p *UserProfileUpdatePayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.Username),
		validation.Field(&p.Email),
	)
}
