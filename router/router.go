package router

import (
	"net/http"
	"regexp"
)

type ServeMux struct {
	Routes      []Route
	Middlewares []MiddlewareFunc
}

func NewRouter() *ServeMux {
	return &ServeMux{}
}

type Route struct {
	Pattern string
	Method  string
	Handler HandlerFunc
}

type Response struct {
	Status int
	Data   any
}

func (r *ServeMux) getHandler(method, path string) HandlerFunc {
	for _, route := range r.Routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return NotFoundHandler()
}

func NotFoundHandler() HandlerFunc {
	return func(ctx *RouterContext) error {
		return ctx.SetResponse(404, "Not Found")
	}
}

func (r *ServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)

	// Wrap the handler with all middlewares
	for i := len(r.Middlewares) - 1; i >= 0; i-- {
		handler = r.Middlewares[i](handler)
	}

	handler.ServeHTTP(w, req)
}

func (r *ServeMux) GET(path string, handler HandlerFunc) {
	r.AddRoute("GET", path, handler)
}

func (r *ServeMux) POST(path string, handler HandlerFunc) {
	r.AddRoute("POST", path, handler)
}

func (r *ServeMux) PUT(path string, handler HandlerFunc) {
	r.AddRoute("PUT", path, handler)
}

func (r *ServeMux) DELETE(path string, handler HandlerFunc) {
	r.AddRoute("DELETE", path, handler)
}

func (r *ServeMux) AddRoute(method, path string, handler HandlerFunc) {
	r.Routes = append(r.Routes, Route{Method: method, Pattern: path, Handler: handler})
}

func (mux *ServeMux) Use(m ...MiddlewareFunc) {
	mux.Middlewares = append(mux.Middlewares, m...)
}
