package services

import (
	"fmt"
	"sort"

	"github.com/benidevo/website/internal/models"
	"maps"
)

// ProjectService provides methods to manage projects, integrating with GitHub and technology services.
// It maintains a map of project details for quick access.
type ProjectService struct {
	githubService     *GitHubService
	technologyService *TechnologyService
	projects          map[string]*models.ProjectDetails
}

// NewProjectService creates a new project service
func NewProjectService(githubService *GitHubService) *ProjectService {
	return &ProjectService{
		githubService:     githubService,
		technologyService: NewTechnologyService(),
		projects:          make(map[string]*models.ProjectDetails),
	}
}

// LoadProjects initializes the project data with predefined projects
func (p *ProjectService) LoadProjects() {
	featuredProjects := p.getFeaturedProjectsData()

	maps.Copy(p.projects, featuredProjects)
}

// GetFeaturedProjects returns featured projects
func (p *ProjectService) GetFeaturedProjects() []models.Project {
	if len(p.projects) == 0 {
		p.LoadProjects()
	}

	var featured []models.Project

	for _, project := range p.projects {
		if project.Featured {
			featured = append(featured, project.Project)
		}
	}

	// Sort by last updated (most recent first)
	sort.Slice(featured, func(i, j int) bool {
		return featured[i].UpdatedAt.After(featured[j].UpdatedAt)
	})

	return featured
}

// GetProjectDetails returns detailed information for a specific project
func (p *ProjectService) GetProjectDetails(id string) (*models.ProjectDetails, error) {
	project, exists := p.projects[id]
	if !exists {
		return nil, fmt.Errorf("project not found: %s", id)
	}

	return project, nil
}

// GetSkillCategories returns skill categories for the home page
func (p *ProjectService) GetSkillCategories() []models.SkillCategory {
	return []models.SkillCategory{
		{
			Category: "Languages",
			Skills: []models.Skill{
				{Name: "Go", Icon: p.technologyService.GetTechnology("Go").Icon, Category: "Languages"},
				{Name: "Java", Icon: p.technologyService.GetTechnology("Java").Icon, Category: "Languages"},
				{Name: "Python", Icon: p.technologyService.GetTechnology("Python").Icon, Category: "Languages"},
				{Name: "JavaScript", Icon: p.technologyService.GetTechnology("JavaScript").Icon, Category: "Languages"},
			},
		},
		{
			Category: "Frameworks",
			Skills: []models.Skill{
				{Name: "Spring Boot", Icon: p.technologyService.GetTechnology("Spring Boot").Icon, Category: "Frameworks"},
				{Name: "Gin", Icon: p.technologyService.GetTechnology("Gin").Icon, Category: "Frameworks"},
				{Name: "Echo", Icon: p.technologyService.GetTechnology("Echo").Icon, Category: "Frameworks"},
				{Name: "FastAPI", Icon: p.technologyService.GetTechnology("FastAPI").Icon, Category: "Frameworks"},
			},
		},
		{
			Category: "Infrastructure",
			Skills: []models.Skill{
				{Name: "Kubernetes", Icon: p.technologyService.GetTechnology("Kubernetes").Icon, Category: "Infrastructure"},
				{Name: "Docker", Icon: p.technologyService.GetTechnology("Docker").Icon, Category: "Infrastructure"},
				{Name: "AWS", Icon: p.technologyService.GetTechnology("AWS").Icon, Category: "Infrastructure"},
				{Name: "Terraform", Icon: p.technologyService.GetTechnology("Terraform").Icon, Category: "Infrastructure"},
			},
		},
		{
			Category: "Architecture",
			Skills: []models.Skill{
				{Name: "Microservices", Icon: p.technologyService.GetTechnology("Microservices").Icon, Category: "Architecture"},
				{Name: "Event Sourcing", Icon: p.technologyService.GetTechnology("Event Sourcing").Icon, Category: "Architecture"},
				{Name: "CQRS", Icon: p.technologyService.GetTechnology("CQRS").Icon, Category: "Architecture"},
				{Name: "DDD", Icon: p.technologyService.GetTechnology("DDD").Icon, Category: "Architecture"},
			},
		},
	}
}
