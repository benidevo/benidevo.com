package repository

import (
	"fmt"

	"github.com/benidevo/website/internal/models"
)

// InMemoryTechnologyRepository implements TechnologyRepository with in-memory data
type InMemoryTechnologyRepository struct {
	technologies map[string]models.Technology
}

// NewInMemoryTechnologyRepository creates a new in-memory technology repository
func NewInMemoryTechnologyRepository() *InMemoryTechnologyRepository {
	return &InMemoryTechnologyRepository{
		technologies: getTechnologyData(),
	}
}

// GetTechnology returns a technology by name
func (r *InMemoryTechnologyRepository) GetTechnology(name string) (*models.Technology, error) {
	tech, exists := r.technologies[name]
	if !exists {
		return nil, fmt.Errorf("technology '%s' not found", name)
	}
	return &tech, nil
}

// GetTechnologies returns multiple technologies by names (optimized for batch lookups)
func (r *InMemoryTechnologyRepository) GetTechnologies(names []string) []models.Technology {
	// Pre-allocate slice with exact capacity to avoid reallocations
	techs := make([]models.Technology, 0, len(names))
	for _, name := range names {
		if tech, exists := r.technologies[name]; exists {
			techs = append(techs, tech)
		}
	}
	return techs
}

// GetAllTechnologies returns all available technologies
func (r *InMemoryTechnologyRepository) GetAllTechnologies() (map[string]models.Technology, error) {
	// Return a copy to prevent external modification
	result := make(map[string]models.Technology)
	for k, v := range r.technologies {
		result[k] = v
	}
	return result, nil
}

// getTechnologyData returns the technology to icon mapping
func getTechnologyData() map[string]models.Technology {
	return map[string]models.Technology{
		// Programming Languages
		"Go": {
			Name: "Go",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg",
		},
		"Java": {
			Name: "Java",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/java/java-original.svg",
		},
		"Python": {
			Name: "Python",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/python/python-original.svg",
		},
		"JavaScript": {
			Name: "JavaScript",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/javascript/javascript-original.svg",
		},

		// Web Frameworks
		"Spring Boot": {
			Name: "Spring Boot",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/spring/spring-original.svg",
		},
		"Gin": {
			Name: "Gin",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg",
		},
		"FastAPI": {
			Name: "FastAPI",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/fastapi/fastapi-original.svg",
		},
		"Django": {
			Name: "Django",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/django/django-plain.svg",
		},
		"Flask": {
			Name: "Flask",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/flask/flask-original.svg",
		},

		// Databases
		"PostgreSQL": {
			Name: "PostgreSQL",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg",
		},
		"MySQL": {
			Name: "MySQL",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mysql/mysql-original.svg",
		},
		"MongoDB": {
			Name: "MongoDB",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mongodb/mongodb-original.svg",
		},
		"Redis": {
			Name: "Redis",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/redis/redis-original.svg",
		},

		// Cloud & Infrastructure
		"Docker": {
			Name: "Docker",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg",
		},
		"Kubernetes": {
			Name: "Kubernetes",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/kubernetes/kubernetes-plain.svg",
		},
		"Google Cloud Platform": {
			Name: "Google Cloud Platform",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/googlecloud/googlecloud-original.svg",
		},
		"Nginx": {
			Name: "Nginx",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nginx/nginx-original.svg",
		},
		"Terraform": {
			Name: "Terraform",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/terraform/terraform-original.svg",
		},
		"Linux": {
			Name: "Linux",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/linux/linux-original.svg",
		},
		"Virtualization": {
			Name: "Virtualization",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vagrant/vagrant-original.svg",
		},
		"Helm": {
			Name: "Helm",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/helm/helm-original.svg",
		},

		// Message Queues & Streaming
		"Apache Kafka": {
			Name: "Apache Kafka",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apachekafka/apachekafka-original.svg",
		},

		// Monitoring & Observability
		"Prometheus": {
			Name: "Prometheus",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/prometheus/prometheus-original.svg",
		},
		"Grafana": {
			Name: "Grafana",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/grafana/grafana-original.svg",
		},
		"OpenTelemetry": {
			Name: "OpenTelemetry",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/opentelemetry/opentelemetry-original.svg",
		},

		"GitHub Actions": {
			Name: "GitHub Actions",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/github/github-original.svg",
		},

		// Communication Protocols
		"gRPC": {
			Name: "gRPC",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/grpc/grpc-original.svg",
		},

		// Architecture Patterns
		"Microservices": {
			Name: "Microservices",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg",
		},
		"Event Sourcing": {
			Name: "Event Sourcing",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apachekafka/apachekafka-original.svg",
		},
		"CQRS": {
			Name: "CQRS",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/graphql/graphql-plain.svg",
		},
		"Domain-Driven Design": {
			Name: "Domain-Driven Design",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/unifiedmodelinglanguage/unifiedmodelinglanguage-original.svg",
		},
		"Distributed Systems": {
			Name: "Distributed Systems",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/networkx/networkx-original.svg",
		},
		"System Integration": {
			Name: "System Integration",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/azuredevops/azuredevops-original.svg",
		},
	}
}
