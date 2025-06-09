package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/benidevo/website/internal/models"
	"github.com/benidevo/website/internal/services"
)

// ProjectHandler handles project-related requests
type ProjectHandler struct {
	projectService *services.ProjectService
}

// NewProjectHandler creates a new project handler
func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// GetProjectDetail renders the project detail page
func (h *ProjectHandler) GetProjectDetail(c *gin.Context) {
	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		log.Error().Err(err).Str("project_id", projectIDStr).Msg("Invalid project ID")
		c.HTML(http.StatusNotFound, "404", nil)
		return
	}

	log.Info().Int("project_id", projectID).Msg("Fetching project details")

	project := h.projectService.GetProjectByID(projectID)
	if project == nil {
		log.Warn().Int("project_id", projectID).Msg("Project not found")
		c.HTML(http.StatusNotFound, "404", nil)
		return
	}

	data := models.ProjectDetailPageData{
		Title:        project.Title,
		Description:  project.Description,
		CanonicalURL: c.Request.URL.String(),
		CurrentYear:  time.Now().Year(),
		Project:      project,
	}

	c.HTML(http.StatusOK, "project_detail", data)
}

// GetProjectDetailFragment returns a partial HTML fragment for HTMX requests
func (h *ProjectHandler) GetProjectDetailFragment(c *gin.Context) {
	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		log.Error().Err(err).Str("project_id", projectIDStr).Msg("Invalid project ID for fragment")
		c.String(http.StatusBadRequest, "Invalid project ID")
		return
	}

	log.Info().Int("project_id", projectID).Msg("Fetching project detail fragment")

	project := h.projectService.GetProjectByID(projectID)
	if project == nil {
		log.Warn().Int("project_id", projectID).Msg("Project not found for fragment")
		c.String(http.StatusNotFound, "Project not found")
		return
	}
	c.Header("HX-Trigger", "project-loaded")
	c.Header("HX-Push-Url", "/projects/"+projectIDStr)

	c.HTML(http.StatusOK, "project_detail_fragment", gin.H{
		"Project": project,
	})
}
