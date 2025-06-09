package repository

import (
	"fmt"

	"github.com/benidevo/website/internal/models"
)

// InMemoryProjectRepository implements ProjectRepository with in-memory data
type InMemoryProjectRepository struct {
	projects map[int]*models.Project
}

// NewInMemoryProjectRepository creates a new in-memory project repository
func NewInMemoryProjectRepository(techRepo TechnologyRepository) *InMemoryProjectRepository {
	repo := &InMemoryProjectRepository{
		projects: make(map[int]*models.Project),
	}
	repo.initializeData(techRepo)
	return repo
}

// GetAllProjects returns all projects
func (r *InMemoryProjectRepository) GetAllProjects() ([]*models.Project, error) {
	var projects []*models.Project
	for _, project := range r.projects {
		projects = append(projects, project)
	}
	return projects, nil
}

// GetProjectByID returns a project by its ID
func (r *InMemoryProjectRepository) GetProjectByID(id int) (*models.Project, error) {
	project, exists := r.projects[id]
	if !exists {
		return nil, fmt.Errorf("project with ID %d not found", id)
	}
	return project, nil
}

// initializeData sets up the in-memory data
func (r *InMemoryProjectRepository) initializeData(techRepo TechnologyRepository) {
	projectsData := r.getProjectsData(techRepo)
	for _, project := range projectsData {
		r.projects[project.ID] = project
	}
}

// getProjectsData returns the projects data with pre-resolved technologies and all details
func (r *InMemoryProjectRepository) getProjectsData(techRepo TechnologyRepository) map[string]*models.Project {
	return map[string]*models.Project{
		"ascentio": {
			ID:                     1,
			Title:                  "Ascentio - Job Prospecting Platform",
			Description:            "A comprehensive job prospecting platform built with Go, featuring microservices architecture, event-driven design, and real-time notifications.",
			DetailedDescription:    "Ascentio is a comprehensive job prospecting platform designed to revolutionize how job seekers discover and apply for opportunities.\n\nBuilt with a microservices architecture using Go and modern cloud technologies, the platform automates job discovery, matches candidates with suitable positions, and streamlines the application process while maintaining personalization.\n\nThe system employs event-driven communication for loose coupling, with distinct services handling:\n• User management and authentication\n• Job aggregation from multiple sources\n• Intelligent matching algorithms\n• Real-time notification delivery\n\nKey technical achievements include:\n• Implementing a weighted scoring algorithm using Go's concurrent processing capabilities for efficient job matching\n• Building an event-driven notification system using message queues for reliable delivery\n• Creating a robust data pipeline with error handling and retry mechanisms for external API integration\n\nUsing PostgreSQL with read replicas and Redis caching, the platform processes over 10,000 job listings daily with sub-second matching response times, achieving 95% notification delivery rates and reducing manual job search time by 70% for users.",
			ArchitectureDiagramURL: "/static/images/diagrams/ascentio-architecture.svg",
			GitHubURL:              "https://github.com/benidevo/ascentio",
			LiveURL:                "https://ascentio.demo.benidevo.com",
			Language:               "Go",
			Technologies:           techRepo.GetTechnologies([]string{"Go", "PostgreSQL", "Redis", "Docker"}),
			Featured:               true,
		},
		"distributed-url-shortener": {
			ID:                     2,
			Title:                  "Distributed URL Shortener",
			Description:            "A distributed URL shortening system with geographic analytics capabilities using FastAPI and deployed on Kubernetes.",
			DetailedDescription:    "A highly scalable distributed URL shortening system built with geographic analytics capabilities using Python FastAPI and deployed on Kubernetes.\n\nThe system addresses the limitations of traditional URL shorteners by providing:\n• Comprehensive analytics and metrics\n• Geographic tracking with privacy protection\n• Horizontal scaling to handle millions of requests per day\n\nThe microservices architecture features:\n• FastAPI services handling URL operations\n• Geographic data processing pipeline using external geolocation APIs with aggressive caching strategies\n• PostgreSQL with read replicas for data persistence alongside Redis for high-speed lookups\n\nTechnical innovations include:\n• Developing a base62 encoding system with collision detection and retry mechanisms\n• Implementing privacy-respecting geographic analytics using IP geolocation with data anonymization\n• Creating a multi-tier caching strategy with TTL-based invalidation for optimal performance\n\nThe system achieves 99.99% uptime while processing 1M+ URL redirections daily, with geographic analytics providing valuable insights at 95% accuracy while maintaining user privacy compliance.",
			ArchitectureDiagramURL: "/static/images/diagrams/url-shortener-architecture.svg",
			GitHubURL:              "https://github.com/benidevo/url-shortener",
			LiveURL:                "https://s.benidevo.com",
			Language:               "Go",
			Technologies:           techRepo.GetTechnologies([]string{"Python", "FastAPI", "Docker", "Kubernetes", "PostgreSQL", "Redis"}),
			Featured:               true,
		},
		"order-fulfillment": {
			ID:                     3,
			Title:                  "Order Fulfillment System",
			Description:            "A robust order fulfillment system using event sourcing, CQRS, and Saga patterns for handling complex business workflows.",
			DetailedDescription:    "A robust order fulfillment system implementing advanced distributed system patterns including event sourcing, CQRS (Command Query Responsibility Segregation), and Saga patterns for handling complex business workflows reliably. The system addresses the challenges of traditional order processing systems that struggle with complex business workflows, data consistency across services, and handling failure scenarios gracefully. Built with an event-driven architecture, command services handle write operations while query services optimize read performance, with Apache Kafka ensuring reliable event delivery and support for event replay and audit trails. Key architectural decisions include building domain aggregates with proper event sourcing patterns using Java and Spring Boot, implementing eventual consistency with compensation patterns for distributed transactions, and creating choreography-based sagas for loose coupling between services. The system successfully reduced order processing time by 60% while improving overall reliability, with event sourcing enabling complete audit trails and simplified debugging of complex business workflows.",
			ArchitectureDiagramURL: "/static/images/diagrams/order-fulfillment-architecture.svg",
			GitHubURL:              "https://github.com/benidevo/order-fulfillment",
			Language:               "Java",
			Technologies:           techRepo.GetTechnologies([]string{"Java", "Spring Boot", "Go", "Apache Kafka", "PostgreSQL", "Docker", "Kubernetes", "Nginx"}),
			Featured:               true,
		},
		"minicon": {
			ID:                     4,
			Title:                  "MiniCon",
			Description:            "A lightweight container implementation in Python that demonstrates core virtualization concepts like process isolation, resource management, and networking.",
			DetailedDescription:    "MiniCon is a lightweight container implementation built in Python that demonstrates core virtualization concepts including process isolation, resource management, and networking. This educational project was designed to provide deep understanding of Linux kernel features like namespaces, cgroups, and chroot that power modern containerization technologies. The implementation leverages Linux kernel primitives directly, using process isolation through namespaces (PID, network, mount), resource management via cgroups, and filesystem isolation through chroot and overlay filesystems. Technical highlights include using Linux namespaces API through Python's ctypes for process isolation, implementing cgroups v2 integration for fine-grained resource control, creating overlay filesystem management for efficient image layering, and building robust cleanup mechanisms to prevent resource leaks. The project successfully demonstrates core container concepts with 90% feature parity to basic Docker functionality, serving as both a practical containerization tool and an educational resource for understanding containerization internals with minimal dependencies to showcase fundamental concepts clearly.",
			ArchitectureDiagramURL: "/static/images/diagrams/minicon-architecture.svg",
			GitHubURL:              "https://github.com/benidevo/minicon",
			Language:               "Java",
			Technologies:           techRepo.GetTechnologies([]string{"Python", "Linux", "Virtualization"}),
			Featured:               true,
		},
	}
}

// InMemorySkillRepository implements SkillRepository with in-memory data
type InMemorySkillRepository struct {
	skillCategories []models.SkillCategory // Pre-built categories for fast access
}

// NewInMemorySkillRepository creates a new in-memory skill repository
func NewInMemorySkillRepository(techRepo TechnologyRepository) *InMemorySkillRepository {
	repo := &InMemorySkillRepository{}
	repo.initializeSkillCategories(techRepo)
	return repo
}

// GetSkillCategories returns all skill categories (pre-built for performance)
func (r *InMemorySkillRepository) GetSkillCategories() ([]models.SkillCategory, error) {
	// Return a copy to prevent external modification
	categories := make([]models.SkillCategory, len(r.skillCategories))
	copy(categories, r.skillCategories)
	return categories, nil
}

// initializeSkillCategories pre-builds skill categories with resolved technology icons
func (r *InMemorySkillRepository) initializeSkillCategories(techRepo TechnologyRepository) {
	r.skillCategories = []models.SkillCategory{
		{
			Category: "Languages",
			Skills: []models.Skill{
				{Name: "Go", Icon: r.getTechIcon(techRepo, "Go"), Category: "Languages"},
				{Name: "Python", Icon: r.getTechIcon(techRepo, "Python"), Category: "Languages"},
				{Name: "Java", Icon: r.getTechIcon(techRepo, "Java"), Category: "Languages"},
				{Name: "JavaScript", Icon: r.getTechIcon(techRepo, "JavaScript"), Category: "Languages"},
			},
		},
		{
			Category: "Frameworks",
			Skills: []models.Skill{
				{Name: "Gin", Icon: r.getTechIcon(techRepo, "Gin"), Category: "Frameworks"},
				{Name: "Django", Icon: r.getTechIcon(techRepo, "Django"), Category: "Frameworks"},
				{Name: "FastAPI", Icon: r.getTechIcon(techRepo, "FastAPI"), Category: "Frameworks"},
				{Name: "Flask", Icon: r.getTechIcon(techRepo, "Flask"), Category: "Frameworks"},
				{Name: "Spring Boot", Icon: r.getTechIcon(techRepo, "Spring Boot"), Category: "Frameworks"},
			},
		},
		{
			Category: "Infrastructure & DevOps",
			Skills: []models.Skill{
				{Name: "GCP", Icon: r.getTechIcon(techRepo, "Google Cloud Platform"), Category: "Infrastructure & DevOps"},
				{Name: "Kubernetes", Icon: r.getTechIcon(techRepo, "Kubernetes"), Category: "Infrastructure & DevOps"},
				{Name: "Docker", Icon: r.getTechIcon(techRepo, "Docker"), Category: "Infrastructure & DevOps"},
				{Name: "Terraform", Icon: r.getTechIcon(techRepo, "Terraform"), Category: "Infrastructure & DevOps"},
				{Name: "GitHub Actions", Icon: r.getTechIcon(techRepo, "GitHub Actions"), Category: "Infrastructure & DevOps"},
				{Name: "Helm", Icon: r.getTechIcon(techRepo, "Helm"), Category: "Infrastructure & DevOps"},
			},
		},
		{
			Category: "Databases & Messaging",
			Skills: []models.Skill{
				{Name: "Redis", Icon: r.getTechIcon(techRepo, "Redis"), Category: "Databases & Messaging"},
				{Name: "MongoDB", Icon: r.getTechIcon(techRepo, "MongoDB"), Category: "Databases & Messaging"},
				{Name: "MySQL", Icon: r.getTechIcon(techRepo, "MySQL"), Category: "Databases & Messaging"},
				{Name: "PostgreSQL", Icon: r.getTechIcon(techRepo, "PostgreSQL"), Category: "Databases & Messaging"},
				{Name: "Apache Kafka", Icon: r.getTechIcon(techRepo, "Apache Kafka"), Category: "Databases & Messaging"},
				{Name: "gRPC", Icon: r.getTechIcon(techRepo, "gRPC"), Category: "Databases & Messaging"},
			},
		},
		{
			Category: "Architecture & Patterns",
			Skills: []models.Skill{
				{Name: "Distributed Systems", Icon: r.getTechIcon(techRepo, "Distributed Systems"), Category: "Architecture & Patterns"},
				{Name: "Microservices", Icon: r.getTechIcon(techRepo, "Microservices"), Category: "Architecture & Patterns"},
				{Name: "Event Sourcing", Icon: r.getTechIcon(techRepo, "Event Sourcing"), Category: "Architecture & Patterns"},
				{Name: "CQRS", Icon: r.getTechIcon(techRepo, "CQRS"), Category: "Architecture & Patterns"},
				{Name: "System Integration", Icon: r.getTechIcon(techRepo, "System Integration"), Category: "Architecture & Patterns"},
				{Name: "Domain-Driven Design", Icon: r.getTechIcon(techRepo, "Domain-Driven Design"), Category: "Architecture & Patterns"},
			},
		},
		{
			Category: "Monitoring",
			Skills: []models.Skill{
				{Name: "Grafana", Icon: r.getTechIcon(techRepo, "Grafana"), Category: "Monitoring"},
				{Name: "Prometheus", Icon: r.getTechIcon(techRepo, "Prometheus"), Category: "Monitoring"},
				{Name: "OpenTelemetry", Icon: r.getTechIcon(techRepo, "OpenTelemetry"), Category: "Monitoring"},
			},
		},
	}
}

func (r *InMemorySkillRepository) getTechIcon(techRepo TechnologyRepository, name string) string {
	tech, err := techRepo.GetTechnology(name)
	if err != nil {
		return ""
	}
	return tech.Icon
}
