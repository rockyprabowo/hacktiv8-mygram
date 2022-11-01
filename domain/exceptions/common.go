package exceptions

import (
	"errors"
)

var Unauthorized = errors.New("unauthorized")
var AuthTokenInvalid = errors.New("invalid authentication token")
var InvalidCredentials = errors.New("invalid credentials")
