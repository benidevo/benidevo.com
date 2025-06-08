package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/benidevo/website/internal/handlers"
	"github.com/benidevo/website/internal/services"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(globalErrorHandler)

	router.LoadHTMLGlob("web/templates/**/*.html")
	router.Static("/static", "./web/static")

	projectService := services.NewProjectService(nil)

	homeHandler := handlers.NewHomeHandler(projectService)

	router.GET("/", homeHandler.HomePage)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().UTC(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "layouts/base.html", gin.H{
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

			c.HTML(http.StatusInternalServerError, "layouts/base.html", gin.H{
				"title":       "Something Went Wrong",
				"page":        "500",
				"currentYear": time.Now().Year(),
			})
			c.Abort()
		}
	}()

	c.Next()

	if len(c.Errors) > 0 || c.Writer.Status() == http.StatusInternalServerError {
		c.HTML(http.StatusInternalServerError, "layouts/base.html", gin.H{
			"title":       "Something Went Wrong",
			"page":        "500",
			"currentYear": time.Now().Year(),
		})
		c.Abort()
	}
}
