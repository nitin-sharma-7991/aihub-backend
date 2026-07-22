package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/config"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

// New creates a new HTTP server.
func New(cfg *config.Config, logger *zap.Logger, router *gin.Engine) *Server {

	httpServer := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.App.WriteTimeout) * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		logger:     logger,
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {

	s.logger.Info("Starting HTTP server",
		zap.String("address", s.httpServer.Addr),
	)

	err := s.httpServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {

	s.logger.Info("Shutting down server")

	return s.httpServer.Shutdown(ctx)
}
