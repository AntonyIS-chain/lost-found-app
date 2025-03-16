package adapters

import (
	"log"
	"net/http"
	"strings"

	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/app"
	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/domain"
	"github.com/gin-gonic/gin"
)

// Mapping routes to backend services
var backendMapping = map[string]string{
	"/reward-service":       "http://reward-service",
	"/payment-service":      "http://payment-service",
	"/notification-service": "http://notification-service",
	"/matching-service":     "http://matching-service",
	"/document-service":     "http://document-service",
	"/user-service":         "http://user-service",
}

// respondWithError provides a consistent error response format
func respondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

// getTargetService determines the backend service based on request path
func getTargetService(path string) (string, string) {
	for prefix, backendURL := range backendMapping {
		if strings.HasPrefix(path, prefix) {
			trimmedPath := strings.TrimPrefix(path, prefix)
			if !strings.HasPrefix(trimmedPath, "/") {
				trimmedPath = "/" + trimmedPath // Ensure leading slash
			}

			return backendURL, trimmedPath
		}
	}
	return "http://localhost:8000", path
}

// RegisterProxyRoutes sets up authentication and proxy routes.
func RegisterProxyRoutes(router *gin.RouterGroup) {
	proxyService := app.NewProxyService()

	router.Any("/*any", func(c *gin.Context) {
		requestPath := c.Param("any")
		targetBaseURL, adjustedPath := getTargetService(requestPath)
		targetURL := targetBaseURL + adjustedPath

		request := domain.ProxyRequest{
			Method:  c.Request.Method,
			Path:    targetURL,
			Headers: map[string]string{"Authorization": c.GetHeader("Authorization")},
		}

		log.Printf("Proxying request to: %s", targetURL)

		response, err := proxyService.ForwardRequest(request)
		if err != nil {
			log.Printf("Proxy request error: %v | Target: %s", err, targetURL)
			respondWithError(c, http.StatusInternalServerError, "Failed to proxy request")
			return
		}

		contentType := response.Headers["Content-Type"]
		if contentType == "" {
			contentType = "application/json"
		}

		c.Data(response.StatusCode, contentType, response.Body)
	})

}
