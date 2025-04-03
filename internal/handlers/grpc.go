package handlers

import (
	"context"

	"github.com/zaibon/go-template/proto"
	"google.golang.org/grpc"
)

type GRPCHandlers struct {
	// ... dependencies
}

func NewGRPCHandlers() *GRPCHandlers {
	return &GRPCHandlers{}
}

func (h *GRPCHandlers) Register(s *grpc.Server) {
	// Register gRPC service here
}

func (h *GRPCHandlers) SomeGRPCMethod(ctx context.Context, req *proto.YourRequest) (*proto.YourResponse, error) {
	// Implement your gRPC method
	return &proto.YourResponse{}, nil
}
