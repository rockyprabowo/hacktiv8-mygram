package rules

import (
	"github.com/jellydator/validation"
)

var UserUsernameRules = []validation.Rule{
	validation.Required,
	validation.Length(1, 8),
}

var UserPasswordRules = []validation.Rule{
	validation.Required,
	validation.Length(8, 64),
}

var UserAgeRules = []validation.Rule{
	validation.Required,
	validation.Min(13),
	validation.Max(128),
}
