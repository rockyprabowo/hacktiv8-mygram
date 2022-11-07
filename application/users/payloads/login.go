package user_payloads

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/domain/entities/rules"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type UserLoginPayload struct {
	Email    value_objects.Email `json:"email"`
	Password string              `json:"password"`
}

func (p *UserLoginPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Email),
		validation.Field(&p.Password, rules.UserPasswordRules...),
	)
}
