package jwt_helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rocky.my.id/git/mygram/domain/exceptions"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func BuildEchoJWTMiddlewareConfig(parserFunc func(string) (any, error)) middleware.JWTConfig {
	return middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (any, error) {
			return parserFunc(auth)
		},
	}
}

// noinspection GoUnusedExportedFunction
func MustGetUserClaims(ctx echo.Context) (*jwt_user.UserClaims, error) {
	tokenFromContext := ctx.Get("user")
	token, ok := tokenFromContext.(*jwt.Token)
	if ok {
		claims, claimsOK := token.Claims.(*jwt_user.UserClaims)
		if claimsOK {
			return claims, nil
		}
	}
	return nil, exceptions.AuthTokenInvalid
}

func ExtractUserClaims(ctx echo.Context) *jwt_user.UserClaims {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt_user.UserClaims)
	return claims
}
