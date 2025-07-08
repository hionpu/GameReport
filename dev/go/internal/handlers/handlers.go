package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"gameReport/internal/config"
	"gameReport/internal/db"
	"gameReport/internal/models"
)

type Handler struct {
	config   *config.Config
	db       *db.DB
	template *template.Template
}

var webFolderPath = "../../web"

func NewHandler(cfg *config.Config, db *db.DB) (*Handler, error) {
	templates, err := loadTemplates()
	if err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}

	return &Handler{
		config:   cfg,
		db:       db,
		template: templates,
	}, nil
}

func loadTemplates() (*template.Template, error) {

	templatePattern := filepath.Join(webFolderPath, "templates", "*.html")
	templates, err := template.ParseGlob(templatePattern)
	if err != nil {
		return nil, err
	}

	componentPattern := filepath.Join(webFolderPath, "templates", "components", "*.html")
	if componentTemplates, err := filepath.Glob(componentPattern); err == nil && len(componentTemplates) > 0 {
		templates, err = templates.ParseGlob(componentPattern)
		if err != nil {
			return nil, err
		}
	}

	return templates, nil
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("Health check request received from %s", r.RemoteAddr)

	dbStartAt := time.Now()
	dbError := h.db.HealthCheck()
	dbLatency := time.Since(dbStartAt)

	healthStatus := models.HealthStatus{
		Status:    "OK",
		Timestamp: time.Now(),
		Version:   "1.0.0, phase 1",
		Database: models.DatabaseHealth{
			Status:      "OK",
			Connections: h.db.GetDatabaseStatistics(),
			Latency:     dbLatency.String(),
		},
		Supabase: models.SupabaseHealth{
			Status:    "OK",
			URL:       h.config.SupabaseURL,
			Connected: true,
		},
	}

	if dbError != nil {
		healthStatus.Status = "BAD"
		healthStatus.Database.Status = "BAD"
		healthStatus.Supabase.Connected = false
		healthStatus.Details = map[string]interface{}{
			"database_error": dbError.Error(),
		}

		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("❌ Health check failed: %v", dbError)
	} else {
		log.Printf("✅ Health check passed in %s", dbLatency.String())
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(healthStatus); err != nil {
		log.Printf("Error encoding health status: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Home page request received from %s", r.RemoteAddr)

	pageData := models.PageData{
		Title:       "GameReport",
		CurrentDate: time.Now().Format("2025년 7월 7일 월요일"),
		Version:     "1.0.0, phase 1",
		Data: map[string]interface{}{
			"environment":  h.config.Environment,
			"supabase_url": h.config.SupabaseURL,
			"features": []string{
				"Daily Gaming Insights",
				"Performance Metrics",
				"Progress Tracking(Supported later)",
				"AI-Powered Analysis(Supported later)",
			},
		},
	}

	if err := h.template.ExecuteTemplate(w, "layout", pageData); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Home page successfully served to %s", r.RemoteAddr)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 Not Found: %s %s", r.Method, r.URL.Path)

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")

	response := models.APIResponse{
		Success: false,
		Message: "Endpoint not found",
		Error:   fmt.Sprintf("The requested endpoint %s %s does not exist", r.Method, r.URL.Path),
	}

	json.NewEncoder(w).Encode(response)
}
