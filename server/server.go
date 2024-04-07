package server

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPServer interface {
	Start() error
}

type ServerContext struct {
	Context *context.Context
	// some other stuff...
}

type HandlerFunc func(ctx *ServerContext) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type Server struct {
	Host string
	Port string

	http.Server
}

func NewServer(host, port string, handler http.Handler) *Server {
	return &Server{
		Server: http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: handler,
		},
	}
}

// serve HTTP
func (s *Server) Start() error {
	return s.ListenAndServe()
}

// serve HTTPS
func (s *Server) StartTLS(certFile, keyFile string) error {
	return s.ListenAndServeTLS(certFile, keyFile)
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
