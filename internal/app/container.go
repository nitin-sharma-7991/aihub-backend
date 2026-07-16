package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nitin-sharma-7991/aihub-backend/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Router *gin.Engine
}

func New(
	cfg *config.Config,
	log *zap.Logger,
	db *gorm.DB,
	router *gin.Engine,
) *Container {

	return &Container{
		Config: cfg,
		Logger: log,
		DB:     db,
		Router: router,
	}
}
