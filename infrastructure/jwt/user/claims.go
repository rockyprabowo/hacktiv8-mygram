package jwt_user

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type UserClaims struct {
	UserID   int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func (u *UserClaims) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.UserID, validation.Required),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
