package middleware_funcs

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
	"rocky.my.id/git/mygram/delivery/http/api/common/exceptions"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
)

func BindPayloadFunc[T any](ctx echo.Context) error {
	c := ctx.(*context.CustomContext)
	payload := new(T)

	if bindErr := c.Bind(payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, http_exceptions.MalformedPayload)
	}
	c.SetPayload(payload)

	return nil
}
