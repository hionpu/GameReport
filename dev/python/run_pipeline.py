from dotenv import load_dotenv
import os
from database.db_handler import DBHandler
from analyzers.lol_analyzer import LolAnalyzer

def run_pipeline():
    load_dotenv()
    db_handler = DBHandler()

    try:
        api_key = os.getenv("RIOT_API_KEY")
        analyzer = LolAnalyzer(api_key, db_handler)
        test_faker_puuid = "I781COdv2FFhLh4DVF1biyd0xW7WAw2yv6mgiD3Wr6elzNAiMvUEhGd7ps4Z0z-7em75Kghrh_tTwA"
        result = analyzer.run_analysis(test_faker_puuid)
        print(f"Analysis result: {result}")
    except Exception as e:
        print(f"Error: {e}")
        
if __name__ == "__main__":
    run_pipeline()

