package router

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/benidevo/website/internal/handlers"
	"github.com/benidevo/website/internal/repository"
	"github.com/benidevo/website/internal/services"
)

func createMultiTemplateRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Set up template functions
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"slice": func(items interface{}, start, end int) interface{} {
			switch v := items.(type) {
			case []string:
				if start >= len(v) {
					return []string{}
				}
				if end > len(v) {
					end = len(v)
				}
				return v[start:end]
			}
			return items
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}

	// Create templates with base layout, partials, and page content
	r.AddFromFilesFuncs("home", funcMap,
		"web/templates/home.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/home.html")

	r.AddFromFilesFuncs("project_detail", funcMap,
		"web/templates/project_detail.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/project_detail.html")

	r.AddFromFilesFuncs("404", funcMap,
		"web/templates/404.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/404.html")

	return r
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(globalErrorHandler)

	router.HTMLRender = createMultiTemplateRenderer()
	router.Static("/static", "./web/static")

	// Initialize repositories
	techRepo := repository.NewInMemoryTechnologyRepository()
	projectRepo := repository.NewInMemoryProjectRepository(techRepo)
	skillRepo := repository.NewInMemorySkillRepository(techRepo)

	// Initialize services
	projectService := services.NewProjectService(projectRepo, skillRepo)

	homeHandler := handlers.NewHomeHandler(projectService)
	projectHandler := handlers.NewProjectHandler(projectService)

	router.GET("/", homeHandler.HomePage)

	// Project routes
	router.GET("/projects/:id", projectHandler.GetProjectDetail)
	router.GET("/projects/:id/details", projectHandler.GetProjectDetailFragment)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().UTC(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Title":       "Page Not Found",
			"Description": "The page you're looking for doesn't exist.",
			"CurrentYear": time.Now().Year(),
		})
	})

	return router
}

// globalErrorHandler is a Gin middleware that recovers from panics and handles internal server errors
// by rendering a generic 500 error page.
//
// It ensures that any unhandled errors or panics result in a
// consistent error response to the client.
func globalErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Err(err.(error)).Msg("Recovered from panic")

			c.HTML(http.StatusInternalServerError, "404", gin.H{
				"Title":       "Something Went Wrong",
				"Description": "An internal server error occurred.",
				"CurrentYear": time.Now().Year(),
			})
			c.Abort()
		}
	}()

	c.Next()

	if len(c.Errors) > 0 || c.Writer.Status() == http.StatusInternalServerError {
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Title":       "Something Went Wrong",
			"Description": "An internal server error occurred.",
			"CurrentYear": time.Now().Year(),
			"CurrentPage": "500",
		})
		c.Abort()
	}
}
