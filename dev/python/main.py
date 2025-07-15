from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
from datetime import datetime
import os

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

if __name__ == "__main__":
    port = int(os.getenv("PORT", "8001"))
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True,  # Enable auto-reload during development
        log_level="info"
    ) 
    
    