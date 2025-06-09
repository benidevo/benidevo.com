package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	Port          string       `json:"port" env:"PORT" default:"8080"`
	IsDevelopment bool         `json:"is_development" env:"IS_DEVELOPMENT" default:"true"`
	LogLevel      string       `json:"log_level" env:"LOG_LEVEL" default:"info"`
	GitHub        GitHubConfig `json:"github"`
}

type GitHubConfig struct {
	Owner      string `json:"owner" env:"GITHUB_OWNER"`
	Repository string `json:"repository" env:"GITHUB_REPOSITORY"`
	Token      string `json:"token" env:"GITHUB_TOKEN"`
	BaseURL    string `json:"base_url" env:"GITHUB_BASE_URL" default:"https://api.github.com"`
}

func NewSettings() *Settings {
	if err := godotenv.Load(); err != nil {
		log.Debug().Err(err).Msg("No .env file found... \nusing environment variables only")
	}

	return &Settings{
		Port:          getEnv("PORT", "8080"),
		IsDevelopment: getEnv("IS_DEVELOPMENT", "true") == "true",
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		GitHub: GitHubConfig{
			Owner:      getEnv("GITHUB_OWNER", ""),
			Repository: getEnv("GITHUB_REPOSITORY", ""),
			Token:      getEnv("GITHUB_TOKEN", ""),
			BaseURL:    getEnv("GITHUB_BASE_URL", "https://api.github.com"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
