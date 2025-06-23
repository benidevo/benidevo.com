package repository

import (
	"github.com/benidevo/website/internal/models"
)

// InMemoryProjectRepository implements ProjectRepository with in-memory data
type InMemoryProjectRepository struct {
	projects map[int]*models.Project
}

// NewInMemoryProjectRepository creates a new in-memory project repository
func NewInMemoryProjectRepository(techRepo TechnologyRepository) *InMemoryProjectRepository {
	return &InMemoryProjectRepository{
		projects: make(map[int]*models.Project),
	}
}

// GetAllProjects returns all projects
func (r *InMemoryProjectRepository) GetAllProjects() ([]*models.Project, error) {
	var projects []*models.Project
	for _, project := range r.projects {
		projects = append(projects, project)
	}
	return projects, nil
}

// InMemorySkillRepository implements SkillRepository with in-memory data
type InMemorySkillRepository struct {
	skillCategories []models.SkillCategory
}

// NewInMemorySkillRepository creates a new in-memory skill repository
func NewInMemorySkillRepository(techRepo TechnologyRepository) *InMemorySkillRepository {
	return &InMemorySkillRepository{
		skillCategories: []models.SkillCategory{},
	}
}

// GetSkillCategories returns all skill categories
func (r *InMemorySkillRepository) GetSkillCategories() ([]models.SkillCategory, error) {
	// Return a copy to prevent external modification
	categories := make([]models.SkillCategory, len(r.skillCategories))
	copy(categories, r.skillCategories)
	return categories, nil
}
