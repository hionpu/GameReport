# Phase 1: Prototype Detailed Plan (2 Weeks)

## 🎯 Phase 1 Goal

**Complete a basic functioning demo**: A demo where searching for `TenZ#NA1` displays statistics + AI advice.

---

## 📅 Week 1: Backend Basic Structure

### Day 1-2: Project Initialization + Chi Router + templ Setup

#### Day 1 Tasks

**Project Structure Setup**:

```
valorant-ai-platform/
├── cmd/server/
│   └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── services/
│   └── templates/
├── pkg/
│   └── utils/
├── web/
│   ├── static/
│   │   ├── css/
│   │   └── js/
│   └── templates/
├── go.mod
└── README.md
```

**Go Module Initialization**: `go mod init valorant-ai-platform`

**Install Core Dependencies**:

- Chi Router: `github.com/go-chi/chi/v5`
- templ: `github.com/a-h/templ`
- Environment Variables: `github.com/joho/godotenv`
- Logging: Standard library or `slog`

**Environment Configuration**: `.env` file, define configuration structs

**Basic HTTP Server**: Start a basic server using Chi Router

**templ Setup**:

- Install templ CLI
- Create basic template structure
- Configure Go generate

#### Day 2 Tasks

**Router Structure Design**:

- API endpoint hierarchical structure
- Page route structure (HTMX-based)

**Middleware Configuration**: Logging, CORS, error handling

**Health Check Endpoint**: Implement `/health` API

**Basic Response Structure**: Standardize JSON responses + HTML partial responses

**Logging System**: Structured log output

**templ Basic Templates**:

- Layout template
- Component system structure
- HTMX integration setup

### Day 3-4: Riot API Client Implementation

#### Day 3 Tasks

**Riot API Client Struct Design**:

Go

```
type RiotClient struct {
    APIKey     string
    BaseURL    string
    HTTPClient *http.Client
}
```

**API Key Management**: Securely load from environment variables

**HTTP Client Configuration**: Timeout, retry logic

**Rate Limiting Basic Structure**: Control request intervals

**Error Handling**: Map Riot API error codes

#### Day 4 Tasks

**Player Search API**: Implement Account-v1 API

**Player Information Lookup**: Account by RiotID

**Basic Data Structs**: Player, Account models

**API Response Parsing**: JSON unmarshalling

**Test Cases**: Test with actual players

### Day 5-7: Player Search + Basic Statistics API

#### Day 5 Tasks

**Match History API**: Implement Match-v1 API

**Recent Match List**: Retrieve the last 10 matches

**Match Details**: Parse individual match data

**Statistics Calculation Logic**: Win rate, average Kill/Death/Assist

**Data Structs**: Match, MatchStats models

#### Day 6 Tasks

**HTTP Handler Implementation**:

- `/api/player/{gameName}/{tagLine}` API endpoint
- `/player/{gameName}/{tagLine}` page route (HTMX-based)

**Player Statistics Aggregation**: Combine data from multiple matches

**Agent-specific Statistics**: Agent pick rate, win rate

**Map-specific Statistics**: Map performance analysis

**Response Handling**: JSON response + HTML component response

#### Day 7 Tasks

**Enhanced Error Handling**: Handle non-existent players, API errors

**Improved Logging**: Request tracking, performance logging

**Basic Caching**: Simple memory-based caching

**API Testing**: Verify endpoints with curl

**Code Cleanup**: Refactoring, adding comments

---

## 📅 Week 2: AI Integration + templ/HTMX UI

### Day 8-10: AI API Integration + Insight Generation

#### Day 8 Tasks

**AI Client Design**: OpenAI or Claude API client

**Prompt Engineering**: Prompts for game statistics-based advice

**API Key Management**: AI service authentication setup

**Basic Prompt Template**: Player performance analysis template

**AI Response Parsing**: Extract structured insights

#### Day 9 Tasks

**Statistics-Insight Connection**: Convert game statistics into AI prompts

**Insight Categories**: Improvement points, strengths, agent recommendations

**Prompt Optimization**: Generate more accurate and useful advice

**Error Handling**: AI API errors, response validation

**Response Caching**: Prevent duplicate requests for the same player

#### Day 10 Tasks

**AI Insight API**: `/api/insights/{gameName}/{tagLine}` endpoint

**Integrated API**: Retrieve statistics + insights in one call

**Performance Optimization**: Parallel processing, reduce response time

**Testing**: Verify with various player profiles

**Logging**: Track AI requests/responses

### Day 11-14: templ + HTMX Web Interface

#### Day 11 Tasks

**templ Template Structure**:

- `layout.templ`: Basic layout (including HTMX)
- `search.templ`: Search form component
- `player-stats.templ`: Player statistics component
- `insights.templ`: AI insights component

**HTMX Configuration**:

- Load HTMX from CDN
- Configure basic HTMX settings
- Setup error handling

**CSS Styling**:

- Tailwind CSS or simple custom CSS
- Responsive layout
- Visual elements matching the game theme

**Static File Serving**: Serve CSS/JS using Chi Router

#### Day 12 Tasks

**HTMX Search Functionality**:

- Search requests with `hx-post` or `hx-get`
- Update result area with `hx-target`
- Set search events with `hx-trigger`
- Loading indicator with `hx-indicator`

**Dynamic Content Rendering**:

- Return search results as partial HTML
- Structured rendering with templ components
- Configure HTMX swap strategy

**Error Handling**:

- Handle HTMX error events
- User-friendly error message components

#### Day 13 Tasks

**Statistics Visualization**:

- Chart.js or HTML/CSS based charts
- Chart updates integrated with HTMX
- Win rate donut chart
- Agent usage bar chart
- Recent performance line chart

**Insight Display**:

- Card-style AI advice components
- Dynamic insight loading with HTMX
- Display by insight category

**HTMX Advanced Features**:

- Improve page navigation with `hx-boost`
- Send additional data with `hx-vals`
- Out-of-band updates with `hx-swap-oob`

#### Day 14 Tasks

**UI/UX Improvement**:

- HTMX animations and transition effects
- User experience optimization
- Accessibility considerations

**Search History**:

- Link HTMX with browser history
- Component to display recent searches

**Share Feature**:

- URL-based player information sharing
- HTMX `pushUrl` for browser history management

**Final Testing**:

- Verify overall flow
- Test HTMX functionality
- Test on various browsers
- Mobile responsive testing

---

## 🎯 Definition of Done

### Functional Completion Criteria

- [ ] Player search (`TenZ#NA1`) works correctly
- [ ] Basic statistics (win rate, KDA, agent performance) displayed
- [ ] AI-based personalized insights generated
- [ ] HTMX-based responsive web interface
- [ ] Basic usage possible on mobile

### Technical Completion Criteria

- [ ] Go backend server runs stably
- [ ] Riot API integrated and data parsed
- [ ] AI API integrated and insights generated
- [ ] templ template system implemented
- [ ] HTMX-based interaction implemented
- [ ] Basic error handling and logging
- [ ] Performance optimized with simple caching

### Quality Criteria

- [ ] Average response time within 3 seconds
- [ ] Common error scenarios handled appropriately
- [ ] Code commented and README documented
- [ ] Git commit history organized
- [ ] Basic security considerations applied
- [ ] HTMX accessibility considerations applied

---

## 🛠️ Technology Stack

### Backend

- **Language**: Go 1.21+
- **Router**: Chi v5
- **Templating**: templ (type-safe HTML templates)
- **HTTP Client**: Standard library `net/http`
- **JSON Processing**: Standard library `encoding/json`
- **Environment Configuration**: `godotenv` package

### Frontend

- **HTML**: Type-safe HTML generated by templ
- **CSS**: Tailwind CSS or custom CSS
- **JavaScript**: HTMX (minimal JavaScript)
- **Charts**: Chart.js (if needed) or CSS-based charts
- **HTTP Requests**: HTMX built-in features

### HTMX Usage

- **Dynamic Content Loading**: `hx-get`, `hx-post`
- **Partial Page Updates**: `hx-target`, `hx-swap`
- **User Interaction**: `hx-trigger`
- **Loading State**: `hx-indicator`
- **History Management**: `hx-boost`, `hx-push-url`

### External Services

- **Riot Games API**: Player and match data
- **OpenAI API or Claude API**: AI insight generation

---

## 📋 Daily Checklist Example

### Day 1 Checklist

- [ ] Go project initialized (`go mod init`)
- [ ] Directory structure created
- [ ] Chi router dependency installed
- [ ] templ CLI installed and configured
- [ ] Basic HTTP server confirmed to be running
- [ ] Environment variable configuration structure implemented
- [ ] Git repository initialized and first commit made

### Day 7 Checklist

- [ ] Player search API completed
- [ ] Basic statistics calculation logic implemented
- [ ] JSON response format standardized
- [ ] Error handling cases implemented
- [ ] API tested with curl
- [ ] Code reviewed and refactored

### Day 14 Checklist (Final)

- [ ] `TenZ#NA1` search → statistics + AI advice display confirmed
- [ ] HTMX-based web interface fully functional
- [ ] Tested on mobile browsers
- [ ] Error cases handled in a user-friendly manner
- [ ] templ templates optimized
- [ ] README document created
- [ ] Demo video or screenshots prepared

---

## 🚀 Next Steps (Phase 2 Preparation)

Preparations for Phase 2 MVP development after completing Phase 1:

- PostgreSQL schema design
- Database integration structure design (in conjunction with templ)
- Plan for applying advanced HTMX patterns
- User interface improvement plan (templ component extension)
- Identify performance optimization points
- Build a templ/HTMX-based component library

---

**Last updated**: 2025-06-18