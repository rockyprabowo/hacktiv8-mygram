package responses

import (
	"github.com/labstack/echo/v4"
)

func EchoErrorResponse(code int, err any) error {
	return echo.NewHTTPError(
		code,
		ErrorResult{
			Error: err,
			Code:  code,
		},
	)
}
