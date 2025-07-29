from abc import ABC, abstractmethod
from typing import Any

class BaseAnalyzer(ABC):
    @abstractmethod
    def __init__(self, **kwargs):
        pass

    @abstractmethod
    def get_user_report(self, **kwargs) -> Any:
        pass

    # get raw data from the game's API
    @abstractmethod
    def fetch_data(self, **kwargs) -> Any:
        pass

    # process the raw data into a structured format
    @abstractmethod
    def process_data(self, raw_data: Any) -> Any:
        pass

    # right report form structured data
    @abstractmethod
    def generate_insights(self, processed_data: Any) -> Any:
        pass
    