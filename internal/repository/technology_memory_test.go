package repository

import (
	"testing"

	"github.com/benidevo/website/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryTechnologyRepository_GetTechnology(t *testing.T) {
	tests := []struct {
		name         string
		technologies map[string]models.Technology
		techName     string
		want         *models.Technology
		wantErr      bool
	}{
		{
			name: "returns technology when found",
			technologies: map[string]models.Technology{
				"Go": {Name: "Go", Icon: "🐹"},
			},
			techName: "Go",
			want:     &models.Technology{Name: "Go", Icon: "🐹"},
			wantErr:  false,
		},
		{
			name: "returns error when not found",
			technologies: map[string]models.Technology{
				"Go": {Name: "Go", Icon: "🐹"},
			},
			techName: "Python",
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryTechnologyRepository{
				technologies: tt.technologies,
			}

			result, err := repo.GetTechnology(tt.techName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

func TestInMemoryTechnologyRepository_GetTechnologies(t *testing.T) {
	tests := []struct {
		name         string
		technologies map[string]models.Technology
		techNames    []string
		want         []models.Technology
	}{
		{
			name: "returns only found technologies",
			technologies: map[string]models.Technology{
				"Go":    {Name: "Go", Icon: "🐹"},
				"React": {Name: "React", Icon: "⚛️"},
			},
			techNames: []string{"Go", "React", "Python"},
			want: []models.Technology{
				{Name: "Go", Icon: "🐹"},
				{Name: "React", Icon: "⚛️"},
			},
		},
		{
			name:         "returns empty slice when none found",
			technologies: map[string]models.Technology{},
			techNames:    []string{"Go", "React"},
			want:         []models.Technology{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryTechnologyRepository{
				technologies: tt.technologies,
			}

			result := repo.GetTechnologies(tt.techNames)

			assert.Len(t, result, len(tt.want))
			for i, tech := range tt.want {
				assert.Equal(t, tech, result[i])
			}
		})
	}
}

func TestInMemoryTechnologyRepository_GetAllTechnologies(t *testing.T) {
	tests := []struct {
		name         string
		technologies map[string]models.Technology
		want         int
	}{
		{
			name: "returns all technologies",
			technologies: map[string]models.Technology{
				"Go":    {Name: "Go", Icon: "🐹"},
				"React": {Name: "React", Icon: "⚛️"},
			},
			want: 2,
		},
		{
			name:         "returns empty map when no technologies",
			technologies: map[string]models.Technology{},
			want:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryTechnologyRepository{
				technologies: tt.technologies,
			}

			result, err := repo.GetAllTechnologies()

			assert.NoError(t, err)
			assert.Len(t, result, tt.want)
			// Verify it returns a copy
			if tt.want > 0 {
				result["Go"] = models.Technology{Name: "Modified", Icon: "X"}
				assert.NotEqual(t, result["Go"], repo.technologies["Go"])
			}
		})
	}
}
