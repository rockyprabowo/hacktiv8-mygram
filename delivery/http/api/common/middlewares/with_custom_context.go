package http_middlewares

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/mygram/delivery/http/api/common/context"
)

func WithCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CustomContext{Context: c}
		return next(cc)
	}
}
