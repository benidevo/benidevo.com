package services

import (
	"github.com/benidevo/website/internal/models"
)

// getFeaturedProjectsData returns a map of featured project identifiers to their corresponding
// ProjectDetails.
//
// The technologies are fetched using the technologyService.
func (p *ProjectService) getFeaturedProjectsData() map[string]*models.ProjectDetails {
	return map[string]*models.ProjectDetails{
		"ascentio": {
			Project: models.Project{
				ID:           "ascentio",
				Title:        "Ascentio - Job Prospecting Platform",
				Description:  "A comprehensive job prospecting platform built with Go, featuring microservices architecture, event-driven design, and real-time notifications.",
				GitHubURL:    "https://github.com/benidevo/ascentio",
				Language:     "Go",
				Technologies: p.technologyService.GetTechnologies([]string{"Go", "PostgreSQL", "Redis", "Docker"}),
				Featured:     true,
			},
		},
		"distributed-redis": {
			Project: models.Project{
				ID:           "distributed-redis",
				Title:        "Distributed Redis Implementation",
				Description:  "A distributed key-value store implementing Redis protocol with consistent hashing, replication, and automatic failover capabilities.",
				GitHubURL:    "https://github.com/benidevo/distributed-redis",
				Language:     "Go",
				Technologies: p.technologyService.GetTechnologies([]string{"Go", "gRPC", "Docker"}),
				Featured:     true,
			},
		},
		"event-driven-order": {
			Project: models.Project{
				ID:           "event-driven-order",
				Title:        "Event-Driven Order Fulfillment",
				Description:  "A robust order fulfillment system using event sourcing, CQRS, and Saga patterns for handling complex business workflows.",
				GitHubURL:    "https://github.com/benidevo/event-driven-order",
				Language:     "Java",
				Technologies: p.technologyService.GetTechnologies([]string{"Java", "Spring Boot", "Apache Kafka", "PostgreSQL", "Docker", "Kubernetes"}),
				Featured:     true,
			},
		},
		"microservices-weather": {
			Project: models.Project{
				ID:           "microservices-weather",
				Title:        "Microservices Weather System",
				Description:  "A cloud-native weather data aggregation system with Spring Boot microservices, featuring circuit breakers, distributed tracing, and auto-scaling.",
				GitHubURL:    "https://github.com/benidevo/microservices-weather",
				Language:     "Java",
				Technologies: p.technologyService.GetTechnologies([]string{"Java", "Spring Boot", "Spring Cloud", "Docker", "Kubernetes", "Prometheus"}),
				Featured:     true,
			},
		},
	}
}
