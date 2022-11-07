package user_handlers

import "github.com/labstack/echo/v4"

type UserHTTPHandlerContract interface {
	GetUser(ctx echo.Context) error
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}
