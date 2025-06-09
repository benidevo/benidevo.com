package config

import (
	"github.com/joho/godotenv"
)

// Config bundles all configuration for the application
type Config struct {
	Settings *Settings
}

// SetupConfig initializes and returns application configuration
func SetupConfig() (*Config, error) {
	// Load environment variables from .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	settings := NewSettings()

	InitializeLogger(settings.IsDevelopment, settings.LogLevel)

	return &Config{
		Settings: settings,
	}, nil
}
