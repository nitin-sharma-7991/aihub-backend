package router

import (
	"github.com/gin-gonic/gin"

	authHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/auth/handler"
	userHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/handler"
)

// New creates and configures the application's router.
func New(
	userHandler *userHandler.UserHandler,
	authHandler *authHandler.AuthHandler,
) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health
	router.GET("/health", Health)

	// User Routes
	router.POST("/users", userHandler.Create)
	router.GET("/users/:id", userHandler.GetByID)
	router.PUT("/users/:id", userHandler.Update)
	router.DELETE("/users/:id", userHandler.Delete)

	// Auth Routes
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/register", authHandler.Register)

	// Protected (after middleware)
	// auth := router.Group("/auth")
	// auth.Use(middleware.JWT(...))
	// {
	//     auth.GET("/me", authHandler.Me)
	//     auth.POST("/logout", authHandler.Logout)
	// }

	return router
}
