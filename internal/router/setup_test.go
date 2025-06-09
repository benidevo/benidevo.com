package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCompressionMiddleware(t *testing.T) {
	middleware := compressionMiddleware()
	assert.NotNil(t, middleware)

	// Test that middleware is a valid gin.HandlerFunc
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware)

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test response")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
