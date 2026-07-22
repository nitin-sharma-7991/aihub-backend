package router

import (
	"github.com/gin-gonic/gin"

	userHandler "github.com/nitin-sharma-7991/aihub-backend/internal/modules/user/handler"
)

// New creates and configures the application's router.
func New(
	userHandler *userHandler.UserHandler,
) *gin.Engine {

	router := gin.New()

	// Built-in Gin middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Public Routes
	router.GET("/health", Health)

	// User Routes
	router.POST("/users", userHandler.Create)
	router.GET("/users/:id", userHandler.GetByID)
	router.PUT("/users/:id", userHandler.Update)
	router.DELETE("/users/:id", userHandler.Delete)

	return router
}
