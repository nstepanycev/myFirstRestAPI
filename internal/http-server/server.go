package httpserver

import (
	"context"
	"net/http"
	"time"
	"test/internal/config"

)

type Server struct{
	httpServer *http.Server
}

func (s *Server) Run(config config.Config, handler http.Handler) error{
	s.httpServer = &http.Server{
		Addr:           ":" + config.HTTP_Server.Port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
