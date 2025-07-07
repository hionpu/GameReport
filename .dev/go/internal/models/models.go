package models

import (
	"time"
)

// HealthStatus represents the health status of the application
type HealthStatus struct {
	Status    string                 `json:"status"`
	Timestamp time.Time              `json:"timestamp"`
	Version   string                 `json:"version"`
	Database  DatabaseHealth         `json:"database"`
	Supabase  SupabaseHealth         `json:"supabase"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// DatabaseHealth represents database health information
type DatabaseHealth struct {
	Status      string                 `json:"status"`
	Connections map[string]interface{} `json:"connections"`
	Latency     string                 `json:"latency"`
}

// SupabaseHealth represents Supabase-specific health information
type SupabaseHealth struct {
	Status    string `json:"status"`
	URL       string `json:"url"`
	Connected bool   `json:"connected"`
}

// PageData represents data passed to HTML templates
type PageData struct {
	Title       string                 `json:"title"`
	CurrentDate string                 `json:"current_date"`
	Version     string                 `json:"version"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}