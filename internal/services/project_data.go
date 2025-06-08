package services

import (
	"github.com/benidevo/website/internal/models"
)

// getFeaturedProjectsData returns a map of featured project identifiers to their corresponding
// ProjectDetails.
//
// The technologies are fetched using the technologyService.
func (p *ProjectService) getFeaturedProjectsData() map[string]*models.Project {
	return map[string]*models.Project{
		"ascentio": {
			ID:           1,
			Title:        "Ascentio - Job Prospecting Platform",
			Description:  "A comprehensive job prospecting platform built with Go, featuring microservices architecture, event-driven design, and real-time notifications.",
			GitHubURL:    "https://github.com/benidevo/ascentio",
			Language:     "Go",
			Technologies: p.technologyService.GetTechnologies([]string{"Go", "PostgreSQL", "Redis", "Docker"}),
			Featured:     true,
		},
		"distributed-url-shortener": {
			ID:           2,
			Title:        "Distributed URL Shortener",
			Description:  "A distributed URL shortening system with geographic analytics capabilities using FastAPI and deployed on Kubernetes.",
			GitHubURL:    "https://github.com/benidevo/url-shortener",
			Language:     "Go",
			Technologies: p.technologyService.GetTechnologies([]string{"Python", "FastAPI", "Docker", "Kubernetes", "PostgreSQL", "Redis"}),
			Featured:     true,
		},
		"order-fulfillment": {
			ID:           3,
			Title:        "Order Fulfillment System",
			Description:  "A robust order fulfillment system using event sourcing, CQRS, and Saga patterns for handling complex business workflows.",
			GitHubURL:    "https://github.com/benidevo/order-fulfillment",
			Language:     "Java",
			Technologies: p.technologyService.GetTechnologies([]string{"Java", "Spring Boot", "Go", "Apache Kafka", "PostgreSQL", "Docker", "Kubernetes", "Nginx"}),
			Featured:     true,
		},
		"minicon": {
			ID:           4,
			Title:        "MiniCon",
			Description:  "A lightweight container implementation in Python that demonstrates core virtualization concepts like process isolation, resource management, and networking.",
			GitHubURL:    "https://github.com/benidevo/minicon",
			Language:     "Java",
			Technologies: p.technologyService.GetTechnologies([]string{"Python", "Linux", "Virtualization"}),
			Featured:     true,
		},
	}
}

// getDetailedProjectsData returns detailed project information for the project detail pages
func (p *ProjectService) getDetailedProjectsData() []*models.ProjectDetail {
	projectsMap := p.getFeaturedProjectsData()

	return []*models.ProjectDetail{
		{
			Project:                projectsMap["ascentio"],
			DetailedDescription:    "Ascentio is a comprehensive job prospecting platform designed to revolutionize how job seekers discover and apply for opportunities. Built with a microservices architecture using Go and modern cloud technologies, the platform automates job discovery, matches candidates with suitable positions, and streamlines the application process while maintaining personalization. The system employs event-driven communication for loose coupling, with distinct services handling user management, job aggregation, intelligent matching algorithms, and real-time notification delivery. Key technical achievements include implementing a weighted scoring algorithm using Go's concurrent processing capabilities for efficient job matching, building an event-driven notification system using message queues for reliable delivery, and creating a robust data pipeline with error handling and retry mechanisms for external API integration. Using PostgreSQL with read replicas and Redis caching, the platform processes over 10,000 job listings daily with sub-second matching response times, achieving 95% notification delivery rates and reducing manual job search time by 70% for users.",
			ArchitectureDiagramURL: "/static/images/diagrams/ascentio-architecture.svg",
		},
		{
			Project:                projectsMap["distributed-url-shortener"],
			DetailedDescription:    "A highly scalable distributed URL shortening system built with geographic analytics capabilities using Python FastAPI and deployed on Kubernetes. The system addresses the limitations of traditional URL shorteners by providing comprehensive analytics, geographic tracking, and horizontal scaling to handle millions of requests per day. The microservices architecture features FastAPI services handling URL operations, a geographic data processing pipeline using external geolocation APIs with aggressive caching strategies, and PostgreSQL with read replicas for data persistence alongside Redis for high-speed lookups. Technical innovations include developing a base62 encoding system with collision detection and retry mechanisms, implementing privacy-respecting geographic analytics using IP geolocation with data anonymization, and creating a multi-tier caching strategy with TTL-based invalidation for optimal performance. The system achieves 99.99% uptime while processing 1M+ URL redirections daily, with geographic analytics providing valuable insights at 95% accuracy while maintaining user privacy compliance.",
			ArchitectureDiagramURL: "/static/images/diagrams/url-shortener-architecture.svg",
		},
		{
			Project:                projectsMap["order-fulfillment"],
			DetailedDescription:    "A robust order fulfillment system implementing advanced distributed system patterns including event sourcing, CQRS (Command Query Responsibility Segregation), and Saga patterns for handling complex business workflows reliably. The system addresses the challenges of traditional order processing systems that struggle with complex business workflows, data consistency across services, and handling failure scenarios gracefully. Built with an event-driven architecture, command services handle write operations while query services optimize read performance, with Apache Kafka ensuring reliable event delivery and support for event replay and audit trails. Key architectural decisions include building domain aggregates with proper event sourcing patterns using Java and Spring Boot, implementing eventual consistency with compensation patterns for distributed transactions, and creating choreography-based sagas for loose coupling between services. The system successfully reduced order processing time by 60% while improving overall reliability, with event sourcing enabling complete audit trails and simplified debugging of complex business workflows.",
			ArchitectureDiagramURL: "/static/images/diagrams/order-fulfillment-architecture.svg",
		},
		{
			Project:                projectsMap["minicon"],
			DetailedDescription:    "MiniCon is a lightweight container implementation built in Python that demonstrates core virtualization concepts including process isolation, resource management, and networking. This educational project was designed to provide deep understanding of Linux kernel features like namespaces, cgroups, and chroot that power modern containerization technologies. The implementation leverages Linux kernel primitives directly, using process isolation through namespaces (PID, network, mount), resource management via cgroups, and filesystem isolation through chroot and overlay filesystems. Technical highlights include using Linux namespaces API through Python's ctypes for process isolation, implementing cgroups v2 integration for fine-grained resource control, creating overlay filesystem management for efficient image layering, and building robust cleanup mechanisms to prevent resource leaks. The project successfully demonstrates core container concepts with 90% feature parity to basic Docker functionality, serving as both a practical containerization tool and an educational resource for understanding containerization internals with minimal dependencies to showcase fundamental concepts clearly.",
			ArchitectureDiagramURL: "/static/images/diagrams/minicon-architecture.svg",
		},
	}
}
