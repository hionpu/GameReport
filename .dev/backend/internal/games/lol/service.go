package lol

import (
	"context"
	"fmt"
	"gameReport/internal/core"
)

type Service struct {
	apiClient *RiotClient
	analyzer  *PerformanceAnalyzer
	mapper    *CoreMapper
	region    string
}

func NewService(apiKey string) core.GameService {
	return NewServiceWithRegion(apiKey, "kr")
}

func NewServiceWithRegion(apiKey, region string) core.GameService {
	return &Service{
		apiClient: NewRiotClientWithRegion(apiKey, region),
		analyzer:  NewPerformanceAnalyzer(),
		mapper:    NewCoreMapper(),
		region:    region,
	}
}

func (s *Service) GetGame() core.GameType {
	return core.GameTypeLoL
}

func (s *Service) GetPlayer(ctx context.Context, identifier string) (*core.Player, error) {
	gameName, tagLine, err := parseIdentifier(identifier)
	if err != nil {
		return nil, fmt.Errorf("invalid player identifier: %w", err)
	}

	account, err := s.apiClient.GetAccountByRiotId(ctx, gameName, tagLine)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	summoner, err := s.apiClient.GetSummonerByPUUID(ctx, account.PUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get summoner: %w", err)
	}

	return &core.Player{
		ID:       account.PUUID,
		GameName: account.GameName,
		TagLine:  account.TagLine,
		Game:     core.GameTypeLoL,
		Level:    summoner.SummonerLevel,
		Region:   s.region,
	}, nil
}

func (s *Service) GetRecentMatches(ctx context.Context, playerID string, count int) ([]core.Match, error) {
	// Get match IDs
	matchIDs, err := s.apiClient.GetMatchHistory(ctx, playerID, count)
	if err != nil {
		return nil, err
	}

	var lolMatches []MatchDTO
	for _, matchID := range matchIDs {
		match, err := s.apiClient.GetMatchDetail(ctx, matchID)
		if err != nil {
			continue // Skip failed matches
		}
		lolMatches = append(lolMatches, *match)
	}

	// Convert LoL DTOs to core models
	return s.mapper.MatchesToCore(lolMatches, playerID), nil
}

func (s *Service) AnalyzePerformance(matches []core.Match) (*core.PerformanceAnalysis, error) {
	// Convert core matches back to Lol DTOs for detailed analysis
	lolMatches := s.mapper.CoreToMatches(matches)

	// Use the rich analyzer to full DTO data
	playerPUUID := matches[0].PlayerID
	return s.analyzer.AnalyzeMatches(lolMatches, playerPUUID)
}

func parseIdentifier(identifier string) (gameName, tagLine string, err error) {
	for i, char := range identifier {
		if char == '#' {
			if i == 0 || i == len(identifier)-1 {
				return "", "", fmt.Errorf("invalid identifier format")
			}
			return identifier[:i], identifier[i+1:], nil
		}
	}
	return "", "", fmt.Errorf("missing '#' separator")
}
