package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI  string
	Database  string
	RulesPath string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, proceeding with environment variables")
	}

	return Config{
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Database:  getEnv("DB_NAME", "warehouse"),
		RulesPath: getEnv("RULES_PATH", "./rules/rules.yaml"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
