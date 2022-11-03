package jwt_helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/delivery/http/api/common/consts"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type ParserFunc = func(string) (any, error)

func BuildEchoJWTMiddlewareConfig(parserFunc ParserFunc) middleware.JWTConfig {
	return middleware.JWTConfig{
		ContextKey: consts.UserClaimsContextKey,
		ParseTokenFunc: func(auth string, c echo.Context) (any, error) {
			return parserFunc(auth)
		},
	}
}

func ExtractUserClaims(ctx echo.Context) *jwt_user.UserClaims {
	token := ctx.Get(consts.UserClaimsContextKey).(*jwt.Token)
	claims := token.Claims.(*jwt_user.UserClaims)
	return claims
}
