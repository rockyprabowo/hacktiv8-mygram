package jwt_helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type ParserFunc = func(string) (any, error)

func BuildEchoJWTMiddleware(parserFunc ParserFunc) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ContextKey: constants.UserClaimsContextKey,
		ParseTokenFunc: func(auth string, c echo.Context) (any, error) {
			return parserFunc(auth)
		},
	})
}

func ExtractUserClaims(ctx echo.Context) *jwt_user.UserClaims {
	token := ctx.Get(constants.UserClaimsContextKey).(*jwt.Token)
	claims := token.Claims.(*jwt_user.UserClaims)
	return claims
}
