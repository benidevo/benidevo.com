package repository

import (
	"encoding/json"
	"fmt"
	"strings"

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
		detailedDescription = data.Description
	}

	technologies := r.techRepo.GetTechnologies(data.Technologies)

	project := &models.Project{
		ID:                     data.ID,
		Title:                  data.Title,
		Description:            data.Description,
		DetailedDescription:    detailedDescription,
		ArchitectureDiagramURL: data.ArchitectureDiagramURL,
		GitHubURL:              data.GitHubURL,
		LiveURL:                data.LiveURL,
		Language:               data.Language,
		Technologies:           technologies,
		Featured:               data.Featured,
	}

	if data.ArchitectureDiagramURL != "" {
		// Extract filename from URL (e.g., "project-architecture.jpg" from the URL)
		parts := strings.Split(data.ArchitectureDiagramURL, "/")
		if len(parts) > 0 {
			filename := parts[len(parts)-1]
			content, err := r.githubClient.FetchBinaryFileContent(fmt.Sprintf("assets/diagrams/%s", filename))
			if err != nil {
				log.Warn().Err(err).Str("diagram", filename).Msg("Failed to fetch architecture diagram")
			} else {
				project.ArchitectureDiagramContent = content
			}
		}
	}

	return project, nil
}
