package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	s *http.Server
}

func (server *Server) Run(port string, timeout int) error {
	log.Print("Starting the server...")
	server.s = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 10,
		ReadTimeout:    time.Duration(timeout) * time.Second,
		WriteTimeout:   time.Duration(timeout) * time.Second,
	}
	return server.s.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.s.Shutdown(ctx)
}
