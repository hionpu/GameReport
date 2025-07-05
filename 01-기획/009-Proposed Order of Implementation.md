## Proposed Order of Implementation

### Go Web Server Functional Specifications

1. **Health Check**
    
    - **Justification**: This is a fundamental infrastructure piece. Implementing it first ensures the server is running and accessible from the very beginning. It's crucial for deployment and monitoring.
        
2. **Player Search and Basic Info**
    
    - **Justification**: This is the first step in the user's journey for generating a report. It's a relatively simple API call that doesn't involve complex analysis, but it's a necessary precursor to the main report generation. It also allows early testing of basic communication between the frontend (web/mobile) and the Go server.
        
3. **Get Loading Partial**
    
    - **Justification**: While small, this UI component is essential for providing immediate user feedback during longer operations (like report generation). Implementing it early improves the perceived responsiveness of the application. It has no complex dependencies.
        
4. **Daily Report Generation**
    
    - **Justification**: This is a core feature that relies on the "Player Search" to get player info and the Python server to perform the actual analysis. It's complex as it involves branching responses (HTML/JSON) and integrating with the Python backend. It should come after basic infrastructure and initial player lookup.
        
5. **Get Report Card Partial**
    
    - **Justification**: This is a direct dependency of the "Daily Report Generation" for the HTMX flow. It reuses much of the logic from report generation but focuses solely on rendering the HTML partial. It should be implemented concurrently or immediately after the full report generation to complete the HTMX user experience.
        
6. **User Registration**
    
    - **Justification**: User management features are important for premium tiers but are not part of the initial MVP's core "daily report card" value proposition. They can be implemented after the primary analysis flow is stable.
        
7. **User Login**
    
    - **Justification**: Depends on User Registration.
        
8. **Get User Profile**
    
    - **Justification**: Depends on User Login and Registration. It's a premium feature.
        

### Python Analysis Server Functional Specifications

1. **`main.py` (FastAPI Application Entry Point)**
    
    - **Justification**: This is the very first piece of the Python server that needs to be functional. It sets up the FastAPI application and defines how other modules are integrated.
        
2. **`AnalysisFactory` & `BaseAnalyzer`**
    
    - **Justification**: These establish the core architectural pattern (Strategy Pattern) for handling different games. Implementing them early ensures that the system is designed for extensibility from the start, even if only one game is supported initially.
        
3. **Riot API Client Functions (`get_high_elo_summoner_ids()`, `get_puuid_by_summoner_id()`, `get_match_ids_by_puuid()`, `get_match_data_by_match_id()`)**
    
    - **Justification**: These are fundamental for data acquisition. Without being able to fetch raw data from Riot, no analysis can be performed. These should be implemented and thoroughly tested early.
        
4. **Database Handler Functions (`bulk_insert_matches()`, `get_users_to_crawl()`, `add_new_puuids_to_tracked_users()`)**
    
    - **Justification**: Core database interaction is critical. `bulk_insert_matches` allows efficient storage of the collected raw data. `get_users_to_crawl` and `add_new_puuids_to_tracked_users` are essential for building and maintaining the pool of users whose matches will be analyzed daily (the "Snowball" effect).
        
5. **`LolAnalyzer` (Initial Implementation)**
    
    - **Justification**: This is the first concrete implementation of a game analyzer. It will integrate the Riot API client functions to perform the game-specific data fetching and initial processing. This can be built concurrently with `process_all_participants_in_match` but needs to define the structure of the data it expects from the Riot API and produces for the processor.
        
6. **`process_all_participants_in_match(match_json)`**
    
    - **Justification**: This function processes the raw match data into a usable format for analysis and storage. It is a direct dependency for populating the `matches` table and for generating the `aggregated_stats`.
        
7. **`run_daily_pipeline()`**
    
    - **Justification**: This orchestrates the entire data collection and processing flow. It pulls together the Riot API calls, data processing, and database interactions. It should be developed after its constituent parts are ready.
        
8. **`update_aggregated_stats()`**
    
    - **Justification**: This function takes the processed raw match data and computes the higher-level aggregated statistics. This data is what the Go server will ultimately request for AI insights. It depends on `matches` table being populated.
        
9. **`create_daily_report` (API Endpoint Definition in `reports.py`)**
    
    - **Justification**: This is the primary API endpoint that the Go server calls. It needs all underlying analysis, data processing, and AI generation logic to be in place. It will be the final step to expose the analysis capabilities.