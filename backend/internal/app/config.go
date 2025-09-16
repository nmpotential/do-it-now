package app

import (
	"fmt"
	"os"
)

// Config struct holds all your app's configuration values.
// Add new fields here as your app grows.
type Config struct {
	DBURL      string // Database connection string (from env: DB_URL)
	ServerPort string // HTTP server port (from env: SERVER_PORT)
	LogLevel   string // Logging level (from env: LOG_LEVEL, defaults to "info")
}

// LoadConfig loads config from environment variables and validates required fields.
// Returns a pointer to Config and an error if validation fails.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		DBURL:      os.Getenv("DB_URL"),      // Reads DB_URL from env
		ServerPort: os.Getenv("SERVER_PORT"), // Reads SERVER_PORT from env
		LogLevel:   os.Getenv("LOG_LEVEL"),   // Reads LOG_LEVEL from env
	}

	// Validate required fields
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is required")
	}
	if cfg.ServerPort == "" {
		return nil, fmt.Errorf("SERVER_PORT is required")
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info" // Set default log level if not provided
	}

	return cfg, nil
}
