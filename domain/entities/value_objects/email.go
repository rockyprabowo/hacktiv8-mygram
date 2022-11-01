package value_objects

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type Email string

func (e Email) Validate() error {
	return validation.Validate(string(e), validation.Required, is.Email)
}
