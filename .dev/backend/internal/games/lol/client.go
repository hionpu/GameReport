package lol

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RiotClient struct {
	apiKey          string
	httpClient      *http.Client
	region          string
	regionalRouting string
}

// Region configuration
var regionConfig = map[string]string{
	"kr":   "asia",     // Korea
	"jp1":  "asia",     // Japan
	"na1":  "americas", // North America
	"br1":  "americas", // Brazil
	"la1":  "americas", // Latin America North
	"la2":  "americas", // Latin America South
	"oc1":  "americas", // Oceania
	"euw1": "europe",   // Europe West
	"eun1": "europe",   // Europe Nordic & East
	"tr1":  "europe",   // Turkey
	"ru":   "europe",   // Russia
}

func NewRiotClient(apiKey string) *RiotClient {
	return NewRiotClientWithRegion(apiKey, "kr")
}

func NewRiotClientWithRegion(apiKey, region string) *RiotClient {
	regionalRouting, exists := regionConfig[region]
	if !exists {
		regionalRouting = "asia"
	}

	return &RiotClient{
		apiKey:          apiKey,
		region:          region,
		regionalRouting: regionalRouting,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *RiotClient) getRegionalURL() string {
	return fmt.Sprintf("https://%s.api.riotgames.com", c.regionalRouting)
}

func (c *RiotClient) getPlatformURL() string {
	return fmt.Sprintf("https://%s.api.riotgames.com", c.region)
}

func (c *RiotClient) GetAccountByRiotID(ctx context.Context, gameName, tagLine string) (*AccountDTO, error) {
	url := fmt.Sprintf("%s/riot/account/v1/accounts/by-riot-id/%s/%s",
		c.getPlatformURL(), gameName, tagLine)

	var account AccountDTO
	err := c.makeRequest(ctx, url, &account)
	return &account, err
}

func (c *RiotClient) GetSummonerByPUUID(ctx context.Context, puuid string) (*SummonerDTO, error) {
	url := fmt.Sprintf("%s/riot/summoner/v4/summoners/by-puuid/%s",
		c.getPlatformURL(), puuid)

	var summoner SummonerDTO
	err := c.makeRequest(ctx, url, &summoner)
	return &summoner, err
}

func (c *RiotClient) GetMatchHistory(ctx context.Context, puuid string, count int) ([]string, error) {
	url := fmt.Sprintf("%s/lol/match/v5/matches/by-puuid/%s/ids?count=%d", 
		c.getRegionalURL(), puuid, count)
	
	var matchIDs []string
	err := c.makeRequest(ctx, url, &matchIDs)
	return matchIDs, err
}

func (c *RiotClient) GetMatchDetail(ctx context.Context, matchID string) (*MatchDTO, error) {
	url := fmt.Sprintf("%s/lol/match/v5/matches/%s", 
		c.getRegionalURL(), matchID)
	
	var match MatchDTO
	err := c.makeRequest(ctx, url, &match)
	return &match, err
}

func (c *RiotClient) makeRequest(ctx context.Context, url string, result interface{}) error {
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create requestuest: %w", err)
	}

	request.Header.Set("X-Riot-Token", c.apiKey)
	request.Header.Set("Accept", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("API error %d: %s", response.StatusCode, string(body))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}
