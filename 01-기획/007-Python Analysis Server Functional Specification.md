# Python Analysis Server Functional Specification

### 1. Overview

The Python Analysis Server is a specialized microservice responsible for all data collection, processing, statistical analysis, and AI-based report generation for the "Daily Gaming Report Card" project. It communicates with the Go Web Server via a REST API. This server is designed with a "Strategy Pattern" to easily extend support for multiple games.

### 2. Core Features

- **Data Collection**: Gathers raw match data from Riot API (and potentially other game APIs).
    
- **Data Processing**: Parses raw JSON data, extracts key performance indicators (KPIs), and calculates match-specific statistics for individual participants.
    
- **Data Aggregation**: Computes aggregated statistics (e.g., champion-specific, position-specific, tier-specific averages) from processed match data.
    
- **Database Storage**: Stores processed and aggregated data into a PostgreSQL database.
    
- **AI Insight Generation**: Utilizes the Google AI API to generate personalized insights based on analyzed player performance.
    
- **Game-Specific Analysis**: Employs a Strategy Pattern to dynamically select and execute game-specific analysis logic (e.g., for League of Legends, Valorant).
    
- **Automated Scheduling**: Designed to be run by a scheduler (Cron, Supabase Cron Jobs, GitHub Actions) for daily data pipeline execution.
    

### 3. Key Functions / Modules

#### 3.1. `main.py`

- **Location**: `main.py`
    
- **Description**: Entry point for the FastAPI application. Initializes the FastAPI app and includes the API routes from `api/v1/reports.py`.
    

#### 3.2. `reports.py` (API Endpoint Definition)

- **Location**: `app/api/v1/reports.py`
    
- **Description**: Defines the `/api/v1/reports/{game_name}` endpoint.
    
    - Receives the `game_name` and `ReportRequest`.
        
    - Uses `AnalysisFactory.get_analyzer()` to get the correct analyzer for the game.
        
    - Calls the `generate_report` method on the obtained analyzer.
        
    - Handles `ValueError` for unsupported games by returning a 404 HTTP exception.
        

#### 3.3. `AnalysisFactory`

- **Location**: `app/services/analysis_factory.py`
    
- **Description**: Implements the Strategy Pattern to provide the correct game analyzer.
    
    - `ANALYZERS` (dictionary): Maps game names (e.g., "lol", "valorant") to their respective analyzer classes.
        
    - `get_analyzer(game_name: str)` (static method): Returns an instance of the appropriate analyzer class based on `game_name`. Raises a `ValueError` if the game is not supported.
        

#### 3.4. `BaseAnalyzer` (Abstract Base Class)

- **Location**: `app/services/base_analyzer.py`
    
- **Description**: Defines the interface for all game-specific analyzers. All concrete analyzer implementations (e.g., `LolAnalyzer`, `ValorantAnalyzer`) must inherit from this class and implement its abstract methods.
    
    - `generate_report(summoner_name: str)` (abstract method): Signature for the main report generation logic for a specific game.
        

#### 3.5. `LolAnalyzer`

- **Location**: `app/services/analyzers/lol_analyzer.py`
    
- **Description**: Concrete implementation of `BaseAnalyzer` for League of Legends. Contains the specific logic for collecting, processing, and analyzing League of Legends data.
    
    - `generate_report(summoner_name: str)`: Implements the detailed steps for LoL data.
        

#### 3.6. `get_high_elo_summoner_ids()`

- **Location**: `collector/riot_api.py`
    
- **Description**: Calls Riot `LEAGUE-V4` API endpoints (`challengerleagues`, `grandmasterleagues`, `masterleagues`) to retrieve lists of summoner IDs from the highest tiers.
    
- **Returns**: A list of summoner IDs.
    

#### 3.7. `get_puuid_by_summoner_id(summoner_id)`

- **Location**: `collector/riot_api.py`
    
- **Description**: Calls Riot `SUMMONER-V4` API to convert an `encryptedSummonerId` into a `PUUID` (and `accountId`).
    
- **Parameters**: `summoner_id`.
    
- **Returns**: `puuid`, `accountId`, etc.
    

#### 3.8. `get_match_ids_by_puuid(puuid)`

- **Location**: `collector/riot_api.py`
    
- **Description**: Calls Riot `MATCH-V5` API (`matches/by-puuid/{puuid}/ids`) to retrieve a list of recent `matchId`s for a given `puuid`.
    
- **Parameters**: `puuid`.
    
- **Returns**: A list of `matchId`s.
    

#### 3.9. `get_match_data_by_match_id(match_id)`

- **Location**: `collector/riot_api.py`
    
- **Description**: Calls Riot `MATCH-V5` API (`matches/{matchId}`) to retrieve detailed match data (MatchDTO) for a specific `matchId`.
    
- **Parameters**: `match_id`.
    
- **Returns**: Detailed match data as JSON (`MatchDTO`).
    

#### 3.10. `process_all_participants_in_match(match_json)`

- **Location**: `processor/process_match.py`
    
- **Description**: Processes a raw `match_json` (MatchDTO) to extract key statistics for _all 10 participants_ in the match.
    
    - Iterates through the `participants` list within `match_json`.
        
    - Calculates KDA, CS per minute, damage per minute, and other KPIs for each participant.
        
    - Formats the extracted data into a list of dictionaries suitable for insertion into the `matches` table.
        
- **Parameters**: `match_json` (raw JSON from `Match API`).
    
- **Returns**: A list of dictionaries (10 entries), each representing a participant's stats in the match.
    

#### 3.11. `update_aggregated_stats()`

- **Location**: `processor/aggregate_stats.py`
    
- **Description**: Reads data from the `matches` table, uses `Pandas` to group and aggregate statistics (`game_version`, `champion_id`, `position` wise), and updates the `aggregated_stats` table. This function is designed to be run periodically by the scheduler.
    

#### 3.12. `bulk_insert_matches(match_data_list)`

- **Location**: `database/db_handler.py`
    
- **Description**: Efficiently inserts multiple match records (from `process_all_participants_in_match`) into the `matches` table in a single database operation to reduce overhead.
    
- **Parameters**: `match_data_list` (list of dictionaries).
    

#### 3.13. `get_users_to_crawl(limit)`

- **Location**: `database/db_handler.py`
    
- **Description**: Retrieves a specified number of users from the `tracked_users` table whose match history needs to be updated (i.e., `last_crawled_at` is the oldest).
    
- **Parameters**: `limit` (int).
    
- **Returns**: A list of user objects.
    

#### 3.14. `add_new_puuids_to_tracked_users(puuid_list)`

- **Location**: `database/db_handler.py`
    
- **Description**: Adds newly discovered `puuid`s (from match participants) to the `tracked_users` table. Prevents duplicates by ignoring `puuid`s that already exist. This implements the "Snowball" data expansion strategy.
    
- **Parameters**: `puuid_list` (list of strings).
    

#### 3.15. `run_daily_pipeline()`

- **Location**: `main.py` (or `run_daily_pipeline.py`)
    
- **Description**: The main orchestrator function for the daily data pipeline.
    
    1. Gets users to crawl via `db_handler.get_users_to_crawl()`.
        
    2. For each user, retrieves match IDs using `riot_api.get_match_ids_by_puuid()`.
        
    3. Filters out already stored `match_id`s.
        
    4. Fetches new match data using `riot_api.get_match_data_by_match_id()`.
        
    5. Processes all 10 participants' data using `processor.process_all_participants_in_match()`.
        
    6. Bulk inserts these 10 records into the `matches` table using `db_handler.bulk_insert_matches()`.
        
    7. Adds new `puuid`s found in the match participants to `tracked_users` using `db_handler.add_new_puuids_to_tracked_users()`.
        
    8. After processing all matches, updates the `aggregated_stats` using `processor.update_aggregated_stats()`.