package router

import (
	"github.com/gin-gonic/gin"

	authHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/handler"
	userHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/handler"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/config"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/middleware"
)

// New creates and configures the application's router.
func New(
	cfg *config.Config,
	userHandler *userHandler.UserHandler,
	authHandler *authHandler.AuthHandler,
) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health Check
	router.GET("/health", Health)

	// API v1
	v1 := router.Group("/api/v1")

	// ---------------- Public Routes ----------------

	// Auth
	v1.POST("/auth/login", authHandler.Login)
	v1.POST("/auth/register", authHandler.Register)

	// ---------------- Protected Routes ----------------

	protected := v1.Group("")
	protected.Use(
		middleware.Auth(cfg.JWT.Secret),
	)

	{
		// Auth
		protected.GET("/auth/me", authHandler.Me)
		protected.POST("/auth/logout", authHandler.Logout)

		// Users
		protected.POST("/users", userHandler.Create)
		protected.GET("/users/:id", userHandler.GetByID)
		protected.PUT("/users/:id", userHandler.Update)
		protected.DELETE("/users/:id", userHandler.Delete)
	}

	return router
}
