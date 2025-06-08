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
			Project:          projectsMap["ascentio"],
			ProblemStatement: "Job seekers often struggle with inefficient prospecting processes, spending countless hours manually searching for relevant opportunities and crafting personalized applications. The challenge was to create a platform that automates job discovery, matches candidates with suitable positions, and streamlines the application process while maintaining personalization.",
			Architecture:     "The system employs a microservices architecture with distinct services for user management, job aggregation, matching algorithms, and notification delivery. Event-driven communication ensures loose coupling, while Redis caching optimizes performance for real-time job matching.",
			TechnicalChallenges: []string{
				"Designing a scalable job matching algorithm that considers multiple criteria (skills, experience, location, salary expectations)",
				"Implementing real-time notifications across multiple channels (email, SMS, in-app)",
				"Handling high-volume job data ingestion from multiple external APIs",
				"Ensuring data consistency across distributed services while maintaining performance",
			},
			Solutions: []string{
				"Implemented a weighted scoring algorithm using Go's concurrent processing capabilities for efficient job matching",
				"Built an event-driven notification system using message queues for reliable delivery",
				"Created a robust data pipeline with error handling and retry mechanisms for external API integration",
				"Used PostgreSQL with read replicas and Redis for caching to balance consistency and performance",
			},
			Results:                "Successfully processed over 10,000 job listings daily with sub-second matching response times. Achieved 95% notification delivery rate and reduced manual job search time by 70% for users.",
			ArchitectureDiagramURL: "/static/images/diagrams/ascentio-architecture.svg",
			TechStackDetails:       "Backend built with Go and Gin framework for high-performance API development. PostgreSQL for persistent data storage with Redis for caching and session management. Docker containerization enables consistent deployment across environments.",
			KeyFeatures: []string{
				"Intelligent job matching with machine learning-based scoring",
				"Real-time notification system across multiple channels",
				"Automated cover letter generation based on job requirements",
				"Comprehensive analytics dashboard for tracking application success",
				"RESTful API with comprehensive documentation",
			},
			Metrics: []models.Metric{
				{Name: "Jobs Processed Daily", Value: "10,000+", Unit: "jobs"},
				{Name: "Matching Response Time", Value: "<500", Unit: "ms"},
				{Name: "API Uptime", Value: "99.9", Unit: "%"},
				{Name: "User Search Time Reduction", Value: "70", Unit: "%"},
			},
		},
		{
			Project:          projectsMap["distributed-url-shortener"],
			ProblemStatement: "Traditional URL shorteners lack geographic analytics and struggle with high-traffic scenarios. The goal was to build a distributed system that provides comprehensive analytics, geographic tracking, and scales horizontally to handle millions of requests per day.",
			Architecture:     "Microservices architecture deployed on Kubernetes with FastAPI services handling URL operations. Geographic data processing pipeline using external geolocation APIs with aggressive caching strategies. PostgreSQL with read replicas for data persistence and Redis for high-speed lookups.",
			TechnicalChallenges: []string{
				"Designing a URL encoding algorithm that generates short, collision-free identifiers",
				"Implementing geographic analytics without compromising user privacy",
				"Handling traffic spikes and ensuring consistent performance under load",
				"Creating an efficient caching strategy that balances memory usage and hit rates",
			},
			Solutions: []string{
				"Developed a base62 encoding system with collision detection and retry mechanisms",
				"Implemented privacy-respecting geographic analytics using IP geolocation with data anonymization",
				"Used Kubernetes horizontal pod autoscaling with Redis Cluster for distributed caching",
				"Created a multi-tier caching strategy with TTL-based invalidation for optimal performance",
			},
			Results:                "Achieved 99.99% uptime while processing 1M+ URL redirections daily. Geographic analytics provided valuable insights with 95% accuracy while maintaining user privacy compliance.",
			ArchitectureDiagramURL: "/static/images/diagrams/url-shortener-architecture.svg",
			TechStackDetails:       "Python FastAPI for high-performance async endpoints. Kubernetes orchestration with horizontal scaling capabilities. PostgreSQL for reliable data persistence with Redis Cluster for distributed caching. Nginx for load balancing and SSL termination.",
			KeyFeatures: []string{
				"High-performance URL shortening with collision-free algorithms",
				"Comprehensive geographic analytics dashboard",
				"Real-time click tracking and statistics",
				"RESTful API with rate limiting and authentication",
				"Custom domain support with SSL certificate management",
			},
			Metrics: []models.Metric{
				{Name: "Daily Redirections", Value: "1M+", Unit: "requests"},
				{Name: "Average Response Time", Value: "<100", Unit: "ms"},
				{Name: "Geographic Accuracy", Value: "95", Unit: "%"},
				{Name: "System Uptime", Value: "99.99", Unit: "%"},
			},
		},
		{
			Project:          projectsMap["order-fulfillment"],
			ProblemStatement: "Traditional order processing systems struggle with complex business workflows, data consistency across services, and handling failure scenarios gracefully. The challenge was to design a resilient system using event sourcing and CQRS patterns that could handle distributed transactions reliably.",
			Architecture:     "Event-driven architecture implementing CQRS and Event Sourcing patterns. Command services handle write operations while query services optimize read performance. Apache Kafka ensures reliable event delivery with support for event replay and audit trails.",
			TechnicalChallenges: []string{
				"Implementing event sourcing with proper aggregate design and conflict resolution",
				"Ensuring data consistency across distributed services without traditional transactions",
				"Designing saga patterns for complex multi-step business processes",
				"Creating efficient event replay mechanisms for system recovery and debugging",
			},
			Solutions: []string{
				"Built domain aggregates with proper event sourcing patterns using Java and Spring Boot",
				"Implemented eventual consistency with compensation patterns for distributed transactions",
				"Created choreography-based sagas for loose coupling between services",
				"Developed event store with snapshotting for efficient aggregate reconstruction",
			},
			Results:                "Reduced order processing time by 60% while improving system reliability. Event sourcing enabled complete audit trails and simplified debugging of complex business workflows.",
			ArchitectureDiagramURL: "/static/images/diagrams/order-fulfillment-architecture.svg",
			TechStackDetails:       "Java Spring Boot microservices with Apache Kafka for event streaming. PostgreSQL for event store with optimized event replay capabilities. Docker and Kubernetes for containerized deployment with service mesh for inter-service communication.",
			KeyFeatures: []string{
				"Event-driven order processing with automatic retry mechanisms",
				"Complete audit trail with event sourcing implementation",
				"Saga pattern implementation for distributed transactions",
				"Real-time order status tracking and notifications",
				"Comprehensive monitoring and alerting system",
			},
			Metrics: []models.Metric{
				{Name: "Processing Time Reduction", Value: "60", Unit: "%"},
				{Name: "Event Processing Rate", Value: "10,000", Unit: "events/sec"},
				{Name: "System Recovery Time", Value: "<30", Unit: "seconds"},
				{Name: "Data Consistency Rate", Value: "99.99", Unit: "%"},
			},
		},
		{
			Project:          projectsMap["minicon"],
			ProblemStatement: "Understanding container virtualization requires deep knowledge of Linux kernel features like namespaces, cgroups, and chroot. The goal was to build a minimal container implementation that demonstrates these core concepts while providing practical containerization capabilities.",
			Architecture:     "Lightweight container runtime built in Python leveraging Linux kernel primitives. Process isolation using namespaces (PID, network, mount), resource management through cgroups, and filesystem isolation via chroot and overlay filesystems.",
			TechnicalChallenges: []string{
				"Implementing proper process isolation without affecting host system stability",
				"Managing resource constraints (CPU, memory, disk) using cgroups effectively",
				"Creating secure filesystem isolation while maintaining performance",
				"Handling container lifecycle management (start, stop, cleanup) reliably",
			},
			Solutions: []string{
				"Used Linux namespaces API through Python's ctypes for process isolation",
				"Implemented cgroups v2 integration for fine-grained resource control",
				"Created overlay filesystem management for efficient image layering",
				"Built robust cleanup mechanisms to prevent resource leaks",
			},
			Results:                "Successfully demonstrated core container concepts with 90% feature parity to basic Docker functionality. Provided educational tool for understanding containerization internals.",
			ArchitectureDiagramURL: "/static/images/diagrams/minicon-architecture.svg",
			TechStackDetails:       "Pure Python implementation utilizing Linux system calls and kernel features. Direct integration with cgroups, namespaces, and filesystem APIs. Minimal dependencies to showcase fundamental containerization concepts clearly.",
			KeyFeatures: []string{
				"Process isolation using Linux namespaces",
				"Resource management through cgroups integration",
				"Filesystem isolation with overlay support",
				"Container image management and layering",
				"Command-line interface similar to Docker",
			},
			Metrics: []models.Metric{
				{Name: "Container Start Time", Value: "<2", Unit: "seconds"},
				{Name: "Memory Overhead", Value: "<10", Unit: "MB"},
				{Name: "Feature Parity", Value: "90", Unit: "%"},
				{Name: "Resource Isolation", Value: "99.9", Unit: "%"},
			},
		},
	}
}
