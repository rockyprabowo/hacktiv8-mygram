package http_api_contracts

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/mygram/infrastructure/jwt/user"
)

type APIWithJWTRouterDeps struct {
	Engine     *echo.Echo
	JWTService *jwt_user.UserJWTService
}

type RouterContract interface {
	Setup()
}
