package models

import (
	"time"
)

// Technology represents a technology with its icon
type Technology struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Project represents a software project
type Project struct {
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	DetailedDescription    string       `json:"detailed_description"`
	ArchitectureDiagramURL string       `json:"architecture_diagram_url,omitempty"`
	GitHubURL              string       `json:"github_url"`
	LiveURL                string       `json:"live_url,omitempty"`
	Language               string       `json:"language"`
	Stars                  int          `json:"stars"`
	Forks                  int          `json:"forks"`
	LastUpdated            string       `json:"last_updated"`
	Technologies           []Technology `json:"technologies"`
	Featured               bool         `json:"featured"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
}

// Skill represents a technical skill with proficiency level
type Skill struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Proficiency int    `json:"proficiency"` // 0-100
	Years       int    `json:"years"`
	Category    string `json:"category"`
}

// SkillCategory groups skills by category
type SkillCategory struct {
	Category string  `json:"category"`
	Skills   []Skill `json:"skills"`
}

// GitHubRepository represents data from GitHub API
type GitHubRepository struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Description     string    `json:"description"`
	HTMLURL         string    `json:"html_url"`
	Language        string    `json:"language"`
	StargazersCount int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	UpdatedAt       time.Time `json:"updated_at"`
	Topics          []string  `json:"topics"`
	Private         bool      `json:"private"`
}

// ProjectDetailPageData represents all data needed for a project detail page
type ProjectDetailPageData struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	CanonicalURL string   `json:"canonical_url"`
	CurrentYear  int      `json:"current_year"`
	Project      *Project `json:"project"`
}

// HomePageData represents all data needed for the home page
type HomePageData struct {
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	CanonicalURL     string          `json:"canonical_url"`
	CurrentYear      int             `json:"current_year"`
	FeaturedProjects []*Project      `json:"featured_projects"`
	SkillCategories  []SkillCategory `json:"skill_categories"`
}
