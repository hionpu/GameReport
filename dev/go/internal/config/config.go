package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
	SupabaseURL string
	SupabaseAnonKey string
	SupabaseServiceRoleKey string
	Environment string
}

func LoadConfig() (*Config) {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	config := &Config{
		Port:               getEnv("PORT", "8080"),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		SupabaseURL:        getEnv("SUPABASE_URL", ""),
		SupabaseAnonKey:    getEnv("SUPABASE_ANON_KEY", ""),
		SupabaseServiceRoleKey: getEnv("SUPABASE_SERVICE_ROLE_KEY", ""),
		Environment:        getEnv("GO_ENV", "development"),
	}

	// Set default Supabase values if not provided
	if config.SupabaseURL == "" {
		config.SupabaseURL = "https://fssbljnxonqzwctasvjk.supabase.co"
	}
	if config.SupabaseAnonKey == "" {
		config.SupabaseAnonKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZzc2Jsam54b25xendjdGFzdmprIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTE4NjYxNjksImV4cCI6MjA2NzQ0MjE2OX0._ecGHtXHup28xAW6svhlVZw4LUzCzQj1vVzxoud5_I4"
	}
	if config.SupabaseServiceRoleKey == "" {
		config.SupabaseServiceRoleKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZzc2Jsam54b25xendjdGFzdmprIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc1MTg2NjE2OSwiZXhwIjoyMDY3NDQyMTY5fQ.UHfkcjGAjG5I269PQPLYyKNs9dsLXWYBApbTCZX-ygk"
	}
	if config.DatabaseURL == "" {
		config.DatabaseURL = "postgresql://postgres:ckathwn2%40@db.fssbljnxonqzwctasvjk.supabase.co:5432/postgres"
	}

	return config
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func (c *Config) Validate() error {
	if c.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	if c.SupabaseURL == "" {
		return fmt.Errorf("SUPABASE_URL is required")
	}
	return nil
}
