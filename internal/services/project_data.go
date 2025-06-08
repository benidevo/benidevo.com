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
