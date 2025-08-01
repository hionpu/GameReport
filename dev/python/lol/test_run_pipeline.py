"""
League of Legends Data Pipeline Runner
"""
import os
import sys
from dotenv import load_dotenv

# Add project root to the Python path to allow module imports
project_root = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
sys.path.insert(0, project_root)

from lol.analyzers.lol_analyzer import LolAnalyzer
from shared.database.db_handler import DBHandler

def main():
    """
    Main function to run the League of Legends data analysis pipeline.
    """
    load_dotenv()

    api_key = os.getenv("RIOT_API_KEY")
    if not api_key:
        print("Error: RIOT_API_KEY not found in environment variables.")
        return

    # --- IMPORTANT ---
    # Replace this with a valid PUUID from the KR region for testing.
    # You can get a PUUID from the Riot API using a summoner name.
    target_puuid = "REPLACE_WITH_YOUR_TEST_PUUID"

    if target_puuid == "REPLACE_WITH_YOUR_TEST_PUUID":
        print("Error: Please replace the placeholder `target_puuid` in `dev/python/lol/run_pipeline.py`.")
        return

    print("Initializing database handler...")
    db_handler = DBHandler()

    print("Initializing LoL Analyzer...")
    lol_analyzer = LolAnalyzer(api_key=api_key, db_handler=db_handler)

    print(f"Running analysis for PUUID: {target_puuid}")
    lol_analyzer.fetch_data(identifier=target_puuid)
    print("LoL pipeline execution finished.")

if __name__ == "__main__":
    main()