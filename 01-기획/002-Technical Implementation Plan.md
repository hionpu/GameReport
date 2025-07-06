# 002 - Technical Implementation Plan

## 🏗️ Technical Architecture

### Tech Stack Selection
**Backend:** Golang (std lib + Chi router) - minimal, high performance
**Frontend:** HTMX + Templ (simple, fast, SEO-friendly)
**Mobile:** Flutter (cross-platform for iOS/Android)
**Database:** Supabase (PostgreSQL + real-time + auth + MCP integration)
**AI API:** Google AI API (Gemini models)
**External APIs:** Riot Games API
**Deployment:** Railway (backend) + Vercel (web assets)

### Project Structure

```

gaming-report-card/
├── main.go                          # 메인 서버 파일
├── go.mod                           # Go 모듈 파일
├── go.sum                           # Go 종속성 체크섬
├── static/                          # 정적 파일들
│   ├── css/
│   │   └── custom.css              # 커스텀 CSS 스타일
│   ├── js/
│   │   └── app.js                  # 추가 JavaScript (필요시)
│   └── images/
│       └── logo.png                # 로고 이미지
├── templates/                       # Go HTML 템플릿들
│   ├── layout.html                 # 기본 레이아웃
│   ├── home.html                   # 홈페이지 콘텐츠
│   └── components/                 # 재사용 가능한 컴포넌트들
│       ├── header.html             # 헤더 컴포넌트
│       ├── footer.html             # 푸터 컴포넌트
│       ├── analysis-result.html    # 분석 결과 컴포넌트
│       └── sample-report.html      # 샘플 리포트 컴포넌트
├── handlers/                        # HTTP 핸들러들 (확장시)
│   ├── home.go
│   ├── analysis.go
│   └── api.go
├── models/                          # 데이터 모델들 (확장시)
│   ├── user.go
│   ├── analysis.go
│   └── match.go
├── services/                        # 비즈니스 로직 (확장시)
│   ├── riot_api.go
│   ├── analysis.go
│   └── ai_service.go
└── config/                          # 설정 파일들 (확장시)
    ├── config.go
    └── database.go
```

### System Architecture
```
User Browser/Mobile App
↓
HTMX Frontend (Vercel) / Flutter App
↓
Golang + Chi API (Railway)
↓
┌─ Riot API  ─┐ ┌─ Google AI API ─┐
│             │ │                 │
└─ Match Data ┘ └─  AI Insights   ┘
↓                       ↓
Supabase Database + Auth + Storage
↓
MCP Server (optional for advanced features)
```

## 📝 API Endpoints Design

### Core Endpoints
```go
// Player search and basic info
GET /api/player/{summonerName}/{tagLine}

// Daily report generation
POST /api/report/{summonerName}/{tagLine}

// User management (premium features)
POST /api/auth/register
POST /api/auth/login
GET /api/user/profile

// Health check
GET /health

// HTMX partial updates
GET /partials/report-card/{summonerName}/{tagLine}
GET /partials/loading
```

### Data Flow
1. **User Request:** HTMX frontend/Flutter app sends summoner name + tag
2. **Data Fetch:** Backend queries Riot API for recent matches
3. **Analysis:** Process match data for trends
4. **AI Generation:** Send analysis to Google AI API for insights
5. **Response:** Return formatted report card to frontend
6. **Caching:** Store results in Supabase for 24 hours

## 🔌 Riot API Integration

### Required Endpoints
```go
// Account API - Get player info
GET /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}

// Match API - Get recent matches
GET /lol/match/v5/matches/by-puuid/{puuid}/ids?count=5

// Match details
GET /lol/match/v5/matches/{matchId}
```

### Rate Limiting Strategy
- **Limit:** 120 requests per 2 minutes
- **Solution:** In-memory request queue with sync.Map
- **Caching:** Store match data in Supabase for 24 hours
- **Optimization:** Batch process multiple users

### Data Processing
```go
type MatchAnalysis struct {
    WinRate        float64 `json:"win_rate"`
    AvgKDA         float64 `json:"avg_kda"`
    TrendDirection string  `json:"trend_direction"`
    BestChampion   string  `json:"best_champion"`
    WeakestArea    string  `json:"weakest_area"`
}

func analyzePlayerPerformance(matches []Match) MatchAnalysis {
    // Implementation for analyzing match data
    return MatchAnalysis{
        WinRate:        calculateWinRate(matches),
        AvgKDA:         calculateAverageKDA(matches),
        TrendDirection: detectTrend(matches),
        BestChampion:   findBestPerformingChampion(matches),
        WeakestArea:    identifyWeaknesses(matches),
    }
}
```

## 🤖 Google AI Integration

### API Usage
- **Model:** Gemini Pro (for quality) / Gemini Flash (for speed)
- **Approach:** Structured prompts with game data
- **Cost Management:** Cache identical analyses for 24 hours in Supabase

### Prompt Template
```go
func generateInsightPrompt(playerData MatchAnalysis) string {
    return fmt.Sprintf(`You are a League of Legends coach analyzing a player's recent performance.

Player Data:
- Recent Win Rate: %.1f%%
- Average KDA: %.2f
- Performance Trend: %s
- Best Champion: %s
- Games Analyzed: %d

Generate exactly 3 insights:
1. Performance Trend Analysis (1-2 sentences)
2. Champion/Role Recommendation (1-2 sentences)
3. Key Improvement Focus (1-2 sentences)

Keep each insight actionable and encouraging.`,
        playerData.WinRate,
        playerData.AvgKDA,
        playerData.TrendDirection,
        playerData.BestChampion,
        5) // assuming 5 games analyzed
}
```

### AI Response Processing
```go
type AIInsight struct {
    PerformanceTrend   string `json:"performance_trend"`
    Recommendation     string `json:"recommendation"`
    ImprovementFocus   string `json:"improvement_focus"`
    Confidence         string `json:"confidence"`
}

func parseAIResponse(response string) AIInsight {
    lines := strings.Split(strings.TrimSpace(response), "\n")
    
    var insights []string
    for _, line := range lines {
        if trimmed := strings.TrimSpace(line); trimmed != "" {
            insights = append(insights, trimmed)
        }
    }
    
    if len(insights) >= 3 {
        return AIInsight{
            PerformanceTrend: insights[0],
            Recommendation:   insights[1],
            ImprovementFocus: insights[2],
            Confidence:       "high",
        }
    }
    
    // Fallback insights
    return AIInsight{
        PerformanceTrend: "Keep playing consistently to track your improvement trends.",
        Recommendation:   "Focus on your best performing champions to climb ranks.",
        ImprovementFocus: "Work on map awareness and objective control for better results.",
        Confidence:       "medium",
    }
}
```

## 💾 Supabase Database Schema

### Core Tables
```sql
-- Users table
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255),
    subscription_tier VARCHAR(20) DEFAULT 'free',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Players table
CREATE TABLE players (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    summoner_name VARCHAR(100),
    tag_line VARCHAR(10),
    puuid VARCHAR(100) UNIQUE,
    region VARCHAR(10),
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(summoner_name, tag_line)
);

-- Daily reports table
CREATE TABLE daily_reports (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    player_id UUID REFERENCES players(id),
    report_date DATE,
    match_data JSONB,
    ai_insights JSONB,
    performance_score INTEGER,
    win_rate DECIMAL(5,2),
    avg_kda DECIMAL(5,2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(player_id, report_date)
);

-- API usage tracking
CREATE TABLE api_usage (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    endpoint VARCHAR(100),
    request_count INTEGER DEFAULT 1,
    date DATE DEFAULT CURRENT_DATE,
    UNIQUE(user_id, endpoint, date)
);
```

## 🏗️ Backend Implementation (Golang + Chi)

### Main Server Setup
```go
// cmd/server/main.go
package main

import (
    "log"
    "net/http"
    "os"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "gaming-analytics/internal/handlers"
    "gaming-analytics/internal/services"
)

func main() {
    r := chi.NewRouter()
    
    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Compress(5))
    
    // Services
    gameService := services.NewGameService()
    
    // Handlers
    h := handlers.NewHandler(gameService)
    
    // Routes
    r.Route("/api", func(r chi.Router) {
        r.Get("/player/{summonerName}/{tagLine}", h.GetPlayer)
        r.Post("/report/{summonerName}/{tagLine}", h.GenerateReport)
        r.Post("/auth/register", h.Register)
        r.Post("/auth/login", h.Login)
        r.Get("/user/profile", h.GetProfile)
    })
    
    // HTMX partials
    r.Route("/partials", func(r chi.Router) {
        r.Get("/report-card/{summonerName}/{tagLine}", h.GetReportCardPartial)
        r.Get("/loading", h.GetLoadingPartial)
    })
    
    r.Get("/health", h.Health)
    
    // Static files
    r.Handle("/*", http.FileServer(http.Dir("./web/static/")))
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
```

### Handlers Implementation
```go
// internal/handlers/handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "gaming-analytics/internal/services"
)

type Handler struct {
    gameService *services.GameService
    templates   *template.Template
}

func NewHandler(gs *services.GameService) *Handler {
    templates := template.Must(template.ParseGlob("web/templates/*.html"))
    return &Handler{
        gameService: gs,
        templates:   templates,
    }
}

func (h *Handler) GenerateReport(w http.ResponseWriter, r *http.Request) {
    summonerName := chi.URLParam(r, "summonerName")
    tagLine := chi.URLParam(r, "tagLine")
    
    if summonerName == "" || tagLine == "" {
        http.Error(w, "Missing summoner name or tag line", http.StatusBadRequest)
        return
    }
    
    report, err := h.gameService.GetDailyReport(summonerName, tagLine)
    if err != nil {
        http.Error(w, "Failed to generate report", http.StatusInternalServerError)
        return
    }
    
    // Check if request wants JSON or HTML
    if r.Header.Get("Accept") == "application/json" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(report)
        return
    }
    
    // Return HTMX partial
    h.renderReportCard(w, report)
}

func (h *Handler) GetReportCardPartial(w http.ResponseWriter, r *http.Request) {
    summonerName := chi.URLParam(r, "summonerName")
    tagLine := chi.URLParam(r, "tagLine")
    
    report, err := h.gameService.GetDailyReport(summonerName, tagLine)
    if err != nil {
        h.renderError(w, "Failed to generate report")
        return
    }
    
    h.renderReportCard(w, report)
}

func (h *Handler) renderReportCard(w http.ResponseWriter, report *services.DailyReportData) {
    data := struct {
        Player   services.Player     `json:"player"`
        Date     string             `json:"date"`
        WinRate  float64            `json:"win_rate"`
        AvgKDA   float64            `json:"avg_kda"`
        Insights services.AIInsight `json:"insights"`
    }{
        Player:   report.Player,
        Date:     time.Now().Format("January 2, 2006"),
        WinRate:  report.WinRate * 100,
        AvgKDA:   report.AvgKDA,
        Insights: report.Insights,
    }
    
    w.Header().Set("Content-Type", "text/html")
    if err := h.templates.ExecuteTemplate(w, "report-card", data); err != nil {
        http.Error(w, "Template error", http.StatusInternalServerError)
        return
    }
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":    "healthy",
        "timestamp": time.Now().ISO(),
        "version":   os.Getenv("APP_VERSION"),
    })
}
```

## 🎨 Frontend Implementation (HTMX + Templ)

### HTML Templates
```html
<!-- web/templates/base.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Daily Gaming Report Card</title>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-100 min-h-screen">
    <div class="container mx-auto px-4 py-8">
        {{template "content" .}}
    </div>
</body>
</html>

<!-- web/templates/search.html -->
{{define "content"}}
<div class="max-w-md mx-auto bg-white rounded-lg shadow-md p-6">
    <h1 class="text-2xl font-bold text-center mb-6">Daily Gaming Report Card</h1>
    
    <form hx-post="/api/report/{summonerName}/{tagLine}" 
          hx-target="#report-container" 
          hx-indicator="#loading"
          class="space-y-4">
        
        <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
                Summoner Name
            </label>
            <input type="text" 
                   name="summonerName" 
                   placeholder="Enter summoner name" 
                   class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                   required>
        </div>
        
        <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
                Tag Line
            </label>
            <input type="text" 
                   name="tagLine" 
                   placeholder="e.g., NA1" 
                   class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                   required>
        </div>
        
        <button type="submit" 
                class="w-full bg-blue-600 text-white p-3 rounded-lg hover:bg-blue-700 transition-colors">
            Generate Daily Report
        </button>
    </form>
    
    <div id="loading" class="htmx-indicator text-center mt-4">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <p class="text-gray-600 mt-2">Analyzing your performance...</p>
    </div>
    
    <div id="report-container" class="mt-6"></div>
</div>
{{end}}

<!-- web/templates/report-card.html -->
{{define "report-card"}}
<div class="bg-white rounded-lg shadow-lg p-6 mt-6">
    <div class="text-center mb-6">
        <h2 class="text-2xl font-bold text-gray-800">{{.Player.GameName}}#{{.Player.TagLine}}</h2>
        <p class="text-gray-600">Daily Report - {{.Date}}</p>
    </div>
    
    <div class="grid grid-cols-2 gap-4 mb-6">
        <div class="text-center p-4 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{printf "%.1f%%" .WinRate}}</div>
            <div class="text-sm text-gray-600">Win Rate</div>
        </div>
        <div class="text-center p-4 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{printf "%.2f" .AvgKDA}}</div>
            <div class="text-sm text-gray-600">Average KDA</div>
        </div>
    </div>
    
    <div class="space-y-4">
        <div class="p-4 bg-gray-50 rounded-lg">
            <h3 class="font-semibold text-gray-800 mb-2">📈 Performance Trend</h3>
            <p class="text-gray-700">{{.Insights.PerformanceTrend}}</p>
        </div>
        
        <div class="p-4 bg-gray-50 rounded-lg">
            <h3 class="font-semibold text-gray-800 mb-2">🎯 Recommendation</h3>
            <p class="text-gray-700">{{.Insights.Recommendation}}</p>
        </div>
        
        <div class="p-4 bg-gray-50 rounded-lg">
            <h3 class="font-semibold text-gray-800 mb-2">⚡ Focus Area</h3>
            <p class="text-gray-700">{{.Insights.ImprovementFocus}}</p>
        </div>
    </div>
    
    <button hx-get="/partials/report-card/{{.Player.GameName}}/{{.Player.TagLine}}" 
            hx-target="#report-container"
            class="w-full mt-6 bg-gray-600 text-white p-3 rounded-lg hover:bg-gray-700 transition-colors">
        🔄 Refresh Report
    </button>
</div>
{{end}}
```

## 🚀 Service Layer Implementation

### Game Service
```go
// internal/services/game_service.go
package services

import (
    "encoding/json"
    "fmt"
    "time"
    "gaming-analytics/internal/clients"
)

type GameService struct {
    riotClient  *clients.RiotClient
    supabase    *clients.SupabaseClient
    googleAI    *clients.GoogleAIClient
}

type Player struct {
    PUUID    string `json:"puuid"`
    GameName string `json:"gameName"`
    TagLine  string `json:"tagLine"`
}

type DailyReportData struct {
    Player   Player      `json:"player"`
    WinRate  float64     `json:"win_rate"`
    AvgKDA   float64     `json:"avg_kda"`
    Insights AIInsight   `json:"insights"`
    Matches  []MatchSummary `json:"matches"`
}

type AIInsight struct {
    PerformanceTrend   string `json:"performance_trend"`
    Recommendation     string `json:"recommendation"`
    ImprovementFocus   string `json:"improvement_focus"`
    Confidence         string `json:"confidence"`
}

func NewGameService() *GameService {
    return &GameService{
        riotClient: clients.NewRiotClient(),
        supabase:   clients.NewSupabaseClient(),
        googleAI:   clients.NewGoogleAIClient(),
    }
}

func (gs *GameService) GetDailyReport(gameName, tagLine string) (*DailyReportData, error) {
    // 1. Get player info from Riot API
    riotPlayer, err := gs.riotClient.GetPlayerByRiotID(gameName, tagLine)
    if err != nil {
        return nil, fmt.Errorf("failed to get player: %w", err)
    }
    
    player := &Player{
        PUUID:    riotPlayer.PUUID,
        GameName: riotPlayer.GameName,
        TagLine:  riotPlayer.TagLine,
    }
    
    // 2. Check if we have today's report cached
    if report, err := gs.supabase.GetTodaysReport(player.PUUID); err == nil {
        return gs.parseStoredReport(report, player)
    }
    
    // 3. Generate new report
    return gs.generateNewReport(player)
}

func (gs *GameService) generateNewReport(player *Player) (*DailyReportData, error) {
    // Get recent matches
    matchIDs, err := gs.riotClient.GetRecentMatches(player.PUUID, 5)
    if err != nil {
        return nil, fmt.Errorf("failed to get matches: %w", err)
    }
    
    var matches []MatchSummary
    var totalKills, totalDeaths, totalAssists int
    var wins int
    
    // Process each match
    for _, matchID := range matchIDs {
        match, err := gs.riotClient.GetMatchDetails(matchID)
        if err != nil {
            continue // Skip failed matches
        }
        
        // Find player's data in match
        for _, participant := range match.Participants {
            if participant.PUUID == player.PUUID {
                kda := float64(participant.Kills + participant.Assists)
                if participant.Deaths > 0 {
                    kda = float64(participant.Kills + participant.Assists) / float64(participant.Deaths)
                }
                
                matches = append(matches, MatchSummary{
                    Champion: participant.ChampionName,
                    KDA:      kda,
                    Win:      participant.Win,
                    Duration: match.GameDuration,
                })
                
                totalKills += participant.Kills
                totalDeaths += participant.Deaths
                totalAssists += participant.Assists
                if participant.Win {
                    wins++
                }
                break
            }
        }
    }
    
    if len(matches) == 0 {
        return nil, fmt.Errorf("no recent matches found")
    }
    
    // Calculate statistics
    winRate := float64(wins) / float64(len(matches))
    avgKDA := float64(totalKills + totalAssists)
    if totalDeaths > 0 {
        avgKDA = float64(totalKills + totalAssists) / float64(totalDeaths)
    }
    
    // Create analysis for AI
    analysis := &MatchAnalysis{
        WinRate:        winRate * 100,
        AvgKDA:         avgKDA,
        BestChampion:   gs.findBestChampion(matches),
        WeakestArea:    gs.identifyWeakness(matches, winRate, avgKDA),
        TrendDirection: gs.analyzeTrend(matches),
        GamesAnalyzed:  len(matches),
    }
    
    // Generate AI insights using Google AI API
    insights, err := gs.googleAI.GenerateInsights(analysis)
    if err != nil {
        // Fallback insights if AI fails
        insights = &AIInsight{
            PerformanceTrend: "Keep playing consistently to track your improvement trends.",
            Recommendation:   "Focus on your best performing champions to climb ranks.",
            ImprovementFocus: "Work on map awareness and objective control for better results.",
            Confidence:       "medium",
        }
    }
    
    // Save to Supabase
    report := &DailyReport{
        PlayerPUUID: player.PUUID,
        ReportDate:  time.Now().Format("2006-01-02"),
        MatchData:   matches,
        AIInsights:  insights,
        WinRate:     winRate,
        AvgKDA:      avgKDA,
    }
    
    if err := gs.supabase.SaveDailyReport(report); err != nil {
        // Log error but don't fail the request
        fmt.Printf("Failed to save report: %v\n", err)
    }
    
    return &DailyReportData{
        Player:   *player,
        WinRate:  winRate,
        AvgKDA:   avgKDA,
        Insights: *insights,
        Matches:  matches,
    }, nil
}
```

## 📱 Flutter Mobile App Structure

### Project Structure
```
flutter_app/
├── lib/
│   ├── main.dart
│   ├── models/
│   │   ├── player.dart
│   │   ├── daily_report.dart
│   │   └── ai_insight.dart
│   ├── services/
│   │   ├── api_service.dart
│   │   └── auth_service.dart
│   ├── screens/
│   │   ├── search_screen.dart
│   │   ├── report_screen.dart
│   │   └── profile_screen.dart
│   └── widgets/
│       ├── report_card.dart
│       ├── loading_indicator.dart
│       └── search_form.dart
└── pubspec.yaml
```

### Main App
```dart
// lib/main.dart
import 'package:flutter/material.dart';
import 'screens/search_screen.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Daily Gaming Report Card',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: SearchScreen(),
    );
  }
}
```

### API Service
```dart
// lib/services/api_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/daily_report.dart';

class ApiService {
  static const String baseUrl = 'https://your-app.railway.app/api';
  
  Future<DailyReport> generateReport(String summonerName, String tagLine) async {
    final response = await http.post(
      Uri.parse('$baseUrl/report/$summonerName/$tagLine'),
      headers: {'Content-Type': 'application/json'},
    );
    
    if (response.statusCode == 200) {
      return DailyReport.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to generate report');
    }
  }
}
```

## 📦 Deployment Configuration

### Dockerfile
```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/web ./web

CMD ["./main"]
```

### Railway Deployment
```toml
# railway.toml
[build]
builder = "DOCKERFILE"
dockerfilePath = "Dockerfile"

[deploy]
healthcheckPath = "/health"
healthcheckTimeout = 300
restartPolicyType = "ON_FAILURE"
restartPolicyMaxRetries = 10

[environments.production.variables]
PORT = "8080"
GO_ENV = "production"
```

### Environment Variables
```bash
# .env for development
PORT=8080
RIOT_API_KEY=RGAPI-your-key-here
GOOGLE_AI_API_KEY=your-google-ai-key-here
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key-here

# Production environment variables (Railway)
RIOT_API_KEY=
GOOGLE_AI_API_KEY=
SUPABASE_URL=
SUPABASE_ANON_KEY=
```

## 🔧 MCP Integration (Optional Advanced Feature)

### Supabase MCP Server
```go
// internal/mcp/supabase_server.go
package mcp

import (
    "context"
    "encoding/json"
    "gaming-analytics/internal/clients"
)

type MCPServer struct {
    supabase *clients.SupabaseClient
}

func NewMCPServer(supabase *clients.SupabaseClient) *MCPServer {
    return &MCPServer{supabase: supabase}
}

func (m *MCPServer) QueryReports(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Advanced querying capabilities for reports
    playerID := params["player_id"].(string)
    days := int(params["days"].(float64))
    
    reports, err := m.supabase.GetReportsForPlayer(playerID, days)
    if err != nil {
        return nil, err
    }
    
    return reports, nil
}

func (m *MCPServer) GenerateAdvancedAnalytics(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Advanced analytics using MCP
    playerID := params["player_id"].(string)
    
    analytics, err := m.supabase.GetPlayerAnalytics(playerID)
    if err != nil {
        return nil, err
    }
    
    return analytics, nil
}
```

This updated technical implementation provides a solid foundation using Golang (std + Chi), HTMX + Templ for web, Flutter for mobile, Supabase for database, and Google AI API for insights. The architecture is optimized for performance, simplicity, and scalability.

---
**Last Updated:** July 1, 2025  
**Status:** Technical Implementation Updated - Golang/Chi/HTMX/Flutter/Supabase Stack