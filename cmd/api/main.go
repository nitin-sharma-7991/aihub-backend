package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/config"
	"github.com/nitin-sharma-7991/aihub-backend/internal/logger"
	"github.com/nitin-sharma-7991/aihub-backend/internal/router"
	"github.com/nitin-sharma-7991/aihub-backend/internal/server"
	"go.uber.org/zap"
)

func main() {

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Initialize logger
	logg, err := logger.New(cfg.App.LogLevel)
	if err != nil {
		panic(err)
	}
	defer logg.Sync()

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize router
	r := router.New()

	// Initialize server
	srv := server.New(cfg, logg, r)

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, context.Canceled) {
			logg.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	logg.Info("Application Started")

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logg.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logg.Fatal("Server shutdown failed", zap.Error(err))
	}

	logg.Info("Server stopped gracefully")
}
