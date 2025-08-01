""" Main entry point for the Python analysis server."""
from dotenv import load_dotenv
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
from datetime import datetime
import os

from lol.analyzers import lol_analyzer
from shared.database.db_handler import DBHandler


# Initialize FastAPI app
app = FastAPI(
    title="Daily Gaming Report Card - Analysis Server",
    description="Python analysis server for gaming performance insights",
    version="1.0.0"
)

# Add CORS middleware for Go server communication
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # In production, restrict to Go server URL
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/health")
async def health_check():
    """Health check endpoint to verify server status"""
    return {
        "status": "healthy",
        "timestamp": datetime.now().isoformat(),
        "server": "python-analysis",
        "version": "1.0.0"
    }
    
db_handler = DBHandler()
load_dotenv()

riot_api_key = os.getenv("RIOT_API_KEY")
if not riot_api_key:
    raise ValueError("RIOT_API_KEY is not set in environment variables.")
else:
    lol_instance = lol_analyzer.LolAnalyzer(riot_api_key, db_handler)
    
    
@app.post("/api/v1/pipelines/lol/run")
async def trigger_lol_pipeline(identifier: str, background_tasks: BackgroundTasks):
    """Wrapper function for background task."""
    print(f"Starting LoL background analysis for {identifier}")
    lol_instance.fetch_data(identifier)

    background_tasks

if __name__ == "__main__":
    port = int(os.getenv("PORT", "8001"))
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True,  # Enable auto-reload during development
        log_level="info"
    ) 
    
    