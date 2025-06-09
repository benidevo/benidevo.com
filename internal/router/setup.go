package router

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/russross/blackfriday/v2"

	"github.com/benidevo/website/internal/config"
	"github.com/benidevo/website/internal/handlers"
	"github.com/benidevo/website/internal/services"
)

func createMultiTemplateRenderer() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

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
		"markdown": func(text string) template.HTML {
			html := blackfriday.Run([]byte(text))
			return template.HTML(html)
		},
		"safeHTML": func(text string) template.HTML {
			return template.HTML(text)
		},
	}

	// Create templates with base layout, partials, and page content
	renderer.AddFromFilesFuncs("home", funcMap,
		"web/templates/home.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/home.html")

	renderer.AddFromFilesFuncs("project_detail", funcMap,
		"web/templates/project_detail.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/project_detail.html")

	renderer.AddFromFilesFuncs("404", funcMap,
		"web/templates/404.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/404.html")

	renderer.AddFromFilesFuncs("500", funcMap,
		"web/templates/500.html",
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/500.html")

	renderer.AddFromFilesFuncs("project_detail_fragment", funcMap,
		"web/templates/partials/project_detail_fragment.html")

	return renderer
}

// SetupRouter initializes and configures the Gin router with all dependencies
func SetupRouter(cfg *config.Config) (*gin.Engine, error) {
	services, err := services.SetupServices(cfg)
	if err != nil {
		return nil, err
	}

	handlers := handlers.SetupHandlers(services)

	router := gin.Default()

	router.Use(globalErrorHandler)
	router.Use(compressionMiddleware())

	router.HTMLRender = createMultiTemplateRenderer()
	router.Static("/static", "./web/static")

	router.GET("/", handlers.HomeHandler.HomePage)
	router.GET("/projects/:id", handlers.ProjectHandler.GetProjectDetail)
	router.GET("/projects/:id/details", handlers.ProjectHandler.GetProjectDetailFragment)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().UTC(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404", nil)
	})

	return router, nil
}

// globalErrorHandler is a Gin middleware that recovers from panics and handles internal server errors
// by rendering a 500 error page.
//
// It ensures that any unhandled errors or panics result in a
// consistent error response to the client.
func globalErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Err(err.(error)).Msg("Recovered from panic")

			c.HTML(http.StatusInternalServerError, "500", nil)
			c.Abort()
		}
	}()

	c.Next()

	if len(c.Errors) > 0 || c.Writer.Status() == http.StatusInternalServerError {
		c.HTML(http.StatusInternalServerError, "500", nil)
		c.Abort()
	}
}

// compressionMiddleware returns a Gin middleware handler that applies gzip compression
// to HTTP responses, excluding certain file extensions such as images and fonts.
func compressionMiddleware() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{
		".jpg", ".jpeg", ".png", ".gif", ".svg", ".ico", ".woff", ".woff2",
	}))
}
