package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/config"
	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/adapters"
	"github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// NewReverseProxy creates a reverse proxy to forward requests
func NewReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// GatewayServer initializes and starts the API Gateway
func GatewayServer() {
	cfg := config.LoadConfig() // Load configuration

	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	// Apply global middleware
	router.Use(middlewares.CORSMiddleware())
	// router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.RateLimiterMiddleware())

	// Public Routes (No Authentication)
	router.POST("/auth/login", NewReverseProxy(cfg.UserService))
	router.POST("/auth/signup", NewReverseProxy(cfg.UserService))
	router.POST("/auth/refresh-token", NewReverseProxy(cfg.UserService))

	// Protected Routes (Require valid JWT)
	protectedRoutes := router.Group("/gw")
	protectedRoutes.Use(middlewares.JWTMiddleware())

	// User Service Routes
	userRoutes := protectedRoutes.Group("/user")
	adapters.RegisterProxyRoutes(userRoutes, cfg.UserService)

	// Reward Service Routes
	rewardRoutes := protectedRoutes.Group("/reward")
	adapters.RegisterProxyRoutes(rewardRoutes, cfg.RewardService)

	// Payment Service Routes
	paymentRoutes := protectedRoutes.Group("/payment")
	adapters.RegisterProxyRoutes(paymentRoutes, cfg.PaymentService)

	// Notification Service Routes
	notificationRoutes := protectedRoutes.Group("/notification")
	adapters.RegisterProxyRoutes(notificationRoutes, cfg.NotificationService)

	// Matching Service Routes
	matchingRoutes := protectedRoutes.Group("/matching")
	adapters.RegisterProxyRoutes(matchingRoutes, cfg.MatchingService)

	// Document Service Routes
	documentRoutes := protectedRoutes.Group("/document")
	adapters.RegisterProxyRoutes(documentRoutes, cfg.DocumentService)

	// Start the HTTP Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.GatewayPort),
		Handler: router,
	}

	// Graceful Shutdown
	go func() {
		log.Printf("[INFO] API Gateway running on port %s\n", cfg.GatewayPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[FATAL] Server error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("[INFO] Shutting down API Gateway...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] Server shutdown failed: %v\n", err)
	}

	log.Println("[INFO] API Gateway stopped successfully.")
}
