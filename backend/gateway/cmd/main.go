package cmd

import (
	"log"
	"net/http"

	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/adapters"
	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func GatewayServer() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.RateLimiterMiddleware())

	// Register authentication route (NO JWT or Authorization required)
	router.GET("/token", func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication credentials"})
			return
		}

		// Validate credentials
		if !middlewares.ValidateBasicAuth(username, password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT Token
		token, err := middlewares.GenerateJWT(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})
	protectedRoutes := router.Group("/gw")
	protectedRoutes.Use(middlewares.JWTMiddleware())           // Validate JWT first
	protectedRoutes.Use(middlewares.AuthorizationMiddleware()) // Then check user roles

	adapters.RegisterProxyRoutes(protectedRoutes) // Register proxy routes under /proxy

	log.Println("API Gateway running on port 8080")
	router.Run(":8000")
}
