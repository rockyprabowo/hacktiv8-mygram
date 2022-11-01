package user_payloads

import (
	"errors"
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

const UserIDPayloadKey = "ID"
const UserEmailPayloadKey = "Email"
const UsernamePayloadKey = "Username"

type UserGetPayload struct {
	Key   string
	Value any
}

func (p *UserGetPayload) Validate() error {
	switch p.Key {
	case UserEmailPayloadKey:
		return validation.Validate(&p.Value, validation.Required, is.Email)
	case UsernamePayloadKey:
	case UserIDPayloadKey:
		return validation.Validate(&p.Value, validation.Required)
	}
	return errors.New("invalid key")
}
