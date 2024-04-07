package router

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(c *RouterContext) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := new(RouterContext)
	err := h(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(ctx.Response.Status)
	json.NewEncoder(w).Encode(ctx.Response.Data)
}
