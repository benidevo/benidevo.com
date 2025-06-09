package services

import (
	"errors"
	"testing"

	"github.com/benidevo/website/internal/models"
	"github.com/stretchr/testify/assert"
)

// Mock implementation
type mockTechnologyRepository struct {
	technologies map[string]models.Technology
	err          error
}

func (m *mockTechnologyRepository) GetTechnology(name string) (*models.Technology, error) {
	if m.err != nil {
		return nil, m.err
	}
	if tech, ok := m.technologies[name]; ok {
		return &tech, nil
	}
	return nil, errors.New("technology not found")
}

func (m *mockTechnologyRepository) GetTechnologies(names []string) []models.Technology {
	result := make([]models.Technology, 0, len(names))
	for _, name := range names {
		if tech, ok := m.technologies[name]; ok {
			result = append(result, tech)
		} else {
			result = append(result, models.Technology{Name: name, Icon: "⚙️"})
		}
	}
	return result
}

func (m *mockTechnologyRepository) GetAllTechnologies() (map[string]models.Technology, error) {
	return m.technologies, m.err
}

func TestTechnologyService_GetTechnology(t *testing.T) {
	tests := []struct {
		name         string
		techName     string
		technologies map[string]models.Technology
		err          error
		want         models.Technology
	}{
		{
			name:     "returns technology when found",
			techName: "Go",
			technologies: map[string]models.Technology{
				"Go": {Name: "Go", Icon: "🐹"},
			},
			want: models.Technology{Name: "Go", Icon: "🐹"},
		},
		{
			name:     "returns default technology when not found",
			techName: "Unknown",
			technologies: map[string]models.Technology{
				"Go": {Name: "Go", Icon: "🐹"},
			},
			want: models.Technology{Name: "Unknown", Icon: "⚙️"},
		},
		{
			name:     "returns default technology on error",
			techName: "Go",
			err:      errors.New("db error"),
			want:     models.Technology{Name: "Go", Icon: "⚙️"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockTechnologyRepository{
				technologies: tt.technologies,
				err:          tt.err,
			}
			service := NewTechnologyService(mockRepo)

			result := service.GetTechnology(tt.techName)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTechnologyService_GetTechnologies(t *testing.T) {
	tests := []struct {
		name         string
		techNames    []string
		technologies map[string]models.Technology
		want         []models.Technology
	}{
		{
			name:      "returns all technologies found",
			techNames: []string{"Go", "React"},
			technologies: map[string]models.Technology{
				"Go":    {Name: "Go", Icon: "🐹"},
				"React": {Name: "React", Icon: "⚛️"},
			},
			want: []models.Technology{
				{Name: "Go", Icon: "🐹"},
				{Name: "React", Icon: "⚛️"},
			},
		},
		{
			name:      "returns default for unknown technologies",
			techNames: []string{"Go", "Unknown"},
			technologies: map[string]models.Technology{
				"Go": {Name: "Go", Icon: "🐹"},
			},
			want: []models.Technology{
				{Name: "Go", Icon: "🐹"},
				{Name: "Unknown", Icon: "⚙️"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockTechnologyRepository{
				technologies: tt.technologies,
			}
			service := NewTechnologyService(mockRepo)

			result := service.GetTechnologies(tt.techNames)

			assert.Equal(t, tt.want, result)
		})
	}
}
