import os
from supabase import create_client, Client
from dotenv import load_dotenv

load_dotenv()

class DBHandler:
    """
    Handles all interactions with the Supabase database.
    """
    def __init__(self):
        url: str = os.getenv("SUPABASE_URL")
        key: str = os.getenv("SUPABASE_SERVICE_ROLE_KEY")
        
        try:
            self.client: Client = create_client(url, key)
            print("✅ Supabase client initialized")
        except Exception as e:
            self.client = None
            print(f"❌ Supabase client initialization failed: {e}")

    def get_client(self):
        return self.client
    
    def bulk_insert_matches(self, matches: list):
        try:
            self.client.table("matches").insert(matches).execute()
            print(f"✅ {len(matches)} matches inserted successfully")
        except Exception as e:
            print(f"❌ Failed to insert matches: {e}")
            
if __name__ == '__main__':
    db_handler = DBHandler()
    if db_handler.get_client():
        print("Supabase client test successful.")
    else:
        print("Supabase client test failed.")
