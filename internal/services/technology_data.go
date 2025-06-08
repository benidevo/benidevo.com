package services

import (
	"github.com/benidevo/website/internal/models"
)

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
		"TypeScript": {
			Name: "TypeScript",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/typescript/typescript-original.svg",
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
		"AWS": {
			Name: "AWS",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/amazonwebservices/amazonwebservices-plain-wordmark.svg",
		},
		"Google Cloud": {
			Name: "Google Cloud",
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

		// Message Queues & Streaming
		"Apache Kafka": {
			Name: "Apache Kafka",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apachekafka/apachekafka-original.svg",
		},
		"RabbitMQ": {
			Name: "RabbitMQ",
			Icon: "https://www.svgrepo.com/show/303576/rabbitmq-logo.svg",
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
		"Jaeger": {
			Name: "Jaeger",
			Icon: "https://www.jaegertracing.io/img/jaeger-icon-color.svg",
		},

		// Communication Protocols
		"gRPC": {
			Name: "gRPC",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/grpc/grpc-original.svg",
		},
		"REST": {
			Name: "REST",
			Icon: "https://www.svgrepo.com/show/373903/postman.svg",
		},
		"GraphQL": {
			Name: "GraphQL",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/graphql/graphql-plain.svg",
		},

		"Spring Cloud": {
			Name: "Spring Cloud",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/spring.svg",
		},
		"Eureka": {
			Name: "Eureka",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/spring.svg",
		},
		"Hystrix": {
			Name: "Hystrix",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/netflix.svg",
		},
		"Zuul": {
			Name: "Zuul",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/netflix.svg",
		},
		"Zipkin": {
			Name: "Zipkin",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/zipkin.svg",
		},

		// Testing
		"JUnit": {
			Name: "JUnit",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/java.svg",
		},
		"Mockito": {
			Name: "Mockito",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/oracle.svg",
		},
		"Testify": {
			Name: "Testify",
			Icon: "https://cdn.jsdelivr.net/npm/simple-icons@latest/icons/go.svg",
		},

		// Build Tools & CI/CD
		"GitHub Actions": {
			Name: "GitHub Actions",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/github/github-original.svg",
		},
		"Jenkins": {
			Name: "Jenkins",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/jenkins/jenkins-original.svg",
		},

		// Architecture Patterns (using generic icons)
		"Microservices": {
			Name: "Microservices",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/kubernetes/kubernetes-plain.svg",
		},
		"Event Sourcing": {
			Name: "Event Sourcing",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apachekafka/apachekafka-original.svg",
		},
		"CQRS": {
			Name: "CQRS",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/graphql/graphql-plain.svg",
		},
		"DDD": {
			Name: "DDD",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/databricks/databricks-original.svg",
		},
	}
}
