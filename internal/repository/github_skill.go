package repository

import (
	"encoding/json"
	"fmt"

	"github.com/benidevo/website/internal/client"
	"github.com/benidevo/website/internal/config"
	"github.com/benidevo/website/internal/models"
	"github.com/rs/zerolog/log"
)

// GitHubSkillRepository implements SkillRepository using GitHub API
type GitHubSkillRepository struct {
	cfg             *config.GitHubConfig
	githubClient    *client.GitHubClient
	skillCategories []models.SkillCategory // Cache for skill categories
	initialized     bool                   // Track initialization status
}

// NewGitHubSkillRepository creates a new GitHub-based skill repository
func NewGitHubSkillRepository(cfg *config.GitHubConfig) *GitHubSkillRepository {
	repo := &GitHubSkillRepository{
		cfg:          cfg,
		githubClient: client.NewGitHubClient(cfg),
		initialized:  false,
	}

	go repo.InitializeAsync()

	return repo
}

// GetSkillCategories returns all skill categories
func (r *GitHubSkillRepository) GetSkillCategories() ([]models.SkillCategory, error) {
	if err := r.ensureSkillsLoaded(); err != nil {
		return nil, err
	}

	categories := make([]models.SkillCategory, len(r.skillCategories))
	copy(categories, r.skillCategories)
	return categories, nil
}

// InitializeAsync loads skills asynchronously in the background
func (r *GitHubSkillRepository) InitializeAsync() {
	if err := r.loadSkills(); err != nil {
		log.Error().Err(err).Msg("Failed to load skills asynchronously")
	} else {
		r.initialized = true
		log.Info().Msg("Skills loaded asynchronously")
	}
}

// ensureSkillsLoaded loads skills from GitHub if not already loaded
func (r *GitHubSkillRepository) ensureSkillsLoaded() error {
	if len(r.skillCategories) > 0 {
		return nil
	}

	if !r.initialized {
		log.Warn().Msg("Skills not loaded asynchronously, loading synchronously")
		return r.loadSkills()
	}

	return nil
}

// loadSkills fetches and loads skills from GitHub
func (r *GitHubSkillRepository) loadSkills() error {
	url := "skills/skills.json"

	content, err := r.githubClient.FetchFileContent(url)
	if err != nil {
		return fmt.Errorf("failed to fetch skills.json: %w", err)
	}

	var skillsResponse models.SkillsResponse
	if err := json.Unmarshal([]byte(content), &skillsResponse); err != nil {
		return fmt.Errorf("failed to parse skills.json: %w", err)
	}

	r.skillCategories = skillsResponse.SkillCategories
	log.Info().Int("categories", len(r.skillCategories)).Msg("Loaded skill categories from GitHub")

	return nil
}

// RefreshSkills forces a reload of skills from GitHub
func (r *GitHubSkillRepository) RefreshSkills() error {
	r.skillCategories = nil
	return r.loadSkills()
}
