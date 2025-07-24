## Proposed Order of Implementation

### Elixir/Phoenix Web Server Functional Specifications

1.  **Project Setup & Health Check**
    -   **Justification**: Fundamental infrastructure. Generate the Phoenix project, set up the database connection (`Repo`), and create a `/health` route. This ensures the server is running and deployable from day one.

2.  **Database Schema (Ecto)**
    -   **Justification**: Define the Ecto schemas for `players`, `daily_reports`, and `users`. Create the initial database migrations. A solid data model is the foundation for all other features.

3.  **Player Search & Riot API Client**
    -   **Justification**: This is the entry point for any user interaction. Implement the `RiotAPI` client module to fetch player data. Create a `PlayerController` or `PlayerLive` to handle the initial search and display basic player info.

4.  **HomeLive & PlayerLive (LiveView)**
    -   **Justification**: Build the core user interface. The `HomeLive` view will contain the search form. Upon a successful search, it will redirect to the `PlayerLive` view, which will be the main dashboard for displaying the report.

5.  **Daily Report Generation (Analytics Context)**
    -   **Justification**: This is the core feature. Create an `Analytics` context that orchestrates fetching data from the Riot API, calling the Python analysis server, and saving the results to the database. This is the most complex piece of the Elixir application.

6.  **Display Report Card in LiveView**
    -   **Justification**: Integrate the data from the `Analytics` context into the `PlayerLive` view. This involves rendering the statistics, AI insights, and charts.

7.  **User Registration & Login (Accounts Context)**
    -   **Justification**: User management is essential for premium features. Create an `Accounts` context to handle user registration and authentication using `phx.gen.auth` or a similar library.

8.  **Premium Features & Subscription Logic**
    -   **Justification**: Once the core free features are stable and user accounts exist, implement the logic to restrict access to premium reports and manage subscription status.

### Python/Rust Analysis Server Functional Specifications

**Note:** The following Python functions represent the initial implementation for the **Phase 1 MVP**. The core aggregation logic (e.g., `update_aggregated_stats` and parts of `run_daily_pipeline`) is designed to be replaced by a more robust, high-performance **Rust batch processing tool in Phase 2** to ensure long-term scalability and efficiency.

1.  **`main.py` (FastAPI Application Entry Point)**
    -   **Justification**: This is the very first piece of the Python server that needs to be functional. It sets up the FastAPI application and defines how other modules are integrated.

2.  **`AnalysisFactory` & `BaseAnalyzer`**
    -   **Justification**: These establish the core architectural pattern (Strategy Pattern) for handling different games. Implementing them early ensures that the system is designed for extensibility from the start.

3.  **Riot API Client Functions**
    -   **Justification**: Fundamental for data acquisition. Without being able to fetch raw data from Riot, no analysis can be performed.

4.  **Database Handler Functions**
    -   **Justification**: Core database interaction is critical for storing raw data and retrieving user information for the daily processing pipeline.

5.  **`LolAnalyzer` (Initial Implementation)**
    -   **Justification**: The first concrete implementation of a game analyzer, integrating the Riot API client to fetch and process LoL data.

6.  **`process_all_participants_in_match(match_json)`**
    -   **Justification**: This function processes the raw match data into a usable format for analysis and storage.

7.  **`run_daily_pipeline()`**
    -   **Justification**: This orchestrates the entire data collection and processing flow. It will eventually be triggered to call the Rust engine.

8.  **`update_aggregated_stats()` (to be replaced by Rust)**
    -   **Justification**: This function computes the higher-level aggregated statistics. It's implemented in Python for the MVP but is the primary candidate for migration to Rust for performance.

9.  **`create_daily_report` (API Endpoint)**
    -   **Justification**: This is the primary API endpoint that the Elixir server calls. It needs all underlying analysis, data processing, and AI generation logic to be in place.