from abc import ABC, abstractmethod

class BaseAnalyzer(ABC):
    @abstractmethod
    def __init__(self, puuid: str):
        self.puuid = puuid
        pass

    @abstractmethod
    def get_user_report(self):
        pass

    # get raw data from the game's API
    @abstractmethod
    def fetch_data(self):
        pass

    # process the raw data into a structured format
    @abstractmethod
    def process_data(self, raw_data):
        pass
    
    # right report form structured data
    @abstractmethod
    def generate_insights(self, processed_data):
        pass