""" Main entry point for the Python analysis server."""
from contextlib import asynccontextmanager
from dotenv import load_dotenv
from fastapi import FastAPI, BackgroundTasks, Request
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
from datetime import datetime
import os

from lol.analyzers import lol_analyzer
from shared.database.db_handler import DBHandler

@asynccontextmanager
async def lifespan(app: FastAPI):
    """Lifespan context manager for app initialization."""
    load_dotenv()
    
    # Initialize database handler
    db_url = os.getenv("SUPABASE_URL")
    db_key = os.getenv("SUPABASE_SERVICE_ROLE_KEY")
    if not db_url or not db_key:
        raise ValueError("SUPABASE_URL and SUPABASE_SERVICE_ROLE_KEY must be set in environment variables.")
    db_handler = DBHandler(db_url, db_key)

    # Initialize LoL analyzer
    riot_api_key = os.getenv("RIOT_API_KEY")
    if not riot_api_key:
        raise ValueError("RIOT_API_KEY is not set in environment variables.")
    
    app.state.lol_analyzer = lol_analyzer.LolAnalyzer(riot_api_key, db_handler)

    yield

app = FastAPI(
    lifespan=lifespan,
    title ="Daily Gaming Report Card - Analysis Server",
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
    
    
@app.post("/api/v1/pipelines/lol/run/{identifier}")
async def trigger_lol_pipeline(identifier: str, request: Request, background_tasks: BackgroundTasks):
    """Wrapper function for background task."""
    print(f"Starting LoL background analysis for {identifier}")
    analyzer_instance = request.app.state.lol_analyzer
    background_tasks.add_task(
        analyzer_instance.fetch_data,
        identifier
    )

    return {
        "status": "success",
        "message": f"Lol analysis pipeline for {identifier} has been queued.",
        "timestamp": datetime.now().isoformat()
    }

    

if __name__ == "__main__":
    port = int(os.getenv("PORT", "8001"))
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True,  # Enable auto-reload during development
        log_level="info"
    ) 
    
    