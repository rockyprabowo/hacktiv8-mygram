package middleware_funcs

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
	"rocky.my.id/git/mygram/delivery/http/api/common/exceptions"
	"rocky.my.id/git/mygram/delivery/http/api/common/helpers/jwt"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func BindJWTUserClaimsFunc[T any](ctx echo.Context, binderFunc func(claims *jwt_user.UserClaims, payload *T)) error {
	c := ctx.(*context.CustomContext)

	if !c.HasUserToken() {
		err := errors.New("this middleware (JWTUserClaimsBinder) was called before JWT middleware")
		c.Logger().Error(err)
		return err
	}
	if !c.HasPayload() {
		return responses.EchoErrorResponse(http.StatusBadRequest, http_exceptions.MissingPayload)
	}

	payload := c.GetPayload().(*T)
	claims := jwt_helpers.ExtractUserClaims(c)
	binderFunc(claims, payload)
	c.SetPayload(payload)

	return nil
}
