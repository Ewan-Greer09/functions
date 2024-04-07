package router

import (
	"encoding/json"
	"net/http"
)

type Handler func(c *CustomContext) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := new(CustomContext)
	err := h(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}

	w.WriteHeader(ctx.Response.Status)
	json.NewEncoder(w).Encode(ctx.Response.Data)
}
