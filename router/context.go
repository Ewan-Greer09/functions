package router

import (
	"context"
	"net/http"
)

type CustomContext struct {
	Response Response
	R        *http.Request
	Ctx      context.Context
}

func (c *CustomContext) SetResponse(code int, data any) {
	c.Response = Response{
		Status: code,
		Data:   data,
	}
}
