package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"gameReport/internal/models"
	"gameReport/internal/core"
	
)


// Response types
type PlayerStatsResponse struct {
	Player     core.Player             `json:"player"`
	Matches    []core.Match            `json:"matches"`
	Analysis   core.PerformanceAnalysis `json:"analysis"`
	MatchCount int                     `json:"match_count"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	models.WriteSuccess(w, map[string]string{
		"status": "ok",
		"service": "gameReport-platform",
	})
}

func GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	gameName := chi.URLParam(r, "gameName")
	tagLine := chi.URLParam(r, "tagLine")

	if gameName == "" || tagLine == "" {
		models.WriteError(w, http.StatusBadRequest, "Missing game name or tag line")
		return
	}

	service, err := gameRegistry.GetService(core.GameType(gameName))
	if err != nil {
		slog.Error("Failed to get game service", "error", err)
		models.WriteError(w, http.StatusInternalServerError, "Failed to get game service")
		return
	}

	// Get player info
	identifier := fmt.Sprintf("%s#%s", gameName, tagLine)
	player, err := service.GetPlayer(r.Context(), identifier)
	if err != nil {
		slog.Error("Failed to get player info", "error", err)
		models.WriteError(w, http.StatusNotFound, "Failed to get player info")
		return
	}

	// Get recent matches
	matchCount := 5
	if countStr := r.URL.Query().Get("matches"); countStr != "" {
		if count, err := strconv.Atoi(countStr); err == nil && count > 0  && count <= 10 {
			matchCount = count
		}
	}

	matches, err := service.GetRecentMatches(r.Context(), player.ID, matchCount)
	if err != nil {
		slog.Error("Failed to get recent matches", "error", err, "playerID", player.ID)
		models.WriteError(w, http.StatusInternalServerError, "Failed to get recent matches")
		return
	}
	
	response := PlayerStatsResponse{
		Player: *player,
		Matches: matches,
		Analysis: *analysis,
		MatchCount: len(matches),
	}

	models.WriteSuccess(w, response)
}

func GetPlayerAnalysis(w http.ResponseWriter, r *http.Request) {
	gameName := chi.URLParam(r, "gameName")
	tagLine := chi.URLParam(r, "tagLine")

	if gameName == "" || tagLine == "" {
		models.WriteError(w, http.StatusBadRequest, "Missing game name or tag line")
		return
	}
	
	// Get player stats first
	statsResponse, err := GetPlayerStatsInternal(r.Context(), gameName, tagLine)
	if err != nil {
		slog.Error("Failed to get player stats", "error", err)
		models.WriteError(w, http.StatusInternalServerError, "Failed to get player stats")
		return
	}

	// Generate AI insights (TODO: Replace with real AI service)
	insights := generateMockInsights(statsResponse.Analysis)

	report := core.DailyReport{
		Player: statsResponse.Player,
		Game: core.GameType(gameName),
		Date: time.Now(),
		PerformanceData: statsResponse.Analysis,
		AIInsights: insights,
		CacheExpiry: core.DefaultCacheExpiry(),
	}

	models.WriteSuccess(w, report)
}

// SearchPlayer - HTMX endpoint for player search
func SearchPlayer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		models.WriteError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	gameName := r.FormValue("gameName")
	tagLine := r.FormValue("tagLine")
	
	if gameName == "" || tagLine == "" {
		models.WriteError(w, http.StatusBadRequest, "Both gameName and tagLine are required")
		return
	}

	// For HTMX requests, redirect to player page
	if models.IsHTMXRequest(r) {
		w.Header().Set("HX-Redirect", "/player/"+gameName+"/"+tagLine)
		w.WriteHeader(http.StatusOK)
		return
	}

	// For regular requests, return JSON
	models.WriteSuccess(w, map[string]string{
		"redirect": "/player/" + gameName + "/" + tagLine,
	})
}

func PlayerPage(w http.ResponseWriter, r *http.Request) {
	gameName := chi.URLParam(r, "gameName")
	tagLine := chi.URLParam(r, "tagLine")

	statsResponse, err := getPlayerStatsInternal(r.Context(), gameName, tagLine)
	if err != nil {
		// For now, return Json Error, TODO: Render error page
		slog.Error("Failed to get player stats", "error", err)
		models.WriteError(w, http.StatusNotFound, "Player not found")
		return
	}

	if models.IsHTMXRequest(r) {
		// TODO: Render player stats partial template
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(statsResponse)
		return
	}

	// For full page requests
	// TODO: Render full page template with player data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page": "player",
		"data": statsResponse,
	})
}


// HomePage Handler (existing, keeping for reference)
func HomePage(w http.ResponseWriter, r *http.Request) {
	if models.IsHTMXRequest(r) {
		// TODO: Return search form partial
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<form hx-post="/search" hx-target="#results">
				<input type="text" name="gameName" placeholder="Game Name" required>
				<input type="text" name="tagLine" placeholder="Tag Line (e.g., KR1)" required>
				<button type="submit">Search</button>
			</form>
			<div id="results"></div>
		`))
		return
	}
	
	// TODO: Return full HTML page with layout
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>GameReport - Daily Gaming Analytics</title>
			<script src="https://unpkg.com/htmx.org@1.9.6"></script>
		</head>
		<body>
			<h1>GameReport</h1>
			<form hx-post="/search" hx-target="#results">
				<input type="text" name="gameName" placeholder="Game Name" required>
				<input type="text" name="tagLine" placeholder="Tag Line (e.g., KR1)" required>
				<button type="submit">Search</button>
			</form>
			<div id="results"></div>
		</body>
		</html>
	`))
}

// Helper function to get player stats internally
func getPlayerStatsInternal(ctx context.Context, gameName, tagLine string) (*PlayerStatsResponse, error) {
	service, err := gameRegistry.GetService(core.GameTypeLoL)
	if err != nil {
		return nil, err
	}

	identifier := gameName + "#" + tagLine
	player, err := service.GetPlayer(ctx, identifier)
	if err != nil {
		return nil, err
	}

	matches, err := service.GetRecentMatches(ctx, player.ID, 5)
	if err != nil {
		return nil, err
	}

	analysis, err := service.AnalyzePerformance(matches)
	if err != nil {
		return nil, err
	}

	return &PlayerStatsResponse{
		Player:     *player,
		Matches:    matches,
		Analysis:   *analysis,
		MatchCount: len(matches),
	}, nil
}

// Mock AI insights generation (TODO: Replace with real AI service)
func generateMockInsights(analysis core.PerformanceAnalysis) core.AIInsights {
	var trend, recommendation, focus string

	if analysis.WinRate > 60 {
		trend = "You're on a winning streak! Your performance is above average."
		recommendation = "Keep playing your best champions and maintain this momentum."
		focus = "Consistency in your strong areas"
	} else if analysis.WinRate > 40 {
		trend = "Your performance is steady with room for improvement."
		recommendation = "Focus on your weakest areas while maintaining your strengths."
		focus = analysis.WeakestArea
	} else {
		trend = "Your recent performance suggests areas needing attention."
		recommendation = "Consider reviewing your gameplay fundamentals and champion choices."
		focus = "Basic mechanics and " + analysis.WeakestArea
	}

	return core.AIInsights{
		PerformanceTrend: trend,
		Recommendation:   recommendation,
		ImprovementFocus: focus,
		Confidence:       "Medium",
	}
}

