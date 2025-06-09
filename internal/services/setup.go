package services

import (
	"fmt"

	"github.com/benidevo/website/internal/config"
	"github.com/benidevo/website/internal/repository"
)

// Services bundles all application services
type Services struct {
	ProjectService    *ProjectService
	TechnologyService *TechnologyService
}

// SetupServices initializes and returns all application services with their dependencies
func SetupServices(cfg *config.Config) (*Services, error) {
	repos, err := setupRepositories(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup repositories: %w", err)
	}

	// Create services with repository dependencies
	projectService := NewProjectService(repos.ProjectRepo, repos.SkillRepo)
	technologyService := NewTechnologyService(repos.TechnologyRepo)

	return &Services{
		ProjectService:    projectService,
		TechnologyService: technologyService,
	}, nil
}

// repositoryBundle groups all repositories for internal use
type repositoryBundle struct {
	ProjectRepo    repository.ProjectRepository
	TechnologyRepo repository.TechnologyRepository
	SkillRepo      repository.SkillRepository
}

// setupRepositories creates and configures all repositories
func setupRepositories(cfg *config.Config) (*repositoryBundle, error) {
	var (
		projectRepo    repository.ProjectRepository
		technologyRepo repository.TechnologyRepository
		skillRepo      repository.SkillRepository
	)

	// Determine repository implementation based on GitHub configuration
	if cfg.Settings.GitHub.Token != "" {
		technologyRepo = repository.NewGitHubTechnologyRepository(&cfg.Settings.GitHub)
		skillRepo = repository.NewGitHubSkillRepository(&cfg.Settings.GitHub)
		projectRepo = repository.NewGitHubProjectRepository(&cfg.Settings.GitHub, technologyRepo)
	} else {
		technologyRepo = repository.NewInMemoryTechnologyRepository()
		skillRepo = repository.NewInMemorySkillRepository(technologyRepo)
		projectRepo = repository.NewInMemoryProjectRepository(technologyRepo)
	}

	return &repositoryBundle{
		ProjectRepo:    projectRepo,
		TechnologyRepo: technologyRepo,
		SkillRepo:      skillRepo,
	}, nil
}
