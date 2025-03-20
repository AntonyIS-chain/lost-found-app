package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	GatewayPort       string
	UserService       string
	RewardService     string
	PaymentService    string
	NotificationService string
	MatchingService   string
	DocumentService   string
	AuthSecret        string
	EnableLogging     bool
}

// LoadConfig loads configuration based on the environment
func LoadConfig() *Config {
	env := os.Getenv("ENV") // Set ENV=development or ENV=production

	// Choose the correct .env file
	var envFile string
	if env == "production" {
		envFile = ".env.production"
	} else {
		envFile = ".env.development"
	}

	// Load .env file
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("[ERROR] Failed to load environment file: %v", err)
	}

	// Read values from environment variables
	return &Config{
		GatewayPort:       os.Getenv("GATEWAY_PORT"),
		UserService:       os.Getenv("USER_SERVICE"),
		RewardService:     os.Getenv("REWARD_SERVICE"),
		PaymentService:    os.Getenv("PAYMENT_SERVICE"),
		NotificationService: os.Getenv("NOTIFICATION_SERVICE"),
		MatchingService:   os.Getenv("MATCHING_SERVICE"),
		DocumentService:   os.Getenv("DOCUMENT_SERVICE"),
		AuthSecret:        os.Getenv("AUTH_SECRET"),
		EnableLogging:     os.Getenv("ENABLE_LOGGING") == "true",
	}
}
