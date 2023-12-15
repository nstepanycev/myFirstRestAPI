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

func (s *Server) Run(config config.HTTPServer, handler http.Handler) error{
	s.httpServer = &http.Server{
		Addr:           ":" + config.Host,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
