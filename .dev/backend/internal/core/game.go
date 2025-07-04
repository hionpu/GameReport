package core

import (
	"context"
	"time"
)


type GameService interface {
	GetGame() GameType
	GetPlayer(ctx context.Context, identifier string) (*Player, error)
	GetRecentMatches(ctx context.Context, playerID string, count int) ([]Match, error)
	AnalyzePerformance(matches []Match) (*PerformanceAnalysis, error)
}

type GameType string

const (
	GameTypeLoL      GameType = "lol"
	GameTypeValorant GameType = "valorant"
	GameTypeCS2      GameType = "cs2"
	GameTypeDota2    GameType = "dota2"
)

type GameRegistry interface { 
	Register(gameType GameType, service GameService)
	GetService(gameType GameType) (GameService, error)
	GetSupportedGames() []GameType
}

