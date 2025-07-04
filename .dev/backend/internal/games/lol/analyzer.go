package lol

import (
	"fmt"
	"gameReport/internal/core"
)

type PerformanceAnalyzer struct{}

func NewPerformanceAnalyzer() *PerformanceAnalyzer {
	return &PerformanceAnalyzer{}
}

func (a *PerformanceAnalyzer) AnalyzeMatches(matches []MatchDTO, playerPUUID string) (*core.PerformanceAnalysis, error) {
	if len(matches) == 0 {
		return nil, fmt.Errorf("no matches to analyze")
	}

	var playerPerformances []PlayerMatchPerformance

	// Extract detailed performance for each match
	for _, match := range matches {
		player := match.GetPlayerByPUUID(playerPUUID)
		if player == nil {
			continue
		}

		performance := PlayerMatchPerformance{
			Match: match,
			Player: *player,
			GameLength: match.GetGameDurationMinutes(),
			IsRanked: match.IsRanked(),
			KDA: player.GetKDA(),
			CSPerMin: player.GetCSPerMinute(match.Info.GameDuration),
			DamageShare: a.calculateDamageShare(*player, match),
			VisionScore: player.VisionScore,
			Champion: player.ChampionName,
			Role: player.TeamPosition,
			Win: player.Win,
		}

		playerPerformances = append(playerPerformances, performance)
	}

	return &core.PerformanceAnalysis{
		WinRate: a.calculateWinRate(playerPerformances),
		AvgKDA: a.calculateAvgKDA(playerPerformances),
		BestCharacter: a.findBestChampion(playerPerformances),
		WeakestArea: a.identifyWeakestArea(playerPerformances),
		TrendDirection: a.analyzeTrend(playerPerformances),
		GamesAnalyzed: len(playerPerformances),
		GameSpecificData: map[string]interface{}{
            "avg_cs_per_min":        a.calculateAvgCSPerMin(playerPerformances),
            "avg_damage_share":      a.calculateAvgDamageShare(playerPerformances),
            "avg_vision_score":      a.calculateAvgVisionScore(playerPerformances),
            "preferred_role":        a.findPreferredRole(playerPerformances),
            "champion_diversity":    a.calculateChampionDiversity(playerPerformances),
            "ranked_games_ratio":    a.calculateRankedRatio(playerPerformances),
            "avg_game_length":       a.calculateAvgGameLength(playerPerformances),
            "multikill_frequency":   a.calculateMultikillFrequency(playerPerformances),
            "objective_participation": a.calculateObjectiveParticipation(playerPerformances),
            "early_game_performance": a.analyzeEarlyGamePerformance(playerPerformances),
        },
	}, nil
}

type PlayerMatchPerformance struct {
    Match       MatchDTO
    Player      ParticipantDTO
    GameLength  float64
    IsRanked    bool
    KDA         float64
    CSPerMin    float64
    DamageShare float64
    VisionScore int
    Champion    string
    Role        string
    Win         bool
}

func (a *PerformanceAnalyzer) calculateWinRate(performances []PlayerMatchPerformance) float64 {
	wins := 0
	for _, p := range performances {
		if p.Win {
			wins++
		}
	}
	return float64(wins) / float64(len(performances)) * 100
}

func (a *PerformanceAnalyzer) calculateAvgKDA(performances []PlayerMatchPerformance) float64 {
	total := 0.0
	for _, p := range performances {
		total += p.KDA
	}
	
	return total/ float64(len(performances))
}

type ChampionStats struct {
    Games    int
    Wins     int
    TotalKDA float64
}

func (a *PerformanceAnalyzer) findBestChampion(performances []PlayerMatchPerformance) string {
	champStats := make(map[string]ChampionStats)
	
	for _, p := range performances {
		stats := champStats[p.Champion]
		stats.Games++
		stats.TotalKDA += p.KDA
		if p.Win {
			stats.Wins++
		}
		champStats[p.Champion] = stats
	}

	var bestChamp string
	var bestScore float64

	for champ, stats := range champStats {
		if stats.Games < 2 {
			continue
		}

		winRate := float64(stats.Wins) / float64(stats.Games)
		avgKDA := stats.TotalKDA / float64(stats.Games)
		score := winRate*0.6 + (avgKDA/5.0)*0.4

		if score > bestScore {
			bestScore = score
			bestChamp = champ
		}
	}

	return bestChamp
}

func (a *PerformanceAnalyzer) identifyWeakestArea(performances []PlayerMatchPerformance) string  {
	avgCS := a.calculateAvgCSPerMin(performances)
	avgVision := a.calculateAvgVisionScore(performances)
	avgDamageShare := a.calculateAvgDamageShare(performances)

	issues := []string{}

	if avgCS < 6.0 {
		issues = append(issues, "CS per minute")
	}
	if avgVision < 25 {
		issues = append(issues, "Vision score")
	}
	if avgDamageShare < 20.0 {
		issues = append(issues, "Damage output")
	}

	if len(issues) == 0  {
		return "macro play and positioning"
	}

	return issues[0]
}

func (a *PerformanceAnalyzer) calculateDamageShare(player ParticipantDTO, match MatchDTO) float64 {
    teamTotalDamage := 0
    
    for _, participant := range match.Info.Participants {
        if participant.TeamID == player.TeamID {
            teamTotalDamage += participant.TotalDamageDealtToChampions
        }
    }
    
    return player.GetDamageShare(teamTotalDamage)
}

// Additional helper methods...
func (a *PerformanceAnalyzer) calculateAvgCSPerMin(performances []PlayerMatchPerformance) float64 {
    total := 0.0
    for _, p := range performances {
        total += p.CSPerMin
    }
    return total / float64(len(performances))
}

func (a *PerformanceAnalyzer) calculateAvgVisionScore(performances []PlayerMatchPerformance) float64 {
    total := 0.0
    for _, p := range performances {
        total += float64(p.VisionScore)
    }
    return total / float64(len(performances))
}

func (a *PerformanceAnalyzer) calculateAvgDamageShare(performances []PlayerMatchPerformance) float64 {
	total := 0.0
	for _, p := range performances {
		total += p.DamageShare
	}
	return total / float64(len(performances))
}

func (a *PerformanceAnalyzer) findPreferredRole(performances []PlayerMatchPerformance) string {
	roleCounts := make(map[string]int)
	for _, p := range performances {
		roleCounts[p.Role]++
	}
	
	var preferredRole string
	var maxCount int
	for role, count := range roleCounts {
		if count > maxCount {
			maxCount = count
			preferredRole = role
		}
	}
	
	return preferredRole
}

func (a *PerformanceAnalyzer) calculateChampionDiversity(performances []PlayerMatchPerformance) int {
	champSet := make(map[string]bool)
	for _, p := range performances {
		champSet[p.Champion] = true
	}
	return len(champSet)
}

func (a *PerformanceAnalyzer) calculateRankedRatio(performances []PlayerMatchPerformance) float64 {
	rankedCount := 0
	for _, p := range performances {
		if p.IsRanked {
			rankedCount++
		}
	}
	return float64(rankedCount) / float64(len(performances))
}

func (a *PerformanceAnalyzer) calculateAvgGameLength(performances []PlayerMatchPerformance) float64 {
	total := 0.0
	for _, p := range performances {
		total += p.GameLength
	}
	return total / float64(len(performances))
}

func (a *PerformanceAnalyzer) calculateMultikillFrequency(performances []PlayerMatchPerformance) float64 {
	multikillCount := 0
	for _, p := range performances {
		if p.Player.DoubleKills > 0 || p.Player.TripleKills > 0 || 
		   p.Player.QuadraKills > 0 || p.Player.PentaKills > 0 {
			multikillCount++
		}
	}
	return float64(multikillCount) / float64(len(performances))
}

func (a *PerformanceAnalyzer) calculateObjectiveParticipation(performances []PlayerMatchPerformance) float64 {
	total := 0.0
	for _, p := range performances {
		participation := float64(p.Player.DragonKills + p.Player.BaronKills + p.Player.TurretKills)
		total += participation
	}
	return total / float64(len(performances))
}

func (a *PerformanceAnalyzer) analyzeEarlyGamePerformance(performances []PlayerMatchPerformance) string {
	avgCSAt15 := 0.0
	for _, p := range performances {
		// Rough estimate: CS at 15 minutes
		csAt15 := p.CSPerMin * 15
		avgCSAt15 += csAt15
	}
	avgCSAt15 /= float64(len(performances))
	
	if avgCSAt15 > 120 {
		return "Strong"
	} else if avgCSAt15 > 80 {
		return "Average"
	} else {
		return "Weak"
	}
}

func (a *PerformanceAnalyzer) analyzeTrend(performances []PlayerMatchPerformance) string {
	if len(performances) < 3 {
		return "Insufficient data"
	}
	
	// Simple trend analysis: compare last 3 games vs first 3 games
	recentGames := performances[len(performances)-3:]
	earlyGames := performances[:3]
	
	recentWins := 0
	earlyWins := 0
	
	for _, p := range recentGames {
		if p.Win {
			recentWins++
		}
	}
	
	for _, p := range earlyGames {
		if p.Win {
			earlyWins++
		}
	}
	
	if recentWins > earlyWins {
		return "Improving"
	} else if recentWins < earlyWins {
		return "Declining"
	} else {
		return "Stable"
	}
}