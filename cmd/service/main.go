// cmd/service/main.go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"database/sql"

	"github.com/zaibon/go-template/internal/config"
	"github.com/zaibon/go-template/internal/handlers"
	"github.com/zaibon/go-template/internal/health"
	"github.com/zaibon/go-template/internal/log"
	"github.com/zaibon/go-template/internal/server"
	"github.com/zaibon/go-template/internal/service"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "my-go-service",
		Short: "My Go Microservice",
		Run:   runService,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runService(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	logger := log.NewLogger(cfg.Log.Level)

	svc := service.NewService(logger)

	//Example of database connection
	db, err := sql.Open("postgres", cfg.Database.ConnectionString)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	databaseChecker := health.NewDatabaseChecker(db)
	healthHandler := health.NewHealth(databaseChecker)

	httpHandlers := handlers.NewHandlers(healthHandler)
	httpServer := server.NewHTTPServer(cfg.Server.HTTP.Port, httpHandlers, logger)

	grpcHandlers := handlers.NewGRPCHandlers()
	grpcServer := server.NewGRPCServer(cfg.Server.GRPC.Port, grpcHandlers, logger)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Error("HTTP server failed", "error", err)
			cancel()
		}
	}()

	go func() {
		if err := grpcServer.Start(); err != nil {
			logger.Error("gRPC server failed", "error", err)
			cancel()
		}
	}()

	<-ctx.Done()

	logger.Info("Shutting down servers...")

	if err := httpServer.Stop(context.Background()); err != nil {
		logger.Error("HTTP server shutdown failed", "error", err)
	}

	if err := grpcServer.Stop(context.Background()); err != nil {
		logger.Error("gRPC server shutdown failed", "error", err)
	}

	if err := svc.Stop(context.Background()); err != nil {
		logger.Error("Service shutdown failed", "error", err)
	}
}
