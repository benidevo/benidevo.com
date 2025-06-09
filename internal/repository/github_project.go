package repository

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/benidevo/website/internal/client"
	"github.com/benidevo/website/internal/config"
	"github.com/benidevo/website/internal/models"
	"github.com/rs/zerolog/log"
)

// GitHubProjectRepository implements ProjectRepository using GitHub API
type GitHubProjectRepository struct {
	cfg          *config.GitHubConfig
	githubClient *client.GitHubClient
	techRepo     TechnologyRepository
}

// NewGitHubProjectRepository creates a new GitHub-based project repository
func NewGitHubProjectRepository(cfg *config.GitHubConfig, techRepo TechnologyRepository) *GitHubProjectRepository {
	repo := &GitHubProjectRepository{
		cfg:          cfg,
		githubClient: client.NewGitHubClient(cfg),
		techRepo:     techRepo,
	}

	go repo.PrewarmCache()

	return repo
}

// GetAllProjects fetches all projects from GitHub repository
func (r *GitHubProjectRepository) GetAllProjects() ([]*models.Project, error) {
	projectsData, err := r.fetchProjectsData()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch projects data: %w", err)
	}

	var projects []*models.Project
	for _, projectData := range projectsData.Projects {
		project, err := r.convertToProject(projectData)
		if err != nil {
			log.Error().Err(err).Int("projectID", projectData.ID).Msg("Failed to convert project data")
			continue
		}
		projects = append(projects, project)
	}

	return projects, nil
}

// GetProjectByID fetches a specific project by ID from GitHub repository
func (r *GitHubProjectRepository) GetProjectByID(id int) (*models.Project, error) {
	projects, err := r.GetAllProjects()
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if project.ID == id {
			return project, nil
		}
	}

	return nil, fmt.Errorf("project with ID %d not found", id)
}

// fetchProjectsData fetches and parses projects.json from GitHub
func (r *GitHubProjectRepository) fetchProjectsData() (*models.ProjectsResponse, error) {
	content, err := r.githubClient.FetchFileContent("projects/projects.json")
	if err != nil {
		return nil, err
	}

	var projectsResponse models.ProjectsResponse
	if err := json.Unmarshal([]byte(content), &projectsResponse); err != nil {
		return nil, fmt.Errorf("failed to parse projects.json: %w", err)
	}

	return &projectsResponse, nil
}

// PrewarmCache loads project data asynchronously to warm up the cache
func (r *GitHubProjectRepository) PrewarmCache() {
	log.Info().Msg("Pre-warming project cache asynchronously")
	if _, err := r.GetAllProjects(); err != nil {
		log.Error().Err(err).Msg("Failed to pre-warm project cache")
	} else {
		log.Info().Msg("Project cache pre-warmed successfully")
	}
}

// fetchDetailedDescription fetches the detailed description markdown file
func (r *GitHubProjectRepository) fetchDetailedDescription(filename string) (string, error) {
	return r.githubClient.FetchFileContent(fmt.Sprintf("projects/details/%s", filename))
}

// convertToProject converts ProjectData to models.Project
func (r *GitHubProjectRepository) convertToProject(data models.ProjectData) (*models.Project, error) {
	detailedDescription, err := r.fetchDetailedDescription(data.DetailedDescriptionFile)
	if err != nil {
		log.Warn().Err(err).Str("file", data.DetailedDescriptionFile).Msg("Failed to fetch detailed description")
		detailedDescription = data.Description // Fallback to basic description
	}

	technologies := r.techRepo.GetTechnologies(data.Technologies)

	// Parse last updated time
	var lastUpdated time.Time
	if data.LastUpdated != "" {
		if parsed, err := time.Parse(time.RFC3339, data.LastUpdated); err == nil {
			lastUpdated = parsed
		}
	}

	return &models.Project{
		ID:                     data.ID,
		Title:                  data.Title,
		Description:            data.Description,
		DetailedDescription:    detailedDescription,
		ArchitectureDiagramURL: data.ArchitectureDiagramURL,
		GitHubURL:              data.GitHubURL,
		LiveURL:                data.LiveURL,
		Language:               data.Language,
		Stars:                  data.Stars,
		Forks:                  data.Forks,
		LastUpdated:            data.LastUpdated,
		Technologies:           technologies,
		Featured:               data.Featured,
		UpdatedAt:              lastUpdated,
	}, nil
}

// EnrichWithGitHubData fetches live data from GitHub API for a repository
func (r *GitHubProjectRepository) EnrichWithGitHubData(project *models.Project) error {
	if project.GitHubURL == "" {
		return nil // No GitHub URL, nothing to enrich
	}

	// Extract owner and repo from GitHub URL
	owner, repo, err := r.parseGitHubURL(project.GitHubURL)
	if err != nil {
		return fmt.Errorf("failed to parse GitHub URL: %w", err)
	}

	// Fetch repository data from GitHub API
	repoData, err := r.fetchGitHubRepoData(owner, repo)
	if err != nil {
		log.Warn().Err(err).Str("repo", project.GitHubURL).Msg("Failed to fetch GitHub repository data")
		return nil // Don't fail the entire operation, just log the warning
	}

	// Update project with live data
	project.Stars = repoData.StargazersCount
	project.Forks = repoData.ForksCount
	if !repoData.UpdatedAt.IsZero() {
		project.LastUpdated = repoData.UpdatedAt.Format(time.RFC3339)
		project.UpdatedAt = repoData.UpdatedAt
	}

	return nil
}

// parseGitHubURL extracts owner and repository name from GitHub URL
//
// githubURL format: https://github.com/owner/repo
func (r *GitHubProjectRepository) parseGitHubURL(githubURL string) (string, string, error) {
	githubURL = strings.TrimSuffix(githubURL, "/")
	parts := strings.Split(githubURL, "/")

	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid GitHub URL format: %s", githubURL)
	}

	repo := parts[len(parts)-1]
	owner := parts[len(parts)-2]

	return owner, repo, nil
}

// fetchGitHubRepoData fetches repository metadata from GitHub API
func (r *GitHubProjectRepository) fetchGitHubRepoData(owner, repo string) (*models.GitHubRepository, error) {
	data, err := r.githubClient.FetchRepositoryData(owner, repo)
	if err != nil {
		return nil, err
	}

	var repoData models.GitHubRepository
	if err := json.Unmarshal(data, &repoData); err != nil {
		return nil, fmt.Errorf("failed to decode repository data: %w", err)
	}

	return &repoData, nil
}
