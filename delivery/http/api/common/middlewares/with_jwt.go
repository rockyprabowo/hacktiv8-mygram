package http_middlewares

import (
	"github.com/labstack/echo/v4"
)

func WithJWTValidation(jwtMiddleware echo.MiddlewareFunc, middlewares ...echo.MiddlewareFunc) []echo.MiddlewareFunc {

	jwtMiddlewares := []echo.MiddlewareFunc{
		jwtMiddleware,
		ValidateToken,
	}
	return append(jwtMiddlewares, middlewares...)
}
