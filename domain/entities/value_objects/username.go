package value_objects

import (
	"github.com/jellydator/validation"
	"rocky.my.id/git/mygram/domain/entities/rules"
)

type Username string

func (u Username) Validate() error {
	return validation.Validate(string(u), rules.UserUsernameRules...)
}
