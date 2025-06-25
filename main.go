package main

import (
	"strings"
	"template-go-api/config"
	"template-go-api/database"
	"template-go-api/logger"
	"template-go-api/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	logger := logger.NewLogger()
	logger.Infof("Initiating Template Go API Service....")

	cfg, err := config.NewConfig(logger)
	if err != nil {
		logger.Errorf("error while init config: %+v", err)
		panic(err)
	}

	// CORS Configuration
	allowed := []string{"*"}
	if cfg.CorsAllowOrigins != "" {
		allowed = strings.Split(cfg.CorsAllowOrigins, ",")
	}

	corsConfig := cors.Config{
		AllowOrigins:     allowed,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "access_token", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	// Initialize database
	database.Init(cfg)

	// Initialize routes
	routes.Init(r, cfg, logger)

	// Start server
	logger.Infof("Server starting on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
