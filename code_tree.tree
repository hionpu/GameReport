# Code Tree Structure

Generated from ctags analysis of the codebase.

**Filtered for subpath:** `dev`

## Symbol Legend

| Symbol | Type | Description |
|--------|------|-------------|
| D | Folder | Directory |
| F | File/Header | Source file or header file |
| f | Function/Macro | Standalone function or macro |
| m | Method | Class/struct method |
| v | Field/Variable | Struct field or variable |
| C | Constant | Constant value |
| c | Class/Namespace | Class, namespace, or wrapper |
| s | Struct | Struct definition |
| i | Interface/Heading | Interface or heading |
| t | Type/Typedef | Type definition or typedef |
| e | Enum | Enumeration |
| E | Enum Member | Enumeration member |
| a | Alias/Property | Type alias, property, or binding |
| l | Local Variable | Local variable |
| p | Parameter | Function parameter |
| o | Object | Object definition |
| x | External/Zone | External reference or zone |
| r | Reference | Reference |
| b | Block | Code block |
| k | Keyword/Key | Keyword, key, or JSON key |
| w | Widget | Widget |
| u | Union | Union type |
| A | Array | Array |
| . | CSS Class | CSS class selector |
| S | CSS Selector | CSS selector |
| R | CSS Rule | CSS rule |
| # | CSS/HTML ID | CSS or HTML ID selector |
| ! | Event | Event |
| g | Getter | Getter method |
| h | Heading | Heading |
| M | Macro | Macro definition |
| P | Property | Property |
| T | Typedef | Type definition |
| L | Label | Label |
| H | Header | Header file |
| n | Namespace | Namespace |
| d | Define | Define/macro |
| ? | Unknown | Unknown or unrecognized type |

**Compact Format:** `symbol:name{children}` where `->` indicates return type

D:dev{D:go{D:cmd{D:server{F:main.go{f:main}}} D:internal{D:config{F:config.go{s:Config{v:DatabaseURL->string v:Environment->string v:Port->string v:SupabaseAnonKey->string v:SupabaseServiceRoleKey->string v:SupabaseURL->string f:Validate->error} f:LoadConfig->*Config f:getEnv(key, fallback string)->string}} D:db{F:supabase.go{s:DB{M:DB->*sql.DB f:GetDatabaseStatistics->map[string]interface{} f:HealthCheck->error} s:DBConfig{v:ConnMaxLifetime->time.Duration v:MaxIdleConns->int v:MaxOpenConns->int v:URL->string} f:NewDB(config DBConfig)->*DB, error}} D:handlers{F:handlers.go{s:Handler{f:HealthCheck(w http.ResponseWriter, r *http.Request) f:Home(w http.ResponseWriter, r *http.Request) f:NotFound(w http.ResponseWriter, r *http.Request) v:config->*config.Config v:db->*db.DB v:template->*template.Template} f:NewHandler(cfg *config.Config, db *db.DB)->*Handler, error f:loadTemplates->*template.Template, error C:webFolderPath}} D:middleware{F:middleware.go{f:CORS->func(next http.Handler) http.Handler f:Logger->func(next http.Handler) http.Handler f:Recovery->func(next http.Handler) http.Handler f:Timeout(duration time.Duration)->func(next http.Handler) http.Handler}} D:models{F:models.go{s:APIResponse{v:Data->interface{} v:Error->string v:Message->string v:Success->bool} s:DatabaseHealth{v:Connections->map[string]interface{} v:Latency->string v:Status->string} s:HealthStatus{v:Database->DatabaseHealth v:Details->map[string]interface{} v:Status->string v:Supabase->SupabaseHealth v:Timestamp->time.Time v:Version->string} s:PageData{v:CurrentDate->string v:Data->map[string]interface{} v:Title->string v:Version->string} s:SupabaseHealth{v:Connected->bool v:Status->string v:URL->string}}}} D:web{D:static{D:css{F:custom.css{#:#mobile-menu c:.achievement-glow c:.animate-spin c:.bg-purple-accent c:.bg-purple-bg c:.bg-purple-primary c:.bg-purple-secondary c:.border-purple-100 c:.border-purple-200 c:.border-purple-primary c:.btn-loading c:.btn-loading::after c:.card-hover c:.card-hover:hover c:.custom-scrollbar c:.custom-scrollbar::-webkit-scrollbar c:.custom-scrollbar::-webkit-scrollbar-thumb c:.custom-scrollbar::-webkit-scrollbar-thumb:hover c:.custom-scrollbar::-webkit-scrollbar-track c:.error-message c:.focus c:.from-purple-50 c:.from-purple-600 c:.hover c:.htmx-indicator c:.htmx-request .htmx-indicator c:.htmx-request.htmx-indicator c:.htmx-settling c:.mobile-full c:.mobile-stack c:.mobile-text-center c:.progress-bar c:.success-message c:.text-purple-primary c:.text-purple-secondary c:.to-purple-100 c:.to-purple-700 c:.transition-colors s::bg-purple-600:hover s::ring-2:focus s::ring-purple-primary:focus s::text-purple-600:hover s::text-purple-primary:hover}}} D:templates{D:components{F:analysis-result.html{i:Analysis Results for {{.SummonerName}} #:improvement-tips #:match-details #:share-modal h:🎮 Recent Matches h:📊 Performance Insights} F:footer.html{#:back-to-top #:newsletter-result} F:header.html{#:mobile-menu} F:sample-report.html{h:Ready to get your personalized report? i:Sample Analysis Report h:🎮 Recent Match Analysis h:📊 Performance Overview}} F:home.html{h:AI Analyzes Your Data h:Achievement System h:Daily Performance Analysis h:Enter Your Summoner Name h:Free h:Get Your Daily Report i:How It Works i:Powerful Features for Better Gaming h:Premium i:Ready to Level Up Your Gaming? i:Simple Pricing h:Single Match Deep Dive h:Today's Report h:Your Daily Gaming Report Card #:analysis-result #:features #:how-it-works #:loading #:pricing}}}} D:python{D:analyzers{F:base.py{c:BaseAnalyzer(ABC){m:__init__(self, puuid: str) m:fetch_data(self) m:generate_insights(self, processed_data) m:get_user_report(self) m:process_data(self, raw_data)}} F:lol_analyzer.py{c:LolAnalyzer(BaseAnalyzer){m:__init__(self, api_key) m:fetch_data(self, puuid) m:generate_insights(self, processed_data) m:get_user_report(self) m:process_data(self, raw_data) m:save_data(self, processed_data)}}} D:database{F:db_handler.py{c:DBHandler{m:__init__(self) m:bulk_insert_matches(self, matches: list) m:get_client(self)} v:db_handler}} F:main.py{v:app f:health_check v:port(os.getenv("PORT", "8001")}}}
