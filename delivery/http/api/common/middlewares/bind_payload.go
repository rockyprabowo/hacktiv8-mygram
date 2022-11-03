package http_middlewares

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
	"rocky.my.id/git/mygram/delivery/http/api/common/middlewares/functions"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

func BindPayloadAndValidate[T any](payloadTypes *T) echo.MiddlewareFunc {
	_ = payloadTypes
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctx.(*context.CustomContext)
			bindErr := middleware_funcs.BindPayloadFunc[T](c)
			if bindErr != nil {
				return bindErr
			}

			validateErr := middleware_funcs.ValidatePayloadFunc(c)
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
			bindErr := middleware_funcs.BindPayloadFunc[T](c)
			if bindErr != nil {
				return bindErr
			}

			bindJWTErr := middleware_funcs.BindJWTUserClaimsFunc(c, binderFunc)
			if bindJWTErr != nil {
				return bindJWTErr
			}

			validationErr := middleware_funcs.ValidatePayloadFunc(c)
			if validationErr != nil {
				return validationErr
			}

			return next(c)
		}
	}
}
