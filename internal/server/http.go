package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/zaibon/go-template/internal/handlers"
)

type HTTPServer struct {
	srv    *http.Server
	logger *slog.Logger
}

func NewHTTPServer(port int, handlers *handlers.Handlers, logger *slog.Logger) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.Healthz)
	mux.HandleFunc("/readyz", handlers.Readyz)
	mux.HandleFunc("/some-endpoint", handlers.SomeEndpoint)

	return &HTTPServer{
		srv:    &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: mux},
		logger: logger,
	}
}

func (s *HTTPServer) Start() error {
	s.logger.Info("HTTP server starting", "port", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	s.logger.Info("HTTP server stopping")
	return s.srv.Shutdown(ctx)
}
