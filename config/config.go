package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4/middleware"
)

// Configuration :
type Configuration struct {
	Port       string
	CORSConfig middleware.CORSConfig
}

var config Configuration

func init() {
	configureEnvironment()
	configureCors()
}

// GetConfig :
func GetConfig() Configuration {
	return config
}

func configureEnvironment() {
	if os.Getenv("GO_ENV") == "" {
		godotenv.Load()
	}
	config.Port = os.Getenv("PORT")
}

func configureCors() {

	corsAllowedHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")
	corsAllowedMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
	corsAllowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	config.CORSConfig = middleware.CORSConfig{
		AllowHeaders: corsAllowedHeaders,
		AllowMethods: corsAllowedMethods,
		AllowOrigins: corsAllowedOrigins,
		MaxAge:       7200,
	}
}
