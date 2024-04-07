package router

import (
	"context"
	"net/http"
)

type RouterContext struct {
	Response Response
	R        *http.Request
	Ctx      context.Context
}

func (c *RouterContext) SetResponse(code int, data any) error {
	c.Response = Response{
		Status: code,
		Data:   data,
	}

	return nil
}
