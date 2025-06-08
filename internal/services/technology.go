package services

import (
	"github.com/benidevo/website/internal/models"
)

// TechnologyService manages technology icons and metadata
type TechnologyService struct {
	technologies map[string]models.Technology
}

// NewTechnologyService creates a new technology service
func NewTechnologyService() *TechnologyService {
	return &TechnologyService{
		technologies: getTechnologyData(),
	}
}

// GetTechnology returns technology with icon by name
func (t *TechnologyService) GetTechnology(name string) models.Technology {
	if tech, exists := t.technologies[name]; exists {
		return tech
	}

	return models.Technology{
		Name: name,
		Icon: "⚙️", // Default gear icon
	}
}

// GetTechnologies converts a slice of strings to Technology objects
func (t *TechnologyService) GetTechnologies(names []string) []models.Technology {
	var technologies []models.Technology
	for _, name := range names {
		technologies = append(technologies, t.GetTechnology(name))
	}
	return technologies
}
