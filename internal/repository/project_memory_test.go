package repository

import (
	"testing"

	"github.com/benidevo/website/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryProjectRepository_GetAllProjects(t *testing.T) {
	tests := []struct {
		name     string
		projects map[int]*models.Project
		want     int
	}{
		{
			name: "returns all projects",
			projects: map[int]*models.Project{
				1: {ID: 1, Title: "Project 1"},
				2: {ID: 2, Title: "Project 2"},
			},
			want: 2,
		},
		{
			name:     "returns empty slice when no projects",
			projects: map[int]*models.Project{},
			want:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryProjectRepository{
				projects: tt.projects,
			}

			result, err := repo.GetAllProjects()

			assert.NoError(t, err)
			assert.Len(t, result, tt.want)
		})
	}
}

func TestInMemorySkillRepository_GetSkillCategories(t *testing.T) {
	tests := []struct {
		name       string
		categories []models.SkillCategory
		want       int
	}{
		{
			name: "returns all categories",
			categories: []models.SkillCategory{
				{Category: "Backend", Skills: []models.Skill{{Name: "Go"}}},
				{Category: "Frontend", Skills: []models.Skill{{Name: "React"}}},
			},
			want: 2,
		},
		{
			name:       "returns empty slice when no categories",
			categories: []models.SkillCategory{},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemorySkillRepository{
				skillCategories: tt.categories,
			}

			result, err := repo.GetSkillCategories()

			assert.NoError(t, err)
			assert.Len(t, result, tt.want)
			// Verify it returns a copy
			if tt.want > 0 {
				result[0].Category = "Modified"
				assert.NotEqual(t, result[0].Category, repo.skillCategories[0].Category)
			}
		})
	}
}
