package responses

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func WithData(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, data)
}
