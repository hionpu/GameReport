# Go Web Server API Specification

## 1. Function Specification: Player Search and Basic Info

- **Function Name**: GetPlayer
    
- **Description**: Retrieves basic player information based on summoner name and tag line. This function acts as an initial lookup to get the player's PUUID before generating a full report.
    
- **Key Business Logic**:
    
    - Receive `summonerName` and `tagLine` from the request.
        
    - Call the internal `GameService` to obtain player details, which in turn communicates with the Python Analysis Server or a direct Riot API client (if implemented in Go).
        
    - Respond with the player's basic information.
        
- **Validation Rules**:
    
    - `summonerName` and `tagLine` must be present in the URL path.
        

### 2. API Endpoint

#### HTTP Method and Path

```
GET /api/player/{summonerName}/{tagLine}
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json
```

#### Request Body (for POST/PUT)

N/A (GET request)

#### Response Specification

**Web Response (HTML)**

```html
<div id="player-info-container">
    <p>Player: {summonerName}#{tagLine}</p>
    </div>
```

**Mobile Response (JSON)**

```json
{
  "success": true,
  "data": {
    "puuid": "string",
    "gameName": "string",
    "tagLine": "string"
  },
  "message": "Player info retrieved successfully"
}
```

### 3. Go Handler Structure

```go
func (h *Handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
    summonerName := chi.URLParam(r, "summonerName")
    tagLine := chi.URLParam(r, "tagLine")

    if summonerName == "" || tagLine == "" {
        http.Error(w, "Missing summoner name or tag line", http.StatusBadRequest)
        return
    }

    // In a real scenario, this would call the Python Analysis Server
    // or a dedicated Riot API client in Go to fetch player info.
    // For this spec, we simulate the data.
    player := struct {
        PUUID    string `json:"puuid"`
        GameName string `json:"gameName"`
        TagLine  string `json:"tagLine"`
    }{
        PUUID:    "sample_puuid_" + summonerName + tagLine,
        GameName: summonerName,
        TagLine:  tagLine,
    }

    if isJSONRequest(r) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": true,
            "data":    player,
            "message": "Player info retrieved successfully",
        })
    } else {
        w.Header().Set("Content-Type", "text/html")
        // Assuming a Templ component `PlayerInfo`
        // playerInfoComponent := PlayerInfo(player)
        // playerInfoComponent.Render(r.Context(), w)
        fmt.Fprintf(w, `<div id="player-info-container"><p>Player: %s#%s</p><button hx-post="/api/report/%s/%s" hx-target="#report-container">Generate Report</button></div>`, summonerName, tagLine, summonerName, tagLine)
    }
}
```

### 4. Error Handling

- **Missing Parameters**:
    
    - **Web Response (HTML)**: Returns HTTP 400 Bad Request with a plain text error message "Missing summoner name or tag line".
        
    - **Mobile Response (JSON)**: Returns HTTP 400 Bad Request with JSON:
        
        
        
        ```JSON
        {
          "success": false,
          "error": {
            "code": "BAD_REQUEST",
            "message": "Missing summoner name or tag line"
          }
        }
        ```
        
- **Player Not Found**:
    
    - **Web Response (HTML)**: Returns HTTP 404 Not Found with a plain text error message "Player not found".
        
    - **Mobile Response (JSON)**: Returns HTTP 404 Not Found with JSON:
        
        
        
        ```JSON
        {
          "success": false,
          "error": {
            "code": "PLAYER_NOT_FOUND",
            "message": "Player with the provided summoner name and tag line not found"
          }
        }
        ```
        
- **Internal Server Error**:
    
    - **Web Response (HTML)**: Returns HTTP 500 Internal Server Error with a plain text error message "Internal server error".
        
    - **Mobile Response (JSON)**: Returns HTTP 500 Internal Server Error with JSON:
        
        
        
        ```JSON
        {
          "success": false,
          "error": {
            "code": "INTERNAL_SERVER_ERROR",
            "message": "An unexpected error occurred"
          }
        }
        ```
        

---

## 1. Function Specification: Daily Report Generation

- **Function Name**: GenerateReport
    
- **Description**: Triggers the generation of a personalized daily report card for a given summoner name and tag line. This involves fetching recent match data, performing analysis, and generating AI-powered insights.
    
- **Key Business Logic**:
    
    - Extract `summonerName` and `tagLine` from the request.
        
    - Call the internal `GameService.GetDailyReport()` to initiate the report generation process.
        
    - The `GameService` will:
        
        - Get player info from Riot API.
            
        - Check for a cached report for today; if available, return it.
            
        - If no cached report, fetch recent match IDs and details.
            
        - Process match data to calculate statistics like win rate and average KDA.
            
        - Send analysis data to the Google AI API to generate insights.
            
        - Save the new report to Supabase for caching.
            
    - Render the report data as an HTMX partial for web requests or JSON for mobile requests.
        
- **Validation Rules**:
    
    - `summonerName` and `tagLine` must be provided.
        

### 2. API Endpoint

#### HTTP Method and Path

```
POST /api/report/{summonerName}/{tagLine}
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json
```

#### Request Body (for POST/PUT)

**Web (Form Data)**

```
summonerName={summonerName}&tagLine={tagLine}
```

**Mobile (JSON)**

JSON

JSON

```
{
  "summonerName": "PlayerName",
  "tagLine": "NA1"
}
```

#### Response Specification

**Web Response (HTML)**

```HTML
<div id="report-container">
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
</div>
```

**Mobile Response (JSON)**



```JSON
{
  "success": true,
  "data": {
    "player": {
      "puuid": "string",
      "gameName": "string",
      "tagLine": "string"
    },
    "win_rate": 0.55,
    "avg_kda": 3.45,
    "insights": {
      "performance_trend": "string",
      "recommendation": "string",
      "improvement_focus": "string",
      "confidence": "string"
    },
    "matches": [
      {
        "champion": "string",
        "kda": 0.0,
        "win": true,
        "duration": 0
      }
    ]
  },
  "message": "Daily report generated successfully"
}
```

### 3. Go Handler Structure



```Go
func (h *Handler) GenerateReport(w http.ResponseWriter, r *http.Request) {
    summonerName := chi.URLParam(r, "summonerName")
    tagLine := chi.URLParam(r, "tagLine")

    if summonerName == "" || tagLine == "" {
        http.Error(w, "Missing summoner name or tag line", http.StatusBadRequest)
        return
    }

    report, err := h.gameService.GetDailyReport(summonerName, tagLine)
    if err != nil {
        // Error handling needs to differentiate based on JSON/HTML request type
        if isJSONRequest(r) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   map[string]string{"message": "Failed to generate report: " + err.Error()},
            })
        } else {
            // h.renderError(w, "Failed to generate report: " + err.Error()) // A helper to render HTML error
            fmt.Fprintf(w, `<div id="report-container" class="text-red-500"><p>Failed to generate report: %s</p></div>`, err.Error())
        }
        return
    }

    // Check if request wants JSON or HTML
    if isJSONRequest(r) { // Re-using the common helper
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": true,
            "data":    report,
            "message": "Daily report generated successfully",
        })
        return
    }

    // Return HTMX partial
    // h.renderReportCard(w, report) // Assuming renderReportCard exists and takes report data
    // Example of direct Templ rendering if Templ functions are available
    data := struct {
        Player   services.Player
        Date     string
        WinRate  float64
        AvgKDA   float64
        Insights services.AIInsight
    }{
        Player:   report.Player,
        Date:     time.Now().Format("January 2, 2006"),
        WinRate:  report.WinRate * 100,
        AvgKDA:   report.AvgKDA,
        Insights: report.Insights,
    }
    w.Header().Set("Content-Type", "text/html")
    // Assuming 'report-card' is a Templ component or html template
    // templ.ExecuteTemplate(w, "report-card", data) or TemplComponent.Render(r.Context(), w)
    fmt.Fprint(w, `<div id="report-container">...rendered report card HTML...</div>`) // Placeholder
}
```

### 4. Error Handling

- **Missing Parameters**:
    
    - **Web Response (HTML)**: Returns HTTP 400 Bad Request with a plain text error message "Missing summoner name or tag line".
        
    - **Mobile Response (JSON)**: Returns HTTP 400 Bad Request with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "BAD_REQUEST",
            "message": "Missing summoner name or tag line"
          }
        }
        ```
        
- **Failed to Generate Report (e.g., Riot API error, AI error, internal processing error)**:
    
    - **Web Response (HTML)**: Returns HTTP 500 Internal Server Error with a plain text error message "Failed to generate report". HTMX target will show this message.
        
    - **Mobile Response (JSON)**: Returns HTTP 500 Internal Server Error with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "REPORT_GENERATION_FAILED",
            "message": "Failed to generate report: [detailed error message]"
          }
        }
        ```
        
- **Player Data Not Found (e.g., invalid summoner name or tag line)**:
    
    - **Web Response (HTML)**: Returns HTTP 404 Not Found with a plain text error message "Player not found or no recent matches".
        
    - **Mobile Response (JSON)**: Returns HTTP 404 Not Found with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "PLAYER_DATA_NOT_FOUND",
            "message": "Player data not found or no recent matches available"
          }
        }
        ```
        

---

## 1. Function Specification: User Registration

- **Function Name**: Register
    
- **Description**: Registers a new user for accessing premium features.
    
- **Key Business Logic**:
    
    - Receive user registration details (email, password) from the request.
        
    - Hash the password securely.
        
    - Store user information in the `users` table in Supabase.
        
    - Assign default subscription tier ('free').
        
    - Upon successful registration, provide feedback to the user.
        
- **Validation Rules**:
    
    - `email` must be a valid email format and unique.
        
    - `password` must meet complexity requirements (e.g., minimum length, special characters).
        

### 2. API Endpoint

#### HTTP Method and Path

```
POST /api/auth/register
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json
```

#### Request Body (for POST/PUT)

**Web (Form Data)**

```
email=user@example.com&password=securepassword123
```

**Mobile (JSON)**

JSON

JSON

```
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

#### Response Specification

**Web Response (HTML)**

HTML

HTML

```
<div id="auth-message" class="text-green-500">
    <p>Registration successful! Please log in.</p>
    </div>
```

**Mobile Response (JSON)**

JSON

JSON

```
{
  "success": true,
  "data": {
    "user_id": "uuid",
    "email": "user@example.com",
    "subscription_tier": "free"
  },
  "message": "User registered successfully"
}
```

### 3. Go Handler Structure

Go

Go

```
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    var email, password string

    if isJSONRequest(r) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        email = req.Email
        password = req.Password
    } else {
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Failed to parse form", http.StatusBadRequest)
            return
        }
        email = r.FormValue("email")
        password = r.FormValue("password")
    }

    // Basic validation
    if email == "" || password == "" {
        // Error handling should be consistent with the spec
        if isJSONRequest(r) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   map[string]string{"code": "INVALID_INPUT", "message": "Email and password are required"},
            })
        } else {
            http.Error(w, "Email and password are required", http.StatusBadRequest)
        }
        return
    }
    // More complex validation (email format, password strength) would go here

    // Simulate user registration
    userID, err := h.gameService.RegisterUser(email, password) // This would interact with Supabase auth
    if err != nil {
        if isJSONRequest(r) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest) // Or 409 Conflict for duplicate email
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   map[string]string{"code": "REGISTRATION_FAILED", "message": "Registration failed: " + err.Error()},
            })
        } else {
            fmt.Fprintf(w, `<div id="auth-message" class="text-red-500"><p>Registration failed: %s</p></div>`, err.Error())
        }
        return
    }

    if isJSONRequest(r) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": true,
            "data":    map[string]string{"user_id": userID, "email": email, "subscription_tier": "free"},
            "message": "User registered successfully",
        })
    } else {
        w.Header().Set("Content-Type", "text/html")
        // Render success message or redirect for HTMX
        fmt.Fprintf(w, `<div id="auth-message" class="text-green-500"><p>Registration successful! Please log in.</p></div>`)
    }
}
```

### 4. Error Handling

- **Invalid Input (e.g., missing email/password, invalid email format, weak password)**:
    
    - **Web Response (HTML)**: Returns HTTP 400 Bad Request with a plain text error message, or an HTMX partial updating an error message div on the form.
        
    - **Mobile Response (JSON)**: Returns HTTP 400 Bad Request with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "INVALID_INPUT",
            "message": "Invalid email or password. [Specific validation error]"
          }
        }
        ```
        
- **Email Already Exists**:
    
    - **Web Response (HTML)**: Returns HTTP 409 Conflict with a plain text error message, or an HTMX partial updating an error message div.
        
    - **Mobile Response (JSON)**: Returns HTTP 409 Conflict with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "EMAIL_ALREADY_EXISTS",
            "message": "The provided email is already registered."
          }
        }
        ```
        
- **Internal Server Error**:
    
    - **Web Response (HTML)**: Returns HTTP 500 Internal Server Error with a plain text error message "Internal server error".
        
    - **Mobile Response (JSON)**: Returns HTTP 500 Internal Server Error with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "INTERNAL_SERVER_ERROR",
            "message": "An unexpected error occurred during registration."
          }
        }
        ```
        

---

## 1. Function Specification: User Login

- **Function Name**: Login
    
- **Description**: Authenticates an existing user.
    
- **Key Business Logic**:
    
    - Receive user credentials (email, password).
        
    - Verify credentials against stored hashed passwords.
        
    - If successful, generate and return an authentication token (e.g., JWT).
        
    - If using Supabase, leverage its authentication service.
        
- **Validation Rules**:
    
    - `email` and `password` must be provided.
        

### 2. API Endpoint

#### HTTP Method and Path

```
POST /api/auth/login
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json
```

#### Request Body (for POST/PUT)

**Web (Form Data)**

```
email=user@example.com&password=securepassword123
```

**Mobile (JSON)**

JSON

JSON

```
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

#### Response Specification

**Web Response (HTML)**

HTML

HTML

```
<div id="auth-message" class="text-green-500">
    <p>Login successful! Welcome back.</p>
    </div>
```

**Mobile Response (JSON)**

JSON

JSON

```
{
  "success": true,
  "data": {
    "user_id": "uuid",
    "email": "user@example.com",
    "subscription_tier": "premium",
    "access_token": "jwt_token_string",
    "expires_in": 3600
  },
  "message": "User logged in successfully"
}
```

### 3. Go Handler Structure

Go

Go

```
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    var email, password string

    if isJSONRequest(r) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        email = req.Email
        password = req.Password
    } else {
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Failed to parse form", http.StatusBadRequest)
            return
        }
        email = r.FormValue("email")
        password = r.FormValue("password")
    }

    if email == "" || password == "" {
        if isJSONRequest(r) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   map[string]string{"code": "BAD_REQUEST", "message": "Email and password are required"},
            })
        } else {
            http.Error(w, "Email and password are required", http.StatusBadRequest)
        }
        return
    }

    // Simulate user login
    authResult, err := h.gameService.LoginUser(email, password) // This would interact with Supabase auth
    if err != nil {
        if isJSONRequest(r) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusUnauthorized) // For invalid credentials
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   map[string]string{"code": "INVALID_CREDENTIALS", "message": "Login failed: " + err.Error()},
            })
        } else {
            fmt.Fprintf(w, `<div id="auth-message" class="text-red-500"><p>Login failed: %s</p></div>`, err.Error())
        }
        return
    }

    if isJSONRequest(r) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": true,
            "data":    authResult, // Contains user_id, email, subscription_tier, access_token
            "message": "User logged in successfully",
        })
    } else {
        w.Header().Set("Content-Type", "text/html")
        // On successful login, HTMX might trigger a page reload or swap to a dashboard component
        fmt.Fprintf(w, `<div id="auth-message" class="text-green-500"><p>Login successful! Welcome back.</p></div>`)
    }
}
```

### 4. Error Handling

- **Invalid Credentials (e.g., wrong email/password)**:
    
    - **Web Response (HTML)**: Returns HTTP 401 Unauthorized with a plain text error message, or an HTMX partial updating an error message div.
        
    - **Mobile Response (JSON)**: Returns HTTP 401 Unauthorized with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "INVALID_CREDENTIALS",
            "message": "Invalid email or password."
          }
        }
        ```
        
- **Missing Credentials**:
    
    - **Web Response (HTML)**: Returns HTTP 400 Bad Request with a plain text error message.
        
    - **Mobile Response (JSON)**: Returns HTTP 400 Bad Request with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "BAD_REQUEST",
            "message": "Email and password are required."
          }
        }
        ```
        
- **Internal Server Error**:
    
    - **Web Response (HTML)**: Returns HTTP 500 Internal Server Error with a plain text error message "Internal server error".
        
    - **Mobile Response (JSON)**: Returns HTTP 500 Internal Server Error with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "INTERNAL_SERVER_ERROR",
            "message": "An unexpected error occurred during login."
          }
        }
        ```
        

---

## 1. Function Specification: Get User Profile

- **Function Name**: GetProfile
    
- **Description**: Retrieves the profile information for the authenticated user.
    
- **Key Business Logic**:
    
    - Authenticate the request using a provided token (e.g., JWT in Authorization header).
        
    - Retrieve user details from the database based on the authenticated user ID.
        
    - Return user profile information, including subscription tier.
        
- **Validation Rules**:
    
    - Request must be authenticated.
        

### 2. API Endpoint

#### HTTP Method and Path

```
GET /api/user/profile
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json
```

#### Request Body (for POST/PUT)

N/A (GET request)

#### Response Specification

**Web Response (HTML)**

HTML

HTML

```
<div id="user-profile">
    <h2 class="text-xl font-bold">Welcome, {{.Email}}!</h2>
    <p>Subscription Tier: <span class="font-semibold">{{.SubscriptionTier}}</span></p>
    </div>
```

**Mobile Response (JSON)**

JSON

JSON

```
{
  "success": true,
  "data": {
    "user_id": "uuid",
    "email": "user@example.com",
    "subscription_tier": "premium",
    "created_at": "2025-07-06T12:00:00Z"
  },
  "message": "User profile retrieved successfully"
}
```

### 3. Go Handler Structure

Go

Go

```
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
    // This handler assumes an authentication middleware has already run
    // and set the user ID or user object in the request context.
    // For simplicity, we'll simulate fetching a user ID from context.

    // ctxUser, ok := r.Context().Value("user").(*models.User)
    // if !ok {
    //     http.Error(w, "Unauthorized", http.StatusUnauthorized)
    //     return
    // }

    // Simulate a user being authenticated
    userID := "simulated-user-uuid-123" // Replace with actual user from context

    // Fetch profile from database (via gameService or a dedicated userService)
    profile, err := h.gameService.GetUserProfile(userID) // This would interact with Supabase
    if err != nil {
        if isJSONRequest(r) {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   map[string]string{"code": "PROFILE_RETRIEVAL_FAILED", "message": "Failed to retrieve profile: " + err.Error()},
            })
        } else {
            fmt.Fprintf(w, `<div id="user-profile-error" class="text-red-500"><p>Failed to load profile: %s</p></div>`, err.Error())
        }
        return
    }

    if isJSONRequest(r) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": true,
            "data":    profile,
            "message": "User profile retrieved successfully",
        })
    } else {
        w.Header().Set("Content-Type", "text/html")
        // Assuming a Templ component `UserProfile`
        // userProfileComponent := UserProfile(profile)
        // userProfileComponent.Render(r.Context(), w)
        fmt.Fprintf(w, `<div id="user-profile"><h2 class="text-xl font-bold">Welcome, %s!</h2><p>Subscription Tier: <span class="font-semibold">%s</span></p></div>`, profile.Email, profile.SubscriptionTier)
    }
}
```

### 4. Error Handling

- **Unauthorized**:
    
    - **Web Response (HTML)**: Returns HTTP 401 Unauthorized with a plain text error message, or redirects to login page.
        
    - **Mobile Response (JSON)**: Returns HTTP 401 Unauthorized with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "UNAUTHORIZED",
            "message": "Authentication required or invalid token."
          }
        }
        ```
        
- **User Not Found (after authentication)**:
    
    - **Web Response (HTML)**: Returns HTTP 404 Not Found with a plain text error message "User profile not found".
        
    - **Mobile Response (JSON)**: Returns HTTP 404 Not Found with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "USER_PROFILE_NOT_FOUND",
            "message": "User profile could not be found."
          }
        }
        ```
        
- **Internal Server Error**:
    
    - **Web Response (HTML)**: Returns HTTP 500 Internal Server Error with a plain text error message "Internal server error".
        
    - **Mobile Response (JSON)**: Returns HTTP 500 Internal Server Error with JSON:
        
        JSON
        
        ```
        {
          "success": false,
          "error": {
            "code": "INTERNAL_SERVER_ERROR",
            "message": "An unexpected error occurred while retrieving profile."
          }
        }
        ```
        

---

## 1. Function Specification: Get Report Card Partial

- **Function Name**: GetReportCardPartial
    
- **Description**: Retrieves only the HTML partial for a daily report card, specifically designed for HTMX requests to refresh or display report content within a larger page.
    
- **Key Business Logic**:
    
    - Extract `summonerName` and `tagLine` from the URL.
        
    - Call `h.gameService.GetDailyReport()` to get the latest report data. This will leverage caching if a recent report exists.
        
    - Render _only_ the `report-card.html` template partial with the retrieved data.
        
- **Validation Rules**:
    
    - `summonerName` and `tagLine` must be provided in the URL.
        

### 2. API Endpoint

#### HTTP Method and Path

```
GET /partials/report-card/{summonerName}/{tagLine}
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

(This endpoint is explicitly for HTML partials, so `Accept: application/json` is not expected for its primary use case.)

**Mobile Request (Flutter)**

N/A (This endpoint is not intended for direct mobile consumption of raw JSON data; mobile apps should use the `/api/report/{summonerName}/{tagLine}` endpoint for JSON.)

#### Request Body (for POST/PUT)

N/A (GET request)

#### Response Specification

**Web Response (HTML)**

HTML

HTML

```
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
```

**Mobile Response (JSON)**

N/A (This endpoint strictly returns HTML for HTMX usage.)

### 3. Go Handler Structure

Go

Go

```
func (h *Handler) GetReportCardPartial(w http.ResponseWriter, r *http.Request) {
    summonerName := chi.URLParam(r, "summonerName")
    tagLine := chi.URLParam(r, "tagLine")

    report, err := h.gameService.GetDailyReport(summonerName, tagLine)
    if err != nil {
        // h.renderError(w, "Failed to generate report: " + err.Error()) // Render a simple HTML error partial
        fmt.Fprintf(w, `<div class="text-red-500"><p>Failed to retrieve report: %s</p></div>`, err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // This handler specifically renders the HTML partial, so no JSON branching needed
    // h.renderReportCard(w, report) // Assuming renderReportCard exists and renders the Templ component
    data := struct {
        Player   services.Player
        Date     string
        WinRate  float64
        AvgKDA   float64
        Insights services.AIInsight
    }{
        Player:   report.Player,
        Date:     time.Now().Format("January 2, 2006"),
        WinRate:  report.WinRate * 100,
        AvgKDA:   report.AvgKDA,
        Insights: report.Insights,
    }
    w.Header().Set("Content-Type", "text/html")
    // templ.ExecuteTemplate(w, "report-card", data) or TemplComponent.Render(r.Context(), w)
    fmt.Fprint(w, `<div class="bg-white rounded-lg shadow-lg p-6 mt-6">...rendered report card HTML...</div>`) // Placeholder
}
```

### 4. Error Handling

- **Missing Parameters**:
    
    - **Web Response (HTML)**: Returns HTTP 400 Bad Request with a simple HTML error message "Missing summoner name or tag line".
        
    - **Mobile Response (JSON)**: N/A (not applicable to this endpoint)
        
- **Failed to Retrieve Report (e.g., GameService error, Riot API error)**:
    
    - **Web Response (HTML)**: Returns HTTP 500 Internal Server Error with a simple HTML error message "Failed to generate report". This will update the `hx-target` with the error message.
        
    - **Mobile Response (JSON)**: N/A (not applicable to this endpoint)
        
- **Player Data Not Found**:
    
    - **Web Response (HTML)**: Returns HTTP 404 Not Found with a simple HTML error message "Player data not found or no recent matches".
        
    - **Mobile Response (JSON)**: N/A (not applicable to this endpoint)
        

---

## 1. Function Specification: Get Loading Partial

- **Function Name**: GetLoadingPartial
    
- **Description**: Provides a simple HTML snippet for a loading indicator, primarily used by HTMX during asynchronous operations.
    
- **Key Business Logic**:
    
    - Simply returns a predefined HTML string or renders a simple Templ component that displays a loading animation and message.
        
- **Validation Rules**: N/A
    

### 2. API Endpoint

#### HTTP Method and Path

```
GET /partials/loading
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

(This endpoint is explicitly for HTML partials.)

**Mobile Request (Flutter)**

N/A (This endpoint is not relevant for mobile apps, which would typically use a native loading indicator.)

#### Request Body (for POST/PUT)

N/A (GET request)

#### Response Specification

**Web Response (HTML)**

HTML

HTML

```
<div id="loading" class="htmx-indicator text-center mt-4">
    <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    <p class="text-gray-600 mt-2">Analyzing your performance...</p>
</div>
```

**Mobile Response (JSON)**

N/A

### 3. Go Handler Structure

Go

Go

```
func (h *Handler) GetLoadingPartial(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    // Assuming a Templ component `LoadingIndicator`
    // loadingComponent := LoadingIndicator()
    // loadingComponent.Render(r.Context(), w)
    fmt.Fprint(w, `<div id="loading" class="htmx-indicator text-center mt-4"><div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div><p class="text-gray-600 mt-2">Analyzing your performance...</p></div>`)
}
```

### 4. Error Handling

N/A (This endpoint should not typically encounter errors in its simple function.)

---

## 1. Function Specification: Health Check

- **Function Name**: Health
    
- **Description**: Provides a basic health status check for the Go web server.
    
- **Key Business Logic**:
    
    - Return a fixed success status, current timestamp, and application version.
        
    - Does not involve complex logic or database interaction directly.
        
- **Validation Rules**: N/A
    

### 2. API Endpoint

#### HTTP Method and Path

```
GET /health
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded
```

(While this endpoint is typically consumed by monitoring systems expecting JSON, it could theoretically return plain text or a simple HTML status for a browser.)

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json
```

#### Request Body (for POST/PUT)

N/A (GET request)

#### Response Specification

**Web Response (HTML)**

HTML

HTML

```
<p>Status: Healthy</p>
<p>Timestamp: 2025-07-06T12:00:00Z</p>
<p>Version: 1.0.0</p>
```

(Simplified for HTML, usually not the primary use case)

**Mobile Response (JSON)**

JSON

JSON

```
{
  "status": "healthy",
  "timestamp": "2025-07-06T12:00:00Z",
  "version": "1.0.0"
}
```

### 3. Go Handler Structure

Go

Go

```
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") // Health check commonly returns JSON
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":    "healthy",
        "timestamp": time.Now().Format(time.RFC3339),
        "version":   os.Getenv("APP_VERSION"), // Get app version from environment variable
    })
}
```

### 4. Error Handling

N/A (This endpoint should be designed to always return a response, even in error states, to indicate the health status.)