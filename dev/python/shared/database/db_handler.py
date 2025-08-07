"""db_handler.py"""

import os
from datetime import datetime, timezone
from supabase import create_client, Client
from dotenv import load_dotenv

load_dotenv()


class DBHandler:
    """
    Handles all interactions with the Supabase database.
    """

    def __init__(self, url: str, service_role_key : str):
        try:
            if not url or not service_role_key:
                raise ValueError(
                    "SUPABASE_URL and SUPABASE_SERVICE_ROLE_KEY must be set in environment variables."
                )
            self.client: Client = create_client(url, service_role_key)
            print("✅ Supabase client initialized")
        except Exception as e:
            print(f"❌ Supabase client initialization failed: {e}")

    def get_client(self):
        """Returns the Supabase client if initialized, otherwise None."""
        return self.client

    def bulk_insert_matches(self, matches: list):
        """
        Inserts a list of matches into the database in bulk.
        """
        try:
            self.client.table("matches").insert(matches).execute()
            print(f"✅ {len(matches)} matches inserted successfully")
        except Exception as e:
            print(f"❌ Failed to insert matches: {e}")

    def get_last_match(self, puuid: str) -> tuple[str, int] | None:
        """
        get the last match fetched in DB for a given puuid
        to optimize data fetching and avoid duplicates
        """
        try:
            response = (
                self.client.table("matches")
                .select("match_id, game_creation")
                .eq("user_puuid", puuid)
                .order("game_creation", desc=True)
                .limit(1)
                .execute()
            )

            if not response.data:
                print(f"❌ No matches found for puuid: {puuid}")
                return None

            last_match_data = response.data[0]
            match_id = last_match_data["match_id"]
            game_creation_dt = datetime.fromisoformat(last_match_data["game_creation"])
            game_creation_timestamp = int(game_creation_dt.timestamp())

            return (match_id, game_creation_timestamp)

        except Exception as e:
            print(f"❌ Failed to get last match: {e}")
            return None


if __name__ == "__main__":
    url =os.getenv("SUPABASE_URL")
    key =os.getenv("SUPABASE_SERVICE_ROLE_KEY")
    if not url or not key:
        raise ValueError("SUPABASE_URL and SUPABASE_SERVICE_ROLE_KEY must be set in environment variables.")
    
    db_handler = DBHandler(url, key)
    if db_handler.get_client():
        print("Supabase client test successful.")
    else:
        print("Supabase client test failed.")
