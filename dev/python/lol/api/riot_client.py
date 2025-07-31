import requests
import time
from datetime import datetime
from typing import List, Dict, Optional

class RiotAPIClient:
    def __init__(self, api_key:str, region: str = "kr"):
        self.api_key = api_key
        self.region = region
        self.base_url = f"https://{region}.api.riotgames.com"
        self.headers = {
            "X-Riot-Token": api_key
        }
    def fetch_match_ids_by_puuid(self, puuid: str, start_time: int, count: int = 20) -> List[str]:
        url = f"{self.base_url}/lol/match/v5/matches/by-puuid/{puuid}/ids?startTime={start_time}&count={count}"
        params = {
            "start": 0,
            "count": count,
            "startTime": start_time
        }
        response = requests.get(url, headers=self.headers, params=params)
        response.raise_for_status()
        return response.json() 
    
    def fetch_match_details(self, match_id: str) -> Optional[Dict]:
        url = f"{self.base_url}/lol/match/v5/matches/{match_id}"
        response = requests.get(url, headers=self.headers)
        if response.status_code == 404:
            return None
        response.raise_for_status()
        return response.json()

        