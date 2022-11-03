package exceptions

import (
	"errors"
	"fmt"
)

var UserNotFoundError = fmt.Errorf("user %w", EntityNotFound)
var EmailAlreadyRegistered = errors.New("e-mail already registered")
var UsernameAlreadyRegistered = errors.New("username already registered")
