package server

import (
	"net/http"
	"os"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Server struct {
	Logger Logger
	*http.Server
}

func NewServer(logger Logger, handler http.Handler) *Server {
	return &Server{
		Logger: logger,
		Server: &http.Server{
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	s.Logger.Printf("Server started")

	if os.Getenv("PORT") == "" {
		s.Addr = ":8080"
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
