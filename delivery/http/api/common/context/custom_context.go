package context

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/mygram/delivery/http/api/common/constants"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) HasUserToken() bool {
	return c.Get(constants.UserClaimsContextKey) != nil
}

func (c *CustomContext) HasPayload() bool {
	return c.Get(constants.Payload) != nil
}

func (c *CustomContext) GetPayload() any {
	return c.Get(constants.Payload)
}

func (c *CustomContext) SetPayload(payload any) {
	c.Set(constants.Payload, payload)
}
