package main

import (
	"fmt"
	"context"
	"log/slog"
	"net/http"
	"os"

	"gameReport/internal/handlers"
	"gameReport/internal/config"
	"gameReport/internal/games"
	"gameReport/internal/infrastructure/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// gameReport.gg
func main() {
	ctx := context.Background()
	
	// Initialize configuration
	cfg := config.Load()

	// Initialize structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// initialize database connection
	dbConn, err := db.New(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// initialize game registry
	gameRegistry := games.NewGameRegistryWithRegion(cfg.RiotAPIKey, cfg.Region)

	// Initialize handlers with dependencies
	handlers.SetDependencies(gameRegistry, dbConn)

	r := chi.NewRouter()

	
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(CORSMiddleware)
	r.Use(ErrorHandlingMiddleware)

	// Static file serving
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Health check endpoint
	r.Get("/health", handlers.HealthCheck)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/player", func(r chi.Router) {
			r.Get("/{gameName}/{tagLine}", handlers.GetPlayerStats)
		})
		r.Route("/insights", func(r chi.Router) {
			r.Get("/{gameName}/{tagLine}", handlers.GetPlayerInsights)
		})
	})

	// Page routes (HTMX-based)
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.HomePage)
		r.Get("/player/{gameName}/{tagLine}", handlers.PlayerPage)
		r.Post("/search", handlers.SearchPlayer)
	})

	slog.Info("Server starting", "port", cfg.Port, "region", cfg.Region)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, HX-Request, HX-Trigger, HX-Target, HX-Current-URL")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Panic recovered", "error", err, "path", r.URL.Path)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}