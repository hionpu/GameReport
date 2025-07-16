import requests
import time
from typing import List, Dict, Optional

class RiotAPIClient:
    def __init__(self, api_key:str, region: str = "kr"):
        self.api_key = api_key
        self.region = region
        self.base_url = f"https://{region}.api.riotgames.com"
        self.headers = {
            "X-Riot-Token": api_key
        }

        