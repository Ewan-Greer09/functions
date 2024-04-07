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
	Path    string
	Method  string
	Handler Handler
}

type Response struct {
	Status int
	Data   interface{}
}

func (r *ServeMux) getHandler(method, path string) Handler {
	for _, route := range r.Routes {
		re := regexp.MustCompile(route.Path)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return NotFoundHandler()
}

func NotFoundHandler() Handler

func (r *ServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)

	// handler middlewares go here
	for _, m := range r.Middlewares {
		m(handler)
	}

	handler.ServeHTTP(w, req)
}

func (r *ServeMux) GET(path string, handler Handler) {
	r.AddRoute("GET", path, handler)
}

func (r *ServeMux) POST(path string, handler Handler) {
	r.AddRoute("POST", path, handler)
}

func (r *ServeMux) PUT(path string, handler Handler) {
	r.AddRoute("PUT", path, handler)
}

func (r *ServeMux) DELETE(path string, handler Handler) {
	r.AddRoute("DELETE", path, handler)
}

func (r *ServeMux) AddRoute(method, path string, handler Handler) {
	r.Routes = append(r.Routes, Route{Method: method, Path: path, Handler: handler})
}
