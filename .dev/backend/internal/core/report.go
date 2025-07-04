package core

import (
	"time"
)

// Universal report structure
type DailyReport struct {
	Player          Player              `json:"player"`
	Game            GameType            `json:"game"`
	Date            time.Time           `json:"date"`
	PerformanceData PerformanceAnalysis `json:"performance"`
	AIInsights      AIInsights          `json:"ai_insights"`
	CacheExpiry     CacheExpiry         `json:"cache_expiry"`
}

// Common performnace metrics across games
type PerformanceAnalysis struct {
	WinRate          float64                `json:"win_rate"`
	AvgKDA           float64                `json:"avg_kda,omitempty"`
	AvgScore         float64                `json:"avg_score,omitempty"`
	BestCharacter    string                 `json:"best_character,omitempty"`
	WeakestArea      string                 `json:"weakest_area"`
	TrendDirection   string                 `json:"trend_direction"`
	GamesAnalyzed    int                    `json:"games_analyzed"`
	GameSpecificData map[string]interface{} `json:"game_specific_data,omitempty"`
}

type AIInsights struct {
	PerformanceTrend string `json:"performance_trend"`
	Recommendation   string `json:"recommendation"`
	ImprovementFocus string `json:"improvement_focus"`
	Confidence       string `json:"confidence"`
}
