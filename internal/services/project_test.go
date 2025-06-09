package services

import (
	"errors"
	"testing"

	"github.com/benidevo/website/internal/models"
	"github.com/stretchr/testify/assert"
)

// Mock implementations
type mockProjectRepository struct {
	projects []*models.Project
	err      error
}

func (m *mockProjectRepository) GetAllProjects() ([]*models.Project, error) {
	return m.projects, m.err
}

func (m *mockProjectRepository) GetProjectByID(id int) (*models.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, p := range m.projects {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("project not found")
}

type mockSkillRepository struct {
	categories []models.SkillCategory
	err        error
}

func (m *mockSkillRepository) GetSkillCategories() ([]models.SkillCategory, error) {
	return m.categories, m.err
}

func TestProjectService_GetFeaturedProjects(t *testing.T) {
	tests := []struct {
		name     string
		projects []*models.Project
		err      error
		want     int
	}{
		{
			name: "returns projects successfully",
			projects: []*models.Project{
				{ID: 1, Title: "Project 1"},
				{ID: 2, Title: "Project 2"},
			},
			want: 2,
		},
		{
			name: "returns empty slice on error",
			err:  errors.New("db error"),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockProjectRepository{
				projects: tt.projects,
				err:      tt.err,
			}
			service := NewProjectService(mockRepo, nil)

			result := service.GetFeaturedProjects()

			assert.Len(t, result, tt.want)
			if tt.err == nil && tt.want > 0 {
				assert.Equal(t, tt.projects, result)
			}
		})
	}
}

func TestProjectService_GetProjectByID(t *testing.T) {
	tests := []struct {
		name      string
		projectID int
		projects  []*models.Project
		err       error
		want      *models.Project
	}{
		{
			name:      "returns project when found",
			projectID: 1,
			projects: []*models.Project{
				{ID: 1, Title: "Project 1"},
				{ID: 2, Title: "Project 2"},
			},
			want: &models.Project{ID: 1, Title: "Project 1"},
		},
		{
			name:      "returns nil when not found",
			projectID: 99,
			projects: []*models.Project{
				{ID: 1, Title: "Project 1"},
			},
			want: nil,
		},
		{
			name:      "returns nil on repository error",
			projectID: 1,
			err:       errors.New("db error"),
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockProjectRepository{
				projects: tt.projects,
				err:      tt.err,
			}
			service := NewProjectService(mockRepo, nil)

			result := service.GetProjectByID(tt.projectID)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestProjectService_GetSkillCategories(t *testing.T) {
	tests := []struct {
		name       string
		categories []models.SkillCategory
		err        error
		want       int
	}{
		{
			name: "returns categories successfully",
			categories: []models.SkillCategory{
				{Category: "Backend", Skills: []models.Skill{{Name: "Go"}}},
				{Category: "Frontend", Skills: []models.Skill{{Name: "React"}}},
			},
			want: 2,
		},
		{
			name: "returns empty slice on error",
			err:  errors.New("db error"),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSkillRepo := &mockSkillRepository{
				categories: tt.categories,
				err:        tt.err,
			}
			service := NewProjectService(nil, mockSkillRepo)

			result := service.GetSkillCategories()

			assert.Len(t, result, tt.want)
			if tt.err == nil && tt.want > 0 {
				assert.Equal(t, tt.categories, result)
			}
		})
	}
}
