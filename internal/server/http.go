package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/zaibon/go-template/internal/handlers"
)

type HTTPServer struct {
	srv    *http.Server
	logger *slog.Logger
}

func NewHTTPServer(port int, handlers *handlers.Handlers, logger *slog.Logger) *HTTPServer {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handlers.Healthz)
	r.HandleFunc("/readyz", handlers.Readyz)
	r.HandleFunc("/some-endpoint", handlers.SomeEndpoint)

	// Middleware using mux.Use
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Request-ID"},
		ExposedHeaders:   []string{"X-Request-ID"},
		AllowCredentials: true,
	}).Handler)
	r.Use(requestIDMiddleware)
	r.Use(loggingMiddleware(logger))
	r.Use(recoverMiddleware(logger))

	return &HTTPServer{
		srv: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      r, // Use the mux router
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
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

// loggingMiddleware logs each HTTP request.
func loggingMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			// Wrap the writer to capture the status code.
			rw := &responseWriter{ResponseWriter: w}
			next.ServeHTTP(rw, r)
			duration := time.Since(start)

			// Use slog for structured logging.
			logger.Info("HTTP request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.statusCode,
				"duration", duration,
				"request_id", r.Header.Get("X-Request-ID"), //get request ID from header
			)
		})
	}
}

// recoverMiddleware recovers from panics and logs the error.
func recoverMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Log the panic with slog.
					logger.Error("Panic recovered", "error", err, "stack", fmt.Sprintf("%+v", err)) //include stack
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal server error")) //keep response generic
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// responseWriter is a wrapper for http.ResponseWriter that captures the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// requestIDKey is a custom type to avoid collisions with other context keys.
type requestIDKey struct{}

var RequestIDKey requestIDKey

// requestIDMiddleware generates and adds a request ID to the context.
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
