package services

import (
	"github.com/benidevo/website/internal/models"
	"github.com/benidevo/website/internal/repository"
	"github.com/rs/zerolog/log"
)

// ProjectService provides methods to manage projects, integrating with repositories.
type ProjectService struct {
	projectRepo repository.ProjectRepository
	skillRepo   repository.SkillRepository
}

// NewProjectService creates a new project service
func NewProjectService(projectRepo repository.ProjectRepository, skillRepo repository.SkillRepository) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
		skillRepo:   skillRepo,
	}
}

// GetFeaturedProjects returns all projects (since all are featured)
func (p *ProjectService) GetFeaturedProjects() []*models.Project {
	projects, err := p.projectRepo.GetAllProjects()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get projects")
		return []*models.Project{}
	}
	return projects
}

// GetSkillCategories returns skill categories for the home page
func (p *ProjectService) GetSkillCategories() []models.SkillCategory {
	categories, err := p.skillRepo.GetSkillCategories()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get skill categories")
		return []models.SkillCategory{}
	}
	return categories
}
