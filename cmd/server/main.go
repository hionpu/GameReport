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

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gameReport/internal/config"
	"gameReport/internal/db"
	"gameReport/internal/handlers"
	custommiddleware "gameReport/internal/middleware"
)

func main() {
	log.Println("🚀 Starting RaidMaster Gaming Report Card Server...")
	
	// Load configuration
	cfg := config.LoadConfig()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("❌ Configuration validation failed: %v", err)
	}
	
	log.Printf("🔧 Environment: %s", cfg.Environment)
	log.Printf("📡 Supabase URL: %s", cfg.SupabaseURL)
	
	// Initialize database
	dbConfig := db.DatabaseConfig{
		URL:             cfg.DatabaseURL,
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: 30 * time.Minute,
	}
	
	db, err := db.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer db.Close()
	
	// Initialize handlers
	handler, err := handlers.NewHandler(cfg, db)
	if err != nil {
		log.Fatalf("❌ Failed to initialize handlers: %v", err)
	}
	
	// Setup router
	r := chi.NewRouter()
	
	// Middleware
	r.Use(custommiddleware.Logger())
	r.Use(custommiddleware.Recovery())
	r.Use(custommiddleware.CORS())
	r.Use(middleware.Compress(5))
	r.Use(custommiddleware.Timeout(30 * time.Second))
	
	// Routes
	r.Get("/", handler.Home)
	r.Get("/health", handler.Health)
	
	// Static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web", "static"))
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))
	
	// 404 handler
	r.NotFound(handler.NotFound)
	
	// Server configuration
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Start server in a goroutine
	go func() {
		log.Printf("🌐 Server starting on port %s", cfg.Port)
		log.Printf("🔗 Visit: http://localhost:%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed to start: %v", err)
		}
	}()
	
	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("🛑 Server shutting down...")
	
	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("❌ Server forced to shutdown: %v", err)
	}
	
	log.Println("✅ Server shutdown complete")
} 