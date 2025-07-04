package core

import (
	"time"
)

// Player represnts a player across all games
type Player struct {
	ID       string   `json:"id"`
	GameName string   `json:"game_name"` // Display name
	TagLine  string   `json:"tag_line"`  // Tag (e.g., #NA1)
	Game     GameType `json:"game"`
	Rank     string   `json:"rank,omitempty"`
	Level    int      `json:"level,omitempty"`
	Region   string   `json:"region,omitempty"`
}

// Match represents a single game match across all games
type Match struct {
	ID         string                 `json:"id"`
	PlayerID   string                 `json:"player_id"`
	Game       GameType               `json:"game"`
	StartTime  time.Time              `json:"start_time"`
	Duration   time.Duration          `json:"duration"`
	GameMode   string                 `json:"game_mode"`
	Result     MatchResult            `json:"result"`
	Stats      map[string]interface{} `json:"stats"` // Game-specific stats
	RawData    interface{}            `json:"-"`     // Full game-specific data
}

type MatchResult string

const (
	MatchResultWin  MatchResult = "win"
	MatchResultLoss MatchResult = "loss"
	MatchResultDraw MatchResult = "draw"
)

// CacheExpiry defines cache durations for different data types
type CacheExpiry struct {
	PlayerData    time.Duration `json:"player_data"`
	MatchData     time.Duration `json:"match_data"`
	Analysis      time.Duration `json:"analysis"`
	AIInsights    time.Duration `json:"ai_insights"`
}

// DefaultCacheExpiry returns default cache expiration times
func DefaultCacheExpiry() CacheExpiry {
	return CacheExpiry{
		PlayerData: 15 * time.Minute,
		MatchData:  30 * time.Minute,
		Analysis:   2 * time.Hour,
		AIInsights: 4 * time.Hour,
	}
}