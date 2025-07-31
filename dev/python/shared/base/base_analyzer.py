"""
Defines the abstract base class for game analyzers.

This module provides the `BaseAnalyzer` interface that all specific game
analyzers must implement.
"""

from abc import ABC, abstractmethod
from typing import Any


class BaseAnalyzer(ABC):
    """
    Abstract base class for all game analyzers.

    This class defines the standard interface for an analyzer, which includes
    fetching data, processing it, and generating insights. Each supported game
    should have a concrete implementation of this class.
    """

    @abstractmethod
    def __init__(self, **kwargs):
        """Initializes the analyzer with necessary configurations."""

    @abstractmethod
    def get_user_report(self, identifier: Any, **kwargs) -> Any:
        """Generates a comprehensive report for a specific user."""

    @abstractmethod
    def fetch_data(self, identifier: Any, **kwargs) -> Any:
        """Fetches raw data from the game's API."""

    @abstractmethod
    def process_data(self, raw_data: Any) -> Any:
        """Processes raw data into a structured format."""

    @abstractmethod
    def generate_insights(self, processed_data: Any) -> Any:
        """Generates actionable insights from processed data."""
