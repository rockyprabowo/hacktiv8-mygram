package http_middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/domain/exceptions"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get(constants.UserClaimsContextKey)
		if claims, claimsOK := token.(*jwt.Token).Claims.(*jwt_user.UserClaims); claimsOK {
			if err := claims.Validate(); err != nil {
				return responses.EchoErrorResponse(http.StatusUnauthorized, exceptions.InvalidAuthToken)
			}
		}
		return next(c)
	}
}
