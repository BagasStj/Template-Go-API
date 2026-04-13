package routes

import (
	"golfscoreid-jng/config"
	"golfscoreid-jng/database"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/repositories"

	"github.com/gin-gonic/gin"
)

type route struct {
	cfg    config.Config
	logger logger.Logger
	repo   *repositories.Repository
	group  *gin.RouterGroup
}

func newV1Route(cfg config.Config, logger logger.Logger, r *gin.Engine) *route {
	return &route{
		cfg:    cfg,
		logger: logger,
		repo:   repositories.NewRepository(logger, database.GetReadDB(), database.GetWriteDB()),
		group:  r.Group("/v1"),
	}
}

func Init(r *gin.Engine, cfg config.Config, logger logger.Logger) {
	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Template Go API is running",
		})
	})

	// Initialize v1 routes
	v1 := newV1Route(cfg, logger, r)
	v1.initRoot()
	v1.initUser()
	v1.initPlayer()
	v1.initHole()
	v1.initScore()
}

func (rt *route) initRoot() {
	rt.group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Template Go API v1",
			"version": "1.0.0",
		})
	})
}
