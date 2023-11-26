package httpserver

import (
	"net/http"
	"time"
	"test/internal/config"

)

type Server struct{
	httpServer *http.Server
}

func (s *Server) Run(config config.HTTPServer, handler http.Handler) error{
	s.httpServer = &http.Server{
		Addr:           ":" + config.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}