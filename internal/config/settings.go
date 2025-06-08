package config

import "os"

type Settings struct {
	Port          string `json:"port" env:"PORT" default:"8080"`
	IsDevelopment bool   `json:"is_development" env:"IS_DEVELOPMENT" default:"true"`
	LogLevel      string `json:"log_level" env:"LOG_LEVEL" default:"info"`
}

func NewSettings() *Settings {
	return &Settings{
		Port:          getEnv("PORT", "8080"),
		IsDevelopment: getEnv("IS_DEVELOPMENT", "true") == "true",
		LogLevel:      getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
