package games

import (
    "fmt"
    "gameReport/internal/core"
    "gameReport/internal/games/lol"
)

type registry struct {
    services map[core.GameType]core.GameService
}

func NewRegistry(riotAPIKey string) core.GameRegistry {
    r := &registry{
        services: make(map[core.GameType]core.GameService),
    }
    
    // Register LoL service
    r.Register(core.GameTypeLoL, lol.NewService(riotAPIKey))
    
    // Future: Add other games
    // r.Register(core.GameTypeValorant, valorant.NewService(riotAPIKey))
    // r.Register(core.GameTypeCS2, cs2.NewService(steamAPIKey))
    
    return r
}

func (r *registry) Register(gameType core.GameType, service core.GameService) {
    r.services[gameType] = service
}

func (r *registry) GetService(gameType core.GameType) (core.GameService, error) {
    service, exists := r.services[gameType]
    if !exists {
        return nil, fmt.Errorf("game %s not supported", gameType)
    }
    return service, nil
}

func (r *registry) GetSupportedGames() []core.GameType {
    games := make([]core.GameType, 0, len(r.services))
    for gameType := range r.services {
        games = append(games, gameType)
    }
    return games
}