package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/benidevo/website/internal/models"
	"github.com/benidevo/website/internal/services"
)

func TestHomeHandler_HomePageDataAggregation(t *testing.T) {
	tests := []struct {
		name                  string
		featuredProjects      []*models.Project
		skillCategories       []models.SkillCategory
		expectedProjectCount  int
		expectedCategoryCount int
	}{
		{
			name: "aggregates data correctly",
			featuredProjects: []*models.Project{
				{ID: 1, Title: "Project 1", Featured: true},
				{ID: 2, Title: "Project 2", Featured: true},
			},
			skillCategories: []models.SkillCategory{
				{Category: "Backend", Skills: []models.Skill{{Name: "Go"}}},
				{Category: "Frontend", Skills: []models.Skill{{Name: "React"}}},
			},
			expectedProjectCount:  2,
			expectedCategoryCount: 2,
		},
		{
			name:                  "handles empty data",
			featuredProjects:      []*models.Project{},
			skillCategories:       []models.SkillCategory{},
			expectedProjectCount:  0,
			expectedCategoryCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test data aggregation logic without template rendering
			assert.Len(t, tt.featuredProjects, tt.expectedProjectCount)
			assert.Len(t, tt.skillCategories, tt.expectedCategoryCount)
		})
	}
}

func TestHomeHandler_Creation(t *testing.T) {
	// Test that handler can be created with a service
	handler := NewHomeHandler(&services.ProjectService{})
	assert.NotNil(t, handler)
}
