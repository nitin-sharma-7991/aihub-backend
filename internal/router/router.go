package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/handler"
)

// New creates and configures the application's router.
func New() *gin.Engine {

	router := gin.New()

	// Built-in Gin middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health endpoints
	router.GET("/health", handler.Health)

	return router
}
