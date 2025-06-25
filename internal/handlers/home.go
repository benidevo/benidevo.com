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
		Description:      "Software Engineer specializing in Distributed Systems, Microservices, and Scalable Architecture.",
		CanonicalURL:     c.Request.URL.String(),
		CurrentYear:      time.Now().Year(),
		FeaturedProjects: featuredProjects,
		SkillCategories:  skillCategories,
	}

	c.HTML(http.StatusOK, "home", data)
}
