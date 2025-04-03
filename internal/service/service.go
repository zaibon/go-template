package service

import (
	"context"
	"log/slog"
)

type Service struct {
	logger *slog.Logger
	// ... other service fields
}

func NewService(logger *slog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) Start(ctx context.Context) error {
	s.logger.Info("Service starting", "version", "v0.0.1")
	// ... service logic
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	s.logger.Info("Service stopping")
	// ... cleanup logic
	return nil
}
