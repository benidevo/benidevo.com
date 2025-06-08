package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/benidevo/website/internal/models"
	"github.com/benidevo/website/internal/services"
)

// HomeHandler handles home page related requests
type HomeHandler struct {
	projectService *services.ProjectService
}

// NewHomeHandler creates a new home handler
func NewHomeHandler(projectService *services.ProjectService) *HomeHandler {
	return &HomeHandler{
		projectService: projectService,
	}
}

// HomePage renders the home page
func (h *HomeHandler) HomePage(c *gin.Context) {
	log.Info().Msg("Rendering home page")

	featuredProjects := h.projectService.GetFeaturedProjects()
	skillCategories := h.projectService.GetSkillCategories()

	data := models.HomePageData{
		Title:            "Benjamin Idewor | Backend Engineer",
		Description:      "Backend Engineer specializing in Distributed Systems, Microservices, and Scalable Architecture. Based in Berlin, Germany.",
		CanonicalURL:     c.Request.URL.String(),
		CurrentYear:      time.Now().Year(),
		FeaturedProjects: featuredProjects,
		SkillCategories:  skillCategories,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

// ProjectDetails handles HTMX requests for project details
func (h *HomeHandler) ProjectDetails(c *gin.Context) {
	projectID := c.Param("id")

	log.Info().Str("project_id", projectID).Msg("Fetching project details")

	projectDetails, err := h.projectService.GetProjectDetails(projectID)
	if err != nil {
		log.Error().Err(err).Str("project_id", projectID).Msg("Failed to fetch project details")
		c.HTML(http.StatusNotFound, "project_error.html", gin.H{
			"message": "Project not found",
		})
		return
	}

	// Render project details partial for HTMX
	c.HTML(http.StatusOK, "project_details.html", projectDetails)
}
