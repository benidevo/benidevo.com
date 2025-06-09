package repository

import (
	"fmt"

	"github.com/benidevo/website/internal/models"
)

// InMemoryTechnologyRepository implements TechnologyRepository with in-memory data
type InMemoryTechnologyRepository struct {
	technologies map[string]models.Technology
}

// NewInMemoryTechnologyRepository creates a new in-memory technology repository
func NewInMemoryTechnologyRepository() *InMemoryTechnologyRepository {
	return &InMemoryTechnologyRepository{
		technologies: make(map[string]models.Technology),
	}
}

// GetTechnology returns a technology by name
func (r *InMemoryTechnologyRepository) GetTechnology(name string) (*models.Technology, error) {
	tech, exists := r.technologies[name]
	if !exists {
		return nil, fmt.Errorf("technology '%s' not found", name)
	}
	return &tech, nil
}

// GetTechnologies returns multiple technologies by names (optimized for batch lookups)
func (r *InMemoryTechnologyRepository) GetTechnologies(names []string) []models.Technology {
	// Pre-allocate slice with exact capacity to avoid reallocations
	techs := make([]models.Technology, 0, len(names))
	for _, name := range names {
		if tech, exists := r.technologies[name]; exists {
			techs = append(techs, tech)
		}
	}
	return techs
}

// GetAllTechnologies returns all available technologies
func (r *InMemoryTechnologyRepository) GetAllTechnologies() (map[string]models.Technology, error) {
	// Return a copy to prevent external modification
	result := make(map[string]models.Technology)
	for k, v := range r.technologies {
		result[k] = v
	}
	return result, nil
}

// AddTechnology adds a technology to the in-memory store (for testing)
func (r *InMemoryTechnologyRepository) AddTechnology(name string, technology models.Technology) {
	r.technologies[name] = technology
}
