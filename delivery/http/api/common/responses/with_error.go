package responses

import (
	"errors"
	"net/http"
	"rocky.my.id/git/mygram/domain/exceptions"
)

func WithError[T error](err T) error {
	if errors.Is(err, exceptions.EntityNotFound) {
		return EchoErrorResponse(http.StatusNotFound, err.Error())
	}
	if errors.Is(err, exceptions.Unauthorized) {
		return EchoErrorResponse(http.StatusUnauthorized, err.Error())
	}
	return EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
}
