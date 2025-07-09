
# Phase 1: Foundational Data Pipeline - Detailed Plan (v3.0)

## 1. Overview 🎯

**Goal:** To build an automated pipeline that collects both match results and key timeline events from the Riot API. The data will be processed and stored in a structured database, creating a robust and scalable foundation for all future analysis features in Phases 2 and 3.

---

## 2. Core Data Collection Strategy

- **1 Match, 10-Player Data:** Maximize API efficiency by processing and storing the performance data for all 10 participants from a single match fetch. One API call yields 10 rows of performance data.
- **Proactive Timeline Data Collection:** Fetch both the `MatchDTO` (end-game results) and the `MatchTimelineDTO` (event logs) from the very beginning. This prevents the need to re-fetch data for old matches when developing future features.
- **Organic User Pool Expansion (Snowball Effect):** Use the participants from collected matches to discover new users, organically growing the list of players to track.

---

## 3. Riot API Endpoints to Use

- **A. Acquiring Seed Users:** `LEAGUE-V4`
  - `lol/league/v4/challengerleagues/by-queue/{queue}`
  - `lol/league/v4/grandmasterleagues/by-queue/{queue}`
  - `lol/league/v4/masterleagues/by-queue/{queue}`
- **B. Converting User Identifiers:** `SUMMONER-V4`
  - `lol/summoner/v4/summoners/{encryptedSummonerId}` -> to get `puuid`.
- **C. Collecting Match ID Lists:** `MATCH-V5`
  - `lol/match/v5/matches/by-puuid/{puuid}/ids`
- **D. Collecting Detailed Match Data:** `MATCH-V5`
  - `lol/match/v5/matches/{matchId}` -> `MatchDTO`
  - `lol/match/v5/matches/{matchId}/timeline` -> `MatchTimelineDTO`

**※ Important:** All API calls must respect Riot's rate limits. Implementing appropriate delays (e.g., `time.sleep()`) between requests is mandatory.

---

## 4. Database Table Design (PostgreSQL)

This design includes storing key events from the timeline, adopting a balanced approach between data completeness and initial complexity.

### `matches` (Individual Match Performance)
Stores the final end-game stats for each player in a match. (10 rows generated per match).

| Column Name            | Data Type      | Description                                          |
| :--------------------- | :------------- | :--------------------------------------------------- |
| `match_id`             | `VARCHAR(32)`  | Composite PK, Match unique ID                        |
| `user_puuid`           | `VARCHAR(128)` | Composite PK, Player's unique PUUID                  |
| `game_version`         | `VARCHAR(32)`  | Game version (e.g., "14.1.555")                      |
| `game_duration`        | `INT`          | Game length in seconds                               |
| `champion_id`          | `INT`          | Champion played                                      |
| `champion_name`        | `VARCHAR(32)`  | Champion name (for readability)                      |
| `team_position`        | `VARCHAR(16)`  | Player's role (TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY) |
| `win`                  | `BOOLEAN`      | True if the player's team won                        |
| `kills`                | `INT`          | Number of kills                                      |
| `deaths`               | `INT`          | Number of deaths                                     |
| `assists`              | `INT`          | Number of assists                                    |
| `kda`                  | `FLOAT`        | Pre-calculated KDA ratio                             |
| `cs_per_min`           | `FLOAT`        | Minions + Jungle monsters killed per minute          |
| `damage_per_min`       | `FLOAT`        | Damage dealt to champions per minute                 |
| `gold_per_min`         | `FLOAT`        | Gold earned per minute                               |
| `vision_score_per_min` | `FLOAT`        | Vision score per minute                              |
| `control_wards_placed` | `INT`          | Number of control wards placed                       |
| `created_at`           | `TIMESTAMP`    | Timestamp when the record was created                |

### `match_events` (Key Timeline Events)
A single, flexible table to store the most important events from the timeline, ready for Phase 2 analysis.

| Column Name | Data Type | Description |
| :--- | :--- | :--- |
| `event_id` | `SERIAL` | PK, Auto-incrementing ID |
| `match_id` | `VARCHAR(32)` | FK, links to the match |
| `timestamp` | `INT` | Event timestamp in milliseconds |
| **`event_type`** | **`VARCHAR(32)`**| Type of event (e.g., `CHAMPION_KILL`, `ELITE_MONSTER_KILL`, `BUILDING_KILL`, `ITEM_PURCHASED`, `SKILL_LEVEL_UP`) |
| `participant_id`| `INT` | The participant (1-10) who is the main actor |
| `killer_id` | `INT` | (For kill events) Participant ID of the killer |
| `victim_id` | `INT` | (For kill events) Participant ID of the victim |
| `item_id` | `INT` | (For item events) The ID of the item |
| `skill_slot` | `INT` | (For skill events) The skill slot leveled up (1-4) |
| `monster_type` | `VARCHAR(32)` | (For objective events) e.g., `DRAGON`, `BARON_NASHOR` |
| `building_type`| `VARCHAR(32)` | (For building events) e.g., `TOWER_BUILDING`|
| `position_x` | `INT` | X-coordinate of the event location |
| `position_y` | `INT` | Y-coordinate of the event location |

*(Note: Columns not relevant to a specific `event_type` will contain `NULL`)*

### `aggregated_stats` (Aggregated Statistics)
This table is periodically updated from the `matches` table to serve data to the client API.

| Column Name | Data Type | Description |
| :--- | :--- | :--- |
| `id` | `SERIAL` | PK, Auto-incrementing ID |
| `game_version` | `VARCHAR(32)` | Game version for the stats |
| `tier_group` | `VARCHAR(16)` | Tier group (e.g., HIGH_ELO, MID_ELO) |
| `champion_id` | `INT` | Champion for the stats |
| `team_position` | `VARCHAR(16)` | Role for the stats |
| `total_games` | `INT` | Number of games analyzed |
| `win_rate` | `FLOAT` | Win rate |
| `pick_rate` | `FLOAT` | Pick rate |
| `avg_kda` | `FLOAT` | Average KDA |
| `avg_cs_per_min`| `FLOAT` | Average CS per minute |
| `avg_damage_per_min`|`FLOAT` | Average Damage to Champions per minute |
| `updated_at` | `TIMESTAMP` | When this row was last updated |

---

## 5. Python Script Structure & Key Functions

**Note:** The following Python-based structure is for the initial **Phase 1 MVP** to ensure a rapid launch. The core data processing and aggregation logic will be migrated to a high-performance **C++ engine in Phase 2** to ensure scalability.

A modular structure is recommended for maintainability.


```

lol_pipeline/ 
├── main.py       # Main script to orchestrate the entire pipeline 
├── api/          # FastAPI server code 
│ └── server.py 
├── collector/    # Module for fetching data from Riot API 
│ └── riot_api.py
├── processor/    # Module for processing and transforming data 
│ └── data_processor.py 
├── database/     # Module for all database interactions
│ └── db_handler.py 
└── config.py     # API keys, database credentials, etc.

```

### Key Function Specifications

#### `collector/riot_api.py`
- `fetch_match_data(match_id)`: Fetches the `MatchDTO`.
- `fetch_timeline_data(match_id)`: Fetches the `MatchTimelineDTO`.
- (Other functions for getting user lists and match histories).

#### `processor/data_processor.py`
- **`process_match_and_timeline(match_json, timeline_json)`**:
    1.  Parses `match_json` to generate a **list of 10 dictionaries** for the `matches` table.
    2.  Parses `timeline_json` to generate a **list of dictionaries** for the `match_events` table (only for the core event types).
    3.  Returns a dictionary containing these lists, e.g., `{"matches": [...], "events": [...]}`.

#### `database/db_handler.py`
- **`bulk_insert_matches(match_data_list)`**: Inserts a list of match performance records into the `matches` table.
- **`bulk_insert_events(event_data_list)`**: Inserts a list of event records into the `match_events` table.
- (Other functions for managing `tracked_users` and querying data).

#### `main.py`
- **`run_daily_pipeline()`**:
    1.  Get a batch of users to process from the database.
    2.  For each user, get their recent match history.
    3.  For each new `match_id`, fetch both the `MatchDTO` and `MatchTimelineDTO`.
    4.  Call `processor.process_match_and_timeline()` to process all data.
    5.  Call the `db_handler` functions to bulk-insert the processed data into the `matches` and `match_events` tables.
    6.  Add any new players discovered to the `tracked_users` list.
    7.  After all matches are processed, run the aggregation script to update the `aggregated_stats` table.

---

## 6. API Endpoint Specification

A basic API to serve the aggregated data to the client application.

- **Endpoint:** `GET /stats/champion/{champion_id}`
- **Description:** Retrieves aggregated stats for a specific champion.
- **Query Parameters:**
    - `position` (string, optional): `TOP`, `JUNGLE`, etc.
    - `tier_group` (string, optional): `HIGH_ELO`, `MID_ELO`, etc.
- **Success Response (200 OK):**
    ```json
    {
      "champion_id": 103,
      "position": "MIDDLE",
      "tier_group": "HIGH_ELO",
      "total_games": 1250,
      "win_rate": 0.52,
      "pick_rate": 0.08,
      "avg_kda": 4.1,
      "avg_cs_per_min": 8.5
    }
    ```
```