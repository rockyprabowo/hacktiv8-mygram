package http_middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers/jwt"
)

func WithJWTValidation(parserFunc jwt_helpers.ParserFunc) []echo.MiddlewareFunc {
	jwtMiddlewareConfig := jwt_helpers.BuildEchoJWTMiddlewareConfig(parserFunc)

	return []echo.MiddlewareFunc{
		middleware.JWTWithConfig(jwtMiddlewareConfig),
		ValidateToken,
	}
}
