package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/zaibon/go-template/internal/handlers"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	srv    *grpc.Server
	logger *slog.Logger
	port   int
}

func NewGRPCServer(port int, handlers *handlers.GRPCHandlers, logger *slog.Logger) *GRPCServer {
	s := grpc.NewServer()
	handlers.Register(s) // Register gRPC services
	return &GRPCServer{
		srv:    s,
		logger: logger,
		port:   port,
	}
}

func (s *GRPCServer) Start() error {
	s.logger.Info("gRPC server starting", "port", s.port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	return s.srv.Serve(lis)
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("gRPC server stopping")
	s.srv.GracefulStop()
	return nil
}
