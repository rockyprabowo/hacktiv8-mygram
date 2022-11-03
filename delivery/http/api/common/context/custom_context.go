package context

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) HasUserToken() bool {
	return c.Get("user") != nil
}

func (c *CustomContext) HasPayload() bool {
	return c.Get("payload") != nil
}

func (c *CustomContext) GetPayload() any {
	return c.Get("payload")
}

func (c *CustomContext) SetPayload(payload any) {
	c.Set("payload", payload)
}
