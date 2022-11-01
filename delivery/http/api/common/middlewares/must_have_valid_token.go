package http_middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/domain/exceptions"
	jwt_user "rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func MustHaveValidToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenFromContext := c.Get("user")
		if token, tokenOK := tokenFromContext.(*jwt.Token); tokenOK {
			if claims, claimsOK := token.Claims.(*jwt_user.UserClaims); claimsOK {
				if err := claims.Validate(); err == nil {
					return next(c)
				}
			}
		}

		return responses.EchoErrorResponse(http.StatusUnauthorized, exceptions.AuthTokenInvalid)
	}
}
