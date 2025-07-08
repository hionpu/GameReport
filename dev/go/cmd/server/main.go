package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"gameReport/internal/config"
	"gameReport/internal/db"
	"gameReport/internal/handlers"
	"gameReport/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("🚀 Starting gameReport server...")

	cfg := config.LoadConfig()
	// Add this after cfg := config.LoadConfig()
	log.Printf("🔍 DEBUG - Database URL: %s", cfg.DatabaseURL)
	log.Printf("🔍 DEBUG - Working directory: %s", func() string {
		wd, _ := os.Getwd()
		return wd
	}())
	if err := cfg.Validate(); err != nil {
		log.Fatalf("❌ Configuration validation failed: %v", err)
	}

	log.Printf("🔧 Environment: %s", cfg.Environment)
	log.Printf("📡 Supabase URL: %s", cfg.SupabaseURL)

	dbConfig := db.DBConfig{
		URL:             cfg.DatabaseURL,
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: 30 * time.Minute,
	}

	database, err := db.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer database.Close()

	handler, err := handlers.NewHandler(cfg, database)
	if err != nil {
		log.Fatalf("❌ Failed to initialize handlers: %v", err)
	}

	r := chi.NewRouter()

	// Middleware chain
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())
	r.Use(chimiddleware.Compress(5))
	r.Use(middleware.Timeout(30 * time.Second))

	// Routes
	r.Get("/", handler.Home)
	r.Get("/health", handler.HealthCheck)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "..", "..", "web", "static"))
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))

	r.NotFound(handler.NotFound)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("🌐 Server starting on port %s", cfg.Port)
		log.Printf("🔗 Visit: http://localhost:%s", cfg.Port)
		log.Printf("💚 Health check: http://localhost:%s/health", cfg.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Server shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("❌ Server forced to shutdown: %v", err)
		return
	}

	log.Println("✅ Server shutdown complete")
}
