package repository

import (
	"encoding/json"
	"fmt"

	"github.com/benidevo/website/internal/client"
	"github.com/benidevo/website/internal/config"
	"github.com/benidevo/website/internal/models"
	"github.com/rs/zerolog/log"
)

// GitHubTechnologyRepository implements TechnologyRepository using GitHub API
type GitHubTechnologyRepository struct {
	cfg          *config.GitHubConfig
	githubClient *client.GitHubClient
	technologies map[string]models.Technology // Cache for technologies
	initialized  bool                         // Track initialization status
}

// NewGitHubTechnologyRepository creates a new GitHub-based technology repository
func NewGitHubTechnologyRepository(cfg *config.GitHubConfig) *GitHubTechnologyRepository {
	repo := &GitHubTechnologyRepository{
		cfg:          cfg,
		githubClient: client.NewGitHubClient(cfg),
		technologies: make(map[string]models.Technology),
		initialized:  false,
	}

	go repo.InitializeAsync()

	return repo
}

// GetTechnology returns a technology by name
func (r *GitHubTechnologyRepository) GetTechnology(name string) (*models.Technology, error) {
	if err := r.ensureTechnologiesLoaded(); err != nil {
		return nil, err
	}

	tech, exists := r.technologies[name]
	if !exists {
		return nil, fmt.Errorf("technology '%s' not found", name)
	}
	return &tech, nil
}

// GetTechnologies returns multiple technologies by names
func (r *GitHubTechnologyRepository) GetTechnologies(names []string) []models.Technology {
	if err := r.ensureTechnologiesLoaded(); err != nil {
		log.Error().Err(err).Msg("Failed to load technologies")
		return []models.Technology{}
	}

	techs := make([]models.Technology, 0, len(names))
	for _, name := range names {
		if tech, exists := r.technologies[name]; exists {
			techs = append(techs, tech)
		}
	}
	return techs
}

// GetAllTechnologies returns all available technologies
func (r *GitHubTechnologyRepository) GetAllTechnologies() (map[string]models.Technology, error) {
	if err := r.ensureTechnologiesLoaded(); err != nil {
		return nil, err
	}

	result := make(map[string]models.Technology)
	for k, v := range r.technologies {
		result[k] = v
	}
	return result, nil
}

// InitializeAsync loads technologies asynchronously in the background
func (r *GitHubTechnologyRepository) InitializeAsync() {
	if err := r.loadTechnologies(); err != nil {
		log.Error().Err(err).Msg("Failed to load technologies asynchronously")
	} else {
		r.initialized = true
		log.Info().Msg("Technologies loaded asynchronously")
	}
}

// ensureTechnologiesLoaded loads technologies from GitHub if not already loaded
func (r *GitHubTechnologyRepository) ensureTechnologiesLoaded() error {
	if len(r.technologies) > 0 {
		return nil // Already loaded
	}

	if !r.initialized {
		log.Warn().Msg("Technologies not loaded asynchronously, loading synchronously")
		return r.loadTechnologies()
	}

	return nil
}

// loadTechnologies fetches and loads technologies from GitHub
func (r *GitHubTechnologyRepository) loadTechnologies() error {
	url := fmt.Sprintf("technologies/technologies.json")

	content, err := r.githubClient.FetchFileContent(url)
	if err != nil {
		return fmt.Errorf("failed to fetch technologies.json: %w", err)
	}

	var techResponse models.TechnologiesResponse
	if err := json.Unmarshal([]byte(content), &techResponse); err != nil {
		return fmt.Errorf("failed to parse technologies.json: %w", err)
	}

	r.technologies = techResponse.Technologies
	log.Info().Int("count", len(r.technologies)).Msg("Loaded technologies from GitHub")

	return nil
}

// RefreshTechnologies forces a reload of technologies from GitHub
func (r *GitHubTechnologyRepository) RefreshTechnologies() error {
	r.technologies = make(map[string]models.Technology) // Clear cache
	return r.loadTechnologies()
}
