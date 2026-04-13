package config

import (
	"golfscoreid-jng/logger"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

type Config struct {
	// Server Configuration
	ServerPort string `env:"SERVER_PORT" default:"8080"`
	ENV        string `env:"ENV" default:"development"`

	// Database Configuration
	DBHostRead  string `env:"DB_HOST_RO" required:"false"`
	DBHostWrite string `env:"DB_HOST_RW" required:"false"`
	DBPort      int    `env:"DB_PORT" default:"5432"`
	DBUser      string `env:"DB_USER" required:"false"`
	DBPasswd    string `env:"DB_PASSWD" required:"false"`
	DBName      string `env:"DB_NAME" required:"false"`

	// CORS Configuration
	CorsAllowOrigins string `env:"CORS_ALLOW_ORIGINS" default:"*"`

	// JWT Secrets
	AccessTokenSecret  string `env:"ACCESS_TOKEN_SECRET" required:"false"`
	RefreshTokenSecret string `env:"REFRESH_TOKEN_SECRET" required:"false"`

	// API Configuration
	ApiSecret string `env:"API_SECRET" required:"false"`

	// Basic Auth
	BasicAuthUsername string `env:"BASIC_USERNAME" required:"false"`
	BasicAuthPassword string `env:"BASIC_PASSWORD" required:"false"`
}

func NewConfig(logger logger.Logger) (Config, error) {
	// Load .env file if exists
	err := godotenv.Load(".env")
	if err != nil {
		logger.Warnf("could not parse .env file, using environment variables: %+v", err)
	}

	config := Config{}
	if err := env.Set(&config); err != nil {
		logger.Errorf("could not parse environment variables, error: %+v", err)
		return Config{}, err
	}

	return config, nil
}
