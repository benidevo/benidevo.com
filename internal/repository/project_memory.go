package repository

import (
	"fmt"

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

// GetProjectByID returns a project by its ID
func (r *InMemoryProjectRepository) GetProjectByID(id int) (*models.Project, error) {
	project, exists := r.projects[id]
	if !exists {
		return nil, fmt.Errorf("project with ID %d not found", id)
	}
	return project, nil
}

// AddProject adds a project to the in-memory store (for testing)
func (r *InMemoryProjectRepository) AddProject(project *models.Project) {
	r.projects[project.ID] = project
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

// AddSkillCategory adds a skill category to the in-memory store (for testing)
func (r *InMemorySkillRepository) AddSkillCategory(category models.SkillCategory) {
	r.skillCategories = append(r.skillCategories, category)
}
