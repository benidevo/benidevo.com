package models

// Technology represents a technology with its icon
type Technology struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// Project represents a software project
type Project struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	GitHubURL    string       `json:"github_url"`
	LiveURL      string       `json:"live_url,omitempty"`
	Language     string       `json:"language"`
	Technologies []Technology `json:"technologies"`
	Featured     bool         `json:"featured"`
}

// Skill represents a technical skill
type Skill struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Category string `json:"category"`
}

// SkillCategory groups skills by category
type SkillCategory struct {
	Category string  `json:"category"`
	Skills   []Skill `json:"skills"`
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

// ProjectData represents the JSON structure for a project in GitHub data files
type ProjectData struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	GitHubURL    string   `json:"github_url"`
	LiveURL      string   `json:"live_url"`
	Language     string   `json:"language"`
	Technologies []string `json:"technologies"`
	Featured     bool     `json:"featured"`
}

// ProjectsResponse represents the structure of projects.json from GitHub
type ProjectsResponse struct {
	Projects []ProjectData `json:"projects"`
}

// TechnologiesResponse represents the structure of technologies.json from GitHub
type TechnologiesResponse struct {
	Technologies map[string]Technology `json:"technologies"`
}

// SkillsResponse represents the structure of skills.json from GitHub
type SkillsResponse struct {
	SkillCategories []SkillCategory `json:"skill_categories"`
}
