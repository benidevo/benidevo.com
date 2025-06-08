package models

import (
	"time"
)

// Technology represents a technology with its icon
type Technology struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Project represents a software project with GitHub integration
type Project struct {
	ID           string       `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	GitHubURL    string       `json:"github_url"`
	LiveURL      string       `json:"live_url,omitempty"`
	Language     string       `json:"language"`
	Stars        int          `json:"stars"`
	Forks        int          `json:"forks"`
	LastUpdated  string       `json:"last_updated"`
	Technologies []Technology `json:"technologies"`
	Featured     bool         `json:"featured"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// ProjectDetails extends Project with more detailed information
type ProjectDetails struct {
	Project
	README          string              `json:"readme"`
	Architecture    string              `json:"architecture"`
	KeyFeatures     []string            `json:"key_features"`
	TechnicalSpecs  []TechnicalSpec     `json:"technical_specs"`
	PerformanceData *PerformanceMetrics `json:"performance_data,omitempty"`
	Challenges      []Challenge         `json:"challenges"`
	Learnings       []string            `json:"learnings"`
}

// TechnicalSpec represents technical specifications for a project
type TechnicalSpec struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

// PerformanceMetrics holds performance data for projects
type PerformanceMetrics struct {
	ResponseTime   string `json:"response_time"`
	Throughput     string `json:"throughput"`
	Availability   string `json:"availability"`
	MemoryUsage    string `json:"memory_usage"`
	CPUUtilization string `json:"cpu_utilization"`
}

// Challenge represents a technical challenge faced during development
type Challenge struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Solution    string `json:"solution"`
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

// HomePageData represents all data needed for the home page
type HomePageData struct {
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	CanonicalURL     string          `json:"canonical_url"`
	CurrentYear      int             `json:"current_year"`
	FeaturedProjects []Project       `json:"featured_projects"`
	SkillCategories  []SkillCategory `json:"skill_categories"`
}

// ProjectFilter represents filtering options for projects
type ProjectFilter struct {
	Language   string
	Featured   *bool
	Technology string
	Limit      int
	Offset     int
}
