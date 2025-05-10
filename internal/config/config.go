package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	ServerPort  string
	DatabaseURL string
	JWTSecret   string
}

// LoadConfig loads environment variables and returns a Config struct
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// Ignore error if .env file is not found
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/goaddd?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
	}, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
