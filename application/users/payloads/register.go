package user_payloads

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/domain/entities/rules"
	"rocky.my.id/git/mygram/domain/entities/value_objects"
)

type UserRegisterPayload struct {
	Username value_objects.Username `json:"username" fake:"{username}"`
	Password string                 `json:"password" fake:"skip"`
	Email    value_objects.Email    `json:"email" fake:"{email}"`
	Age      int                    `json:"age" fake:"{number:13,100}"`
}

func (p *UserRegisterPayload) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Username),
		validation.Field(&p.Password, rules.UserPasswordRules...),
		validation.Field(&p.Email),
		validation.Field(&p.Age, rules.UserAgeRules...),
	)
}
