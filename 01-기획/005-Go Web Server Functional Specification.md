# Go Web Server Functional Specification

### 1. Overview

The Go Web Server acts as the user-facing orchestrator for the "Daily Gaming Report Card" application. It handles user requests, serves web pages, and communicates with the Python Analysis Server to retrieve processed gaming data and AI insights. It utilizes `Chi` for routing and `Templ` for HTML rendering, supporting both full page loads and HTMX partial updates.

### 2. Core Features

- **User Interface Provision**: Serves the single web page application.
    
- **Request Orchestration**: Handles user input (summoner name) and coordinates data retrieval and analysis by interacting with the Python Analysis Server.
    
- **HTML Rendering**: Renders dynamic report cards and other UI elements using Templ, with support for HTMX for partial page updates.
    
- **Static File Serving**: Serves CSS, JavaScript, and other static assets.
    
- **Health Check**: Provides a basic endpoint to check server health.
    
- **API Endpoints**: Offers endpoints for player search, report generation, user authentication (register, login, profile), and HTMX partial updates.
    

### 3. Key Functions / Handlers

#### 3.1. `main()`

- **Location**: `cmd/server/main.go`
    
- **Description**: Entry point of the Go application.
    
    - Initializes the `Chi` router and middleware (logger, recoverer, compressor).
        
    - Sets up services (e.g., `GameService`).
        
    - Registers all API routes and HTMX partial routes.
        
    - Serves static files.
        
    - Starts the HTTP server on the configured port.
        

#### 3.2. `isJSONRequest(r *http.Request) bool`

- **Location**: Shared utility function (e.g., `internal/middleware/request_type.go` or similar)
    
- **Description**: Determines if the incoming HTTP request expects a JSON response based on the `Accept` header or has a JSON `Content-Type`.
    
- **Parameters**: `*http.Request`.
    
- **Returns**: `bool` (`true` if JSON request, `false` otherwise).
    

#### 3.3. `GenerateReport(w http.ResponseWriter, r *http.Request)`

- **Location**: `internal/handlers/report_handler.go` (or `handler.go`)
    
- **Description**: Handles `POST /api/report/{summonerName}/{tagLine}` requests.
    
    - Extracts `summonerName` and `tagLine` from URL parameters.
        
    - Calls `h.gameService.GetDailyReport()` to get the report data.
        
    - Checks the `Accept` header:
        
        - If `application/json`, encodes and sends the report as JSON.
            
        - Otherwise (for HTMX/HTML), calls `h.renderReportCard()` to render and send the HTML partial.
            
- **Parameters**: `http.ResponseWriter`, `*http.Request`.
    

#### 3.4. `GetReportCardPartial(w http.ResponseWriter, r *http.Request)`

- **Location**: `internal/handlers/report_handler.go` (or `handler.go`)
    
- **Description**: Handles `GET /partials/report-card/{summonerName}/{tagLine}` requests.
    
    - Extracts `summonerName` and `tagLine` from URL parameters.
        
    - Calls `h.gameService.GetDailyReport()` to get the report data.
        
    - Calls `h.renderReportCard()` to render and send the HTML partial specifically for the report card.
        
- **Parameters**: `http.ResponseWriter`, `*http.Request`.
    

#### 3.5. `Health(w http.ResponseWriter, r *http.Request)`

- **Location**: `internal/handlers/handler.go`
    
- **Description**: Handles `GET /health` requests, returning the service health status.
    
- **Parameters**: `http.ResponseWriter`, `*http.Request`.
    

#### 3.6. `GetDailyReport(gameName, tagLine string) (*DailyReportData, error)`

- **Location**: `internal/services/game_service.go`
    
- **Description**: Orchestrates the process of getting a daily report.
    
    - Calls `gs.riotClient.GetPlayerByRiotID()` to get player info.
        
    - Checks `gs.supabase.GetTodaysReport()` for a cached report. If found, parses and returns it.
        
    - If not cached, calls `gs.generateNewReport()` to create a fresh report.
        
- **Parameters**: `gameName` (string), `tagLine` (string).
    
- **Returns**: `*DailyReportData`, `error`.
    

#### 3.7. `generateNewReport(player *Player) (*DailyReportData, error)`

- **Location**: `internal/services/game_service.go`
    
- **Description**: Generates a new daily report.
    
    - Fetches recent `matchIDs` using `gs.riotClient.GetRecentMatches()`.
        
    - Iterates through `matchIDs` to get `match` details using `gs.riotClient.GetMatchDetails()`.
        
    - Processes `match` data to calculate statistics like win rate and average KDA.
        
    - Sends analysis data to the Google AI API to generate insights. Implements a fallback if AI fails.
        
    - Saves the new report to Supabase for caching.
        
- **Parameters**: `*Player`.
    
- **Returns**: `*DailyReportData`, `error`.
    

#### 3.8. `parseAIResponse(response string) AIInsight`

- **Location**: `internal/services/game_service.go`
    
- **Description**: Parses the raw string response from the Google AI API into a structured `AIInsight` object. Provides fallback insights if parsing fails or the response is incomplete.
    
- **Parameters**: `response` (string).
    
- **Returns**: `AIInsight`.
    

#### 3.9. Riot API Client Functions

- **Location**: `internal/clients/riot_client.go` (implicitly derived)
    
- **Description**: Functions for interacting with the Riot Games API.
    
    - `GetPlayerByRiotID(gameName, tagLine string) (*RiotPlayer, error)`: Fetches player account information by Riot ID.
        
    - `GetRecentMatches(puuid string, count int) ([]string, error)`: Retrieves a list of recent match IDs for a given PUUID.
        
    - `GetMatchDetails(matchID string) (*MatchDTO, error)`: Fetches detailed match data for a given match ID.
        

#### 3.10. Supabase Client Functions

- **Location**: `internal/clients/supabase_client.go` (implicitly derived)
    
- **Description**: Functions for interacting with the Supabase database.
    
    - `GetTodaysReport(puuid string) (*DailyReport, error)`: Retrieves today's cached report for a player.
        
    - `SaveDailyReport(report *DailyReport) error`: Saves a daily report to the database.
        

#### 3.11. Google AI Client Functions

- **Location**: `internal/clients/google_ai_client.go` (implicitly derived)
    
- **Description**: Functions for interacting with the Google AI API.
    
    - `GenerateInsights(analysis *MatchAnalysis) (*AIInsight, error)`: Sends match analysis data to the AI model and receives generated insights.