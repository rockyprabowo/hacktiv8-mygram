package http_middlewares

import (
	"github.com/jellydator/validation"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

func ValidatePayloadFunc(c *context.CustomContext) error {
	if !c.HasPayload() {
		return responses.EchoErrorResponse(http.StatusBadRequest, "payload is empty")
	}

	payload := c.GetPayload().(validation.Validatable)
	if validationErr := payload.Validate(); validationErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, validationErr)
	}
	return nil
}
