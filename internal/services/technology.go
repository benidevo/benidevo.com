package services

import (
	"github.com/benidevo/website/internal/models"
	"github.com/benidevo/website/internal/repository"
)

// TechnologyService manages technology icons and metadata
type TechnologyService struct {
	repo repository.TechnologyRepository
}

// NewTechnologyService creates a new technology service
func NewTechnologyService(repo repository.TechnologyRepository) *TechnologyService {
	return &TechnologyService{
		repo: repo,
	}
}

// GetTechnology returns technology with icon by name
func (t *TechnologyService) GetTechnology(name string) models.Technology {
	tech, err := t.repo.GetTechnology(name)
	if err != nil {
		return models.Technology{
			Name: name,
			Icon: "⚙️", // Default gear icon
		}
	}
	return *tech
}

// GetTechnologies converts a slice of strings to Technology objects
func (t *TechnologyService) GetTechnologies(names []string) []models.Technology {
	return t.repo.GetTechnologies(names)
}
