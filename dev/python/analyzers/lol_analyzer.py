from .base import BaseAnalyzer

class LolAnalyzer(BaseAnalyzer):
    def __init__(self, api_key):
        self.api_key = api_key

    def get_user_report(self):
        pass
    
    def fetch_data(self, puuid):
        """
        Fetches match data for a given user.
        For MVP, this returns mock data.
        """
        print(f"Fetching data for user: {puuid}")
        # Mock structure based on Riot API Match-v5
        return {
            "info": {
                "gameDuration": 1820,
                "participants": [
                    {
                        "puuid": f"puuid_{i}",
                        "championName": f"Champion_{i}",
                        "teamPosition": "MIDDLE",
                        "win": i % 2 == 0,
                        "kills": i + 2,
                        "deaths": i,
                        "assists": 10 - i,
                        "totalMinionsKilled": 150 + i * 5,
                        "neutralMinionsKilled": 10 + i,
                        "totalDamageDealtToChampions": 20000 + i * 500,
                        "goldEarned": 12000 + i * 200,
                        "visionScore": 25 + i,
                        "controlWardsPlaced": 3 + i % 2,
                    } for i in range(1, 11) # Simulating 10 players
                ]
            }
        }

    def process_data(self, raw_data):
        """
        Processes raw match data to calculate and structure key performance indicators.
        This transforms the API response into a format that is ready for database insertion.
        """
        processed_matches = []

        # Extract game duration and convert to minits for calculations
        # This is shared values for all participants in the match
        duration_minutes = raw_data["info"]["gameDuration"] / 60.0
        if duration_minutes == 0:
            duration_minutes = 1.0
            
        for participant in raw_data['info']['participants']:
            # --- KDA ---
            if participant['deaths'] == 0:
                kda = participant['kills'] + participant['assists']
            else: 
                kda = (participant['kills'] + participant['assists']) / participant['deaths']
            
            # --- Per-Minute Stats ---
            # should be normalized by game duration
            total_cs = participant['totalMinionsKilled'] + participant['neutralMinionsKilled']
            cs_per_minute = total_cs / duration_minutes
            damage_per_minute = participant['totalDamageDealtToChampions'] / duration_minutes
            gold_per_minute = participant['goldEarned'] / duration_minutes
            vision_score_per_minute = participant['visionScore'] / duration_minutes

            structured_match = {
                'puuid': participant['puuid'],
                'game_duration': raw_data['info']['gameDuration'],
                'champion_name': participant['championName'],
                'team_position': participant['teamPosition'],
                'win': participant['win'],
                'kills': participant['kills'],
                'deaths': participant['deaths'],
                'assists': participant['assists'],
                'kda': round(kda, 2),
                'cs_per_min': round(cs_per_minute, 2),
                'damage_per_min': round(damage_per_minute, 2),
                'gold_per_min': round(gold_per_minute, 2),
                'vision_score_per_min': round(vision_score_per_minute, 2),
                'control_wards_placed': participant['controlWardsPlaced'],
            }

            processed_matches.append(structured_match)

        return processed_matches

    def save_data(self, processed_data):
        """
        Saves the processed data to the database
        TODO: Implement this
        """            
        print("Saving data to database...")
        for match in processed_data:
            print(f"Saving match: {match['puuid']}")
            print("Data saved")
            
    def generate_insights(self, processed_data):
        pass

