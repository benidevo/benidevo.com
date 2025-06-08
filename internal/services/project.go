package services

import (
	"github.com/benidevo/website/internal/models"
)

// ProjectService provides methods to manage projects, integrating with GitHub and technology services.
// It maintains a map of project details for quick access.
type ProjectService struct {
	githubService     *GitHubService
	technologyService *TechnologyService
}

// NewProjectService creates a new project service
func NewProjectService(githubService *GitHubService) *ProjectService {
	return &ProjectService{
		githubService:     githubService,
		technologyService: NewTechnologyService(),
	}
}

// GetFeaturedProjects returns featured projects
func (p *ProjectService) GetFeaturedProjects() []*models.Project {
	projectData := p.getFeaturedProjectsData()

	var featuredProjects []*models.Project
	for _, data := range projectData {
		featuredProjects = append(featuredProjects, data)
	}

	return featuredProjects
}

// GetProjectDetail returns detailed information for a specific project
func (p *ProjectService) GetProjectDetail(projectID int) *models.ProjectDetail {
	projectData := p.getFeaturedProjectsData()

	var project *models.Project
	for _, proj := range projectData {
		if proj.ID == projectID {
			project = proj
			break
		}
	}

	if project == nil {
		return nil
	}

	detailedProjects := p.getDetailedProjectsData()

	for _, detail := range detailedProjects {
		if detail.Project.ID == projectID {
			return detail
		}
	}

	return nil
}

// GetSkillCategories returns skill categories for the home page
func (p *ProjectService) GetSkillCategories() []models.SkillCategory {
	return []models.SkillCategory{
		{
			Category: "Languages",
			Skills: []models.Skill{
				{Name: "Go", Icon: p.technologyService.GetTechnology("Go").Icon, Category: "Languages"},
				{Name: "Python", Icon: p.technologyService.GetTechnology("Python").Icon, Category: "Languages"},
				{Name: "Java", Icon: p.technologyService.GetTechnology("Java").Icon, Category: "Languages"},
				{Name: "JavaScript", Icon: p.technologyService.GetTechnology("JavaScript").Icon, Category: "Languages"},
			},
		},
		{
			Category: "Frameworks",
			Skills: []models.Skill{
				{Name: "Gin", Icon: p.technologyService.GetTechnology("Gin").Icon, Category: "Frameworks"},
				{Name: "FastAPI", Icon: p.technologyService.GetTechnology("FastAPI").Icon, Category: "Frameworks"},
				{Name: "Django", Icon: p.technologyService.GetTechnology("Django").Icon, Category: "Frameworks"},
				{Name: "Flask", Icon: p.technologyService.GetTechnology("Flask").Icon, Category: "Frameworks"},
				{Name: "Spring Boot", Icon: p.technologyService.GetTechnology("Spring Boot").Icon, Category: "Frameworks"},
			},
		},
		{
			Category: "Infrastructure",
			Skills: []models.Skill{
				{Name: "Kubernetes", Icon: p.technologyService.GetTechnology("Kubernetes").Icon, Category: "Infrastructure"},
				{Name: "Docker", Icon: p.technologyService.GetTechnology("Docker").Icon, Category: "Infrastructure"},
				{Name: "Google Cloud", Icon: p.technologyService.GetTechnology("Google Cloud").Icon, Category: "Infrastructure"},
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
