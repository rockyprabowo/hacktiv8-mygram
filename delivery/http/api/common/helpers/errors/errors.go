package error_helpers

import (
	"errors"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/domain/exceptions"
)

func ExtractError(err error) error {
	if errors.Is(err, exceptions.EntityNotFound) {
		return responses.EchoErrorResponse(http.StatusNotFound, err.Error())
	}
	if errors.Is(err, exceptions.Unauthorized) {
		return responses.EchoErrorResponse(http.StatusUnauthorized, err.Error())
	}
	return responses.EchoErrorResponse(http.StatusUnprocessableEntity, err.Error())
}
