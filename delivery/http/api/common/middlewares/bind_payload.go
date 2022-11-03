package http_middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
	"rocky.my.id/git/mygram/delivery/http/api/common/responses"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func BindPayloadAndValidate[T any](_ *T) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctx.(*context.CustomContext)
			bindErr := BindPayloadFunc[T](c)
			if bindErr != nil {
				return bindErr
			}

			validateErr := ValidatePayloadFunc(c)
			if validateErr != nil {
				return validateErr
			}
			return next(c)
		}
	}
}

func BindPayloadWithUserClaimsAndValidate[T any](
	binderFunc func(claims *jwt_user.UserClaims, payload *T),
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctx.(*context.CustomContext)
			bindErr := BindPayloadFunc[T](c)
			if bindErr != nil {
				return bindErr
			}

			bindJWTErr := BindJWTUserClaimsFunc(c, binderFunc)
			if bindJWTErr != nil {
				return bindJWTErr
			}

			validationErr := ValidatePayloadFunc(c)
			if validationErr != nil {
				return validationErr
			}

			return next(c)
		}
	}
}

func BindPayloadFunc[T any](ctx echo.Context) error {
	c := ctx.(*context.CustomContext)
	payload := new(T)

	if bindErr := c.Bind(payload); bindErr != nil {
		return responses.EchoErrorResponse(http.StatusBadRequest, bindErr.Error())
	}
	c.SetPayload(payload)

	return nil
}
