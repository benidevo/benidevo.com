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
				{Name: "Django", Icon: p.technologyService.GetTechnology("Django").Icon, Category: "Frameworks"},
				{Name: "FastAPI", Icon: p.technologyService.GetTechnology("FastAPI").Icon, Category: "Frameworks"},
				{Name: "Flask", Icon: p.technologyService.GetTechnology("Flask").Icon, Category: "Frameworks"},
				{Name: "Spring Boot", Icon: p.technologyService.GetTechnology("Spring Boot").Icon, Category: "Frameworks"},
			},
		},
		{
			Category: "Infrastructure & DevOps",
			Skills: []models.Skill{
				{Name: "GCP", Icon: p.technologyService.GetTechnology("Google Cloud Platform").Icon, Category: "Infrastructure & DevOps"},
				{Name: "Kubernetes", Icon: p.technologyService.GetTechnology("Kubernetes").Icon, Category: "Infrastructure & DevOps"},
				{Name: "Docker", Icon: p.technologyService.GetTechnology("Docker").Icon, Category: "Infrastructure & DevOps"},
				{Name: "Terraform", Icon: p.technologyService.GetTechnology("Terraform").Icon, Category: "Infrastructure & DevOps"},
				{Name: "GitHub Actions", Icon: p.technologyService.GetTechnology("GitHub Actions").Icon, Category: "Infrastructure & DevOps"},
				{Name: "Helm", Icon: p.technologyService.GetTechnology("Helm").Icon, Category: "Infrastructure & DevOps"},
			},
		},
		{
			Category: "Databases & Messaging",
			Skills: []models.Skill{
				{Name: "Redis", Icon: p.technologyService.GetTechnology("Redis").Icon, Category: "Databases & Messaging"},
				{Name: "MongoDB", Icon: p.technologyService.GetTechnology("MongoDB").Icon, Category: "Databases & Messaging"},
				{Name: "MySQL", Icon: p.technologyService.GetTechnology("MySQL").Icon, Category: "Databases & Messaging"},
				{Name: "PostgreSQL", Icon: p.technologyService.GetTechnology("PostgreSQL").Icon, Category: "Databases & Messaging"},
				{Name: "Apache Kafka", Icon: p.technologyService.GetTechnology("Apache Kafka").Icon, Category: "Databases & Messaging"},
				{Name: "gRPC", Icon: p.technologyService.GetTechnology("gRPC").Icon, Category: "Databases & Messaging"},
			},
		},
		{
			Category: "Architecture & Patterns",
			Skills: []models.Skill{
				{Name: "Distributed Systems", Icon: p.technologyService.GetTechnology("Distributed Systems").Icon, Category: "Architecture & Patterns"},
				{Name: "Microservices", Icon: p.technologyService.GetTechnology("Microservices").Icon, Category: "Architecture & Patterns"},
				{Name: "Event Sourcing", Icon: p.technologyService.GetTechnology("Event Sourcing").Icon, Category: "Architecture & Patterns"},
				{Name: "CQRS", Icon: p.technologyService.GetTechnology("CQRS").Icon, Category: "Architecture & Patterns"},
				{Name: "System Integration", Icon: p.technologyService.GetTechnology("System Integration").Icon, Category: "Architecture & Patterns"},
				{Name: "Domain-Driven Design", Icon: p.technologyService.GetTechnology("Domain-Driven Design").Icon, Category: "Architecture & Patterns"},
			},
		},
		{
			Category: "Monitoring",
			Skills: []models.Skill{
				{Name: "Grafana", Icon: p.technologyService.GetTechnology("Grafana").Icon, Category: "Monitoring"},
				{Name: "Prometheus", Icon: p.technologyService.GetTechnology("Prometheus").Icon, Category: "Monitoring"},
				{Name: "OpenTelemetry", Icon: p.technologyService.GetTechnology("OpenTelemetry").Icon, Category: "Monitoring"},
			},
		},
	}
}
