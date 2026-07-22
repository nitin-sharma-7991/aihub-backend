package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	user "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user"
	"github.com/nitin-sharma-7991/aihub-backend/internal/server"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/config"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/database"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/logger"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/router"

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

	// Set Gin mode
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize database
	db, err := database.New(cfg)
	if err != nil {
		logg.Fatal("Failed to connect database", zap.Error(err))
	}

	// Initialize User Module
	userModule := user.New(db)

	// Initialize router
	r := router.New(userModule.Handler)

	// Initialize server
	srv := server.New(cfg, logg, r)

	// Start server
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
