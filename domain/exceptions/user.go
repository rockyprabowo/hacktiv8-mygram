package exceptions

import "errors"

var UserNotFoundError = errors.New("user doesn't exist")
var EmailAlreadyRegistered = errors.New("e-mail already registered")
var UsernameAlreadyRegistered = errors.New("username already registered")
