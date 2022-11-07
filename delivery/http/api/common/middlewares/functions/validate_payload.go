package middleware_funcs

import (
	"errors"
	"github.com/jellydator/validation"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

func ValidatePayloadFunc(c *context.CustomContext) error {
	if !c.HasPayload() {
		return responses.EchoErrorResponse(http.StatusBadRequest, "payload is empty")
	}

	payload, isValidatable := c.GetPayload().(validation.Validatable)
	if !isValidatable {
		err := errors.New("this middleware (ValidatePayload) was called on a non-validatable payload")
		c.Logger().Error(err)
		return err
	}
	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}
	return nil
}
