package repository

import (
	"github.com/benidevo/website/internal/models"
)

// ProjectRepository defines the interface for project data access
type ProjectRepository interface {
	// GetAllProjects returns all projects
	GetAllProjects() ([]*models.Project, error)

	// GetProjectByID returns a project by its ID (includes all details)
	GetProjectByID(id int) (*models.Project, error)
}

// TechnologyRepository defines the interface for technology data access
type TechnologyRepository interface {
	// GetTechnology returns a technology by name
	GetTechnology(name string) (*models.Technology, error)

	// GetTechnologies returns multiple technologies by names
	GetTechnologies(names []string) []models.Technology

	// GetAllTechnologies returns all available technologies
	GetAllTechnologies() (map[string]models.Technology, error)
}

// SkillRepository defines the interface for skill data access
type SkillRepository interface {
	// GetSkillCategories returns all skill categories
	GetSkillCategories() ([]models.SkillCategory, error)
}
