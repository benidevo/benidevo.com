package repository

import (
	"testing"

	"github.com/benidevo/website/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewGitHubRepositories(t *testing.T) {
	cfg := &config.GitHubConfig{
		Owner:      "test-owner",
		Repository: "test-repo",
		Token:      "test-token",
		BaseURL:    "https://api.github.com",
	}

	t.Run("creates project repository", func(t *testing.T) {
		techRepo := NewInMemoryTechnologyRepository()
		repo := NewGitHubProjectRepository(cfg, techRepo)

		assert.NotNil(t, repo)
		assert.Equal(t, cfg, repo.cfg)
		assert.NotNil(t, repo.githubClient)
		assert.Equal(t, techRepo, repo.techRepo)
	})

	t.Run("creates skill repository", func(t *testing.T) {
		repo := NewGitHubSkillRepository(cfg)

		assert.NotNil(t, repo)
		assert.Equal(t, cfg, repo.cfg)
		assert.NotNil(t, repo.githubClient)
		assert.False(t, repo.initialized)
	})

	t.Run("creates technology repository", func(t *testing.T) {
		repo := NewGitHubTechnologyRepository(cfg)

		assert.NotNil(t, repo)
		assert.Equal(t, cfg, repo.cfg)
		assert.NotNil(t, repo.githubClient)
	})
}
