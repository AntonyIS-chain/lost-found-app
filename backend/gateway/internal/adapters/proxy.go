package adapters

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// RegisterProxyRoutes sets up reverse proxy routes for a given service
func RegisterProxyRoutes(router *gin.RouterGroup, serviceBaseURL string) {
	targetURL, err := url.Parse(serviceBaseURL)
	if err != nil {
		panic("Invalid service base URL: " + serviceBaseURL)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	router.Any("/*proxyPath", func(c *gin.Context) {
		// Modify request before forwarding
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Scheme = targetURL.Scheme
		c.Request.URL.Path = c.Param("proxyPath")

		// Forward request to the target service
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}
