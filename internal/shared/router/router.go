package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	authHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/handler"
	membershipHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/membership/handler"
	organizationHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/organization/handler"
	permissionHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/permission/handler"
	roleHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role/handler"
	rolePermissionHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/role_permission/handler"
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
	roleHandler *roleHandler.RoleHandler,
	permissionHandler *permissionHandler.PermissionHandler,
	rolePermissionHandler *rolePermissionHandler.RolePermissionHandler,
) *gin.Engine {

	router := gin.New()

	router.Use(
		middleware.RequestID(),
		middleware.Recovery(log),
		middleware.Logger(log),
		middleware.Timeout(10*time.Second),
	)
	// Health Check
	router.GET("/health", Health)

	router.GET("/test-timeout", func(c *gin.Context) {

		select {
		case <-time.After(15 * time.Second):
			c.JSON(200, gin.H{
				"message": "completed",
			})

		case <-c.Request.Context().Done():
			c.JSON(504, gin.H{
				"message": "request timeout",
			})
		}
	})

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

	// Roles
	roles := protected.Group("/roles")
	{
		roles.POST("", roleHandler.Create)
		roles.GET("", roleHandler.GetAll)
		roles.GET("/:id", roleHandler.GetByID)
		roles.PUT("/:id", roleHandler.Update)
		roles.DELETE("/:id", roleHandler.Delete)
	}

	// Permissions
	permissions := protected.Group("/permissions")
	{
		permissions.POST("", permissionHandler.Create)
		permissions.GET("", permissionHandler.GetAll)

		// Static route BEFORE /:id
		permissions.GET("/role/:role_id", permissionHandler.GetByRoleID)

		permissions.GET("/:id", permissionHandler.GetByID)
		permissions.PUT("/:id", permissionHandler.Update)
		permissions.DELETE("/:id", permissionHandler.Delete)
	}

	// Role Permissions
	rolePermissions := protected.Group("/roles/:role_id/permissions")
	{
		rolePermissions.POST("", rolePermissionHandler.AssignPermission)
		rolePermissions.GET("", rolePermissionHandler.GetPermissions)
	}

	protected.DELETE("/roles/:role_id/permissions/:permission_id", rolePermissionHandler.RevokePermission)

	return router
}
