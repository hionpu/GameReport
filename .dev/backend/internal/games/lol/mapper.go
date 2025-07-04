package lol

import (
    "gameReport/internal/core"
    "time"
)

type CoreMapper struct{}

func NewCoreMapper() *CoreMapper {
    return &CoreMapper{}
}

func (m *CoreMapper) MatchesToCore(lolMatches []MatchDTO, playerPUUID string) []core.Match {
    var coreMatches []core.Match
    
    for _, match := range lolMatches {
        // Get game duration in seconds
        durationSeconds := time.Duration(match.Info.GameDuration) * time.Second
        
        // Determine match result for this player
        result := core.MatchResultLoss
        if player := match.GetPlayerByPUUID(playerPUUID); player != nil && player.Win {
            result = core.MatchResultWin
        }
        
        coreMatch := core.Match{
            ID:        match.Metadata.MatchID,
            PlayerID:  playerPUUID,
            Game:      core.GameTypeLoL,
            StartTime: time.Unix(match.Info.GameStartTimestamp/1000, 0), // FIXED: was PlayedAt
            Duration:  durationSeconds,
            GameMode:  match.Info.GameMode,
            Result:    result,
            Stats:     make(map[string]interface{}), // TODO: Extract relevant stats
            RawData:   match, // FIXED: was Data
        }
        coreMatches = append(coreMatches, coreMatch)
    }
    
    return coreMatches
}

func (m *CoreMapper) MatchesFromCore(coreMatches []core.Match) []MatchDTO {
    var lolMatches []MatchDTO
    
    for _, match := range coreMatches {
        if lolMatch, ok := match.Data.(MatchDTO); ok {
            lolMatches = append(lolMatches, lolMatch)
        }
    }
    
    return lolMatches
}