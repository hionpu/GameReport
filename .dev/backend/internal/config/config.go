package config

import (
	"os"
)

type Config struct {
	RiotAPIKey string
	Port       string
	Region     string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		RiotAPIKey: getEnv("RIOT_API_KEY", "RIOT_API_KEY_PLACEHOLDER"),
		Port:       getEnv("PORT", "8080"),
		Region:     getEnv("REGION", "kr"), // Korea as default
		DatabaseURL: getEnv("DATABASE_URL", "user=postgres password=postgres dbname=gamereport port=5432 sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}