package router

import (
	"github.com/gin-gonic/gin"

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
	userHandler *userHandler.UserHandler,
	authHandler *authHandler.AuthHandler,
	orgHandler *organizationHandler.OrganizationHandler,
	membershipHandler *membershipHandler.MembershipHandler,
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

		// Organization
		protected.POST("/organizations", orgHandler.Create)
		protected.GET("/organizations", orgHandler.GetAll)
		protected.GET("/organizations/:id", orgHandler.GetByID)
		protected.PUT("/organizations/:id", orgHandler.Update)
		protected.DELETE("/organizations/:id", orgHandler.Delete)

		// Membership
		protected.POST("/memberships", membershipHandler.Create)
		protected.GET("/memberships", membershipHandler.GetAll)
		protected.GET("/memberships/:id", membershipHandler.GetByID)
		protected.PUT("/memberships/:id", membershipHandler.Update)
		protected.DELETE("/memberships/:id", membershipHandler.Delete)
	}

	return router
}
