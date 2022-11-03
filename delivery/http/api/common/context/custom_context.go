package context

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/mygram/delivery/http/api/common/consts"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) HasUserToken() bool {
	return c.Get(consts.UserClaimsContextKey) != nil
}

func (c *CustomContext) HasPayload() bool {
	return c.Get(consts.Payload) != nil
}

func (c *CustomContext) GetPayload() any {
	return c.Get(consts.Payload)
}

func (c *CustomContext) SetPayload(payload any) {
	c.Set(consts.Payload, payload)
}
