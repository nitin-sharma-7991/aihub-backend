package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	authHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/handler"
	membershipHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/handler"
	organizationHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/handler"
	userHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/handler"

	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/config"
	"github.com/nitin-sharma-7991/aihub-backend/internal/shared/middleware"
)

// New creates and configures the application's router.
func New(
	cfg *config.Config,
	log *zap.Logger,

	userHandler *userHandler.UserHandler,
	authHandler *authHandler.AuthHandler,
	orgHandler *organizationHandler.OrganizationHandler,
	membershipHandler *membershipHandler.MembershipHandler,
) *gin.Engine {

	router := gin.New()

	router.Use(
		middleware.RequestID(),
		middleware.Recovery(log),
		middleware.Logger(log),
	)
	// Health Check
	router.GET("/health", Health)

	// API v1
	v1 := router.Group("/api/v1")

	// ----------------------------------------------------------------
	// Public Routes
	// ----------------------------------------------------------------

	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	// ----------------------------------------------------------------
	// Protected Routes
	// ----------------------------------------------------------------

	protected := v1.Group("")
	protected.Use(middleware.Auth(cfg.JWT.Secret))

	// Auth
	protectedAuth := protected.Group("/auth")
	{
		protectedAuth.GET("/me", authHandler.Me)
		protectedAuth.POST("/logout", authHandler.Logout)
	}

	// Users
	users := protected.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.GetAll)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	// Organizations
	organizations := protected.Group("/organizations")
	{
		organizations.POST("", orgHandler.Create)
		organizations.GET("", orgHandler.GetAll)
		organizations.GET("/:id", orgHandler.GetByID)
		organizations.PUT("/:id", orgHandler.Update)
		organizations.DELETE("/:id", orgHandler.Delete)
	}

	// Memberships
	memberships := protected.Group("/memberships")
	{
		memberships.POST("", membershipHandler.Create)
		memberships.GET("", membershipHandler.GetAll)
		memberships.GET("/:id", membershipHandler.GetByID)
		memberships.PUT("/:id", membershipHandler.Update)
		memberships.DELETE("/:id", membershipHandler.Delete)
	}

	return router
}
