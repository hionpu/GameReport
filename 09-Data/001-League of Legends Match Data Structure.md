## League of Legends Match Data Structure

### **Main Match Data** (`/lol/match/v5/matches/{matchId}`)

**Key Structure:**


```json
{
  "metadata": {
    "dataVersion": "2",
    "matchId": "NA1_4483419198", 
    "participants": ["puuid1", "puuid2", ...] // 10 PUUIDs
  },
  "info": {
    "gameCreation": 1640995200000,
    "gameDuration": 1456,
    "participants": [...], // 10 participant objects
    "teams": [...], // 2 team objects
    // ... other game info
  }
}
```

### **Timeline Data** (`/lol/match/v5/matches/{matchId}/timeline`)

Contains frame-by-frame data every minute with events like kills, item purchases, level ups, etc.

### **Key Data Points Available:**

**Per Player:**

- Basic stats: kills, deaths, assists, championName, win status
- Economy: goldEarned, goldSpent, items (0-6)
- Combat: damage dealt/taken, healing, CC time
- Farm: totalMinionsKilled, neutralMinionsKilled
- Vision: visionScore, wardsPlaced, wardsKilled
- Objectives: dragonKills, baronKills, turretKills
- Runes: Complete perk information

**Per Team:**

- Win/loss, bans, objective totals
- Team ID (100 for blue side, 200 for red side) [riotclient package - github.com/torlenor/alolstats/riotclient - Go Packages](https://pkg.go.dev/github.com/torlenor/alolstats/riotclient)

**Game Info:**

- Game duration, creation timestamps, queueID (420 = Ranked Solo) [Crawling matches using the Riot Games API](https://hextechdocs.dev/crawling-matches-using-the-riot-games-api/)
- Map ID, game version, platform ID

The structs I've provided cover all the essential data you'll need for your gaming analytics app. You can access individual player data using `match_data['info']['participants'][player_index]` and get specific stats like KDA, champion name, and win status.

For your daily report card feature, you'll primarily use the main match endpoint, while the timeline data would be useful for more advanced analytics later.

So the go file would like:

```go
package main

import "time"

// MatchDTO represents the complete match data from Riot API Match-v5
type MatchDTO struct {
	Metadata MatchMetadataDTO `json:"metadata"`
	Info     MatchInfoDTO     `json:"info"`
}

// MatchMetadataDTO contains metadata about the match
type MatchMetadataDTO struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"` // Array of participant PUUIDs
}

// MatchInfoDTO contains the main match information
type MatchInfoDTO struct {
	GameCreation       int64               `json:"gameCreation"`       // Unix timestamp in milliseconds
	GameDuration       int                 `json:"gameDuration"`       // Game length in seconds
	GameEndTimestamp   int64               `json:"gameEndTimestamp"`   // Unix timestamp in milliseconds
	GameID             int64               `json:"gameId"`
	GameMode           string              `json:"gameMode"`           // e.g., "CLASSIC", "ARAM"
	GameName           string              `json:"gameName"`
	GameStartTimestamp int64               `json:"gameStartTimestamp"` // Unix timestamp in milliseconds
	GameType           string              `json:"gameType"`           // e.g., "MATCHED_GAME"
	GameVersion        string              `json:"gameVersion"`        // e.g., "12.1.123.456"
	MapID              int                 `json:"mapId"`              // 11 = Summoner's Rift, 12 = ARAM
	Participants       []ParticipantDTO    `json:"participants"`       // Array of 10 participants
	PlatformID         string              `json:"platformId"`         // e.g., "NA1", "EUW1"
	QueueID            int                 `json:"queueId"`            // 420 = Ranked Solo, 400 = Normal Draft
	Teams              []TeamDTO           `json:"teams"`              // Array of 2 teams
	TournamentCode     string              `json:"tournamentCode,omitempty"`
}

// ParticipantDTO contains individual player performance data
type ParticipantDTO struct {
	// Basic Info
	ParticipantID int    `json:"participantId"` // 1-10
	TeamID        int    `json:"teamId"`        // 100 or 200
	PUUID         string `json:"puuid"`
	SummonerName  string `json:"summonerName"`
	SummonerID    string `json:"summonerId"`
	
	// Champion & Game Setup
	ChampionID             int    `json:"championId"`
	ChampionName           string `json:"championName"`
	ChampionTransform      int    `json:"championTransform"` // For Kayn, Kha'Zix etc.
	ChampionLevel          int    `json:"champLevel"`
	Role                   string `json:"role"`         // "DUO_CARRY", "DUO_SUPPORT", etc.
	Lane                   string `json:"lane"`         // "BOTTOM", "JUNGLE", etc.
	TeamPosition           string `json:"teamPosition"` // "BOTTOM", "JUNGLE", "MIDDLE", "TOP", "UTILITY"
	
	// Summoner Spells
	Summoner1ID    int `json:"summoner1Id"`
	Summoner2ID    int `json:"summoner2Id"`
	Summoner1Casts int `json:"summoner1Casts"`
	Summoner2Casts int `json:"summoner2Casts"`
	
	// Combat Stats
	Kills                int  `json:"kills"`
	Deaths               int  `json:"deaths"`
	Assists              int  `json:"assists"`
	LargestKillingSpree  int  `json:"largestKillingSpree"`
	LargestMultiKill     int  `json:"largestMultiKill"`
	DoubleKills          int  `json:"doubleKills"`
	TripleKills          int  `json:"tripleKills"`
	QuadraKills          int  `json:"quadraKills"`
	PentaKills           int  `json:"pentaKills"`
	UnrealKills          int  `json:"unrealKills"`
	FirstBloodKill       bool `json:"firstBloodKill"`
	FirstBloodAssist     bool `json:"firstBloodAssist"`
	
	// Damage Stats
	TotalDamageDealt             int `json:"totalDamageDealt"`
	TotalDamageDealtToChampions  int `json:"totalDamageDealtToChampions"`
	PhysicalDamageDealt          int `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int `json:"physicalDamageDealtToChampions"`
	MagicDamageDealt             int `json:"magicDamageDealt"`
	MagicDamageDealtToChampions  int `json:"magicDamageDealtToChampions"`
	TrueDamageDealt              int `json:"trueDamageDealt"`
	TrueDamageDealtToChampions   int `json:"trueDamageDealtToChampions"`
	LargestCriticalStrike        int `json:"largestCriticalStrike"`
	
	// Damage Taken
	TotalDamageTaken         int `json:"totalDamageTaken"`
	PhysicalDamageTaken      int `json:"physicalDamageTaken"`
	MagicDamageTaken         int `json:"magicDamageTaken"`
	TrueDamageTaken          int `json:"trueDamageTaken"`
	DamageSelfMitigated      int `json:"damageSelfMitigated"`
	
	// Healing & Shielding
	TotalHeal               int `json:"totalHeal"`
	TotalHealsOnTeammates   int `json:"totalHealsOnTeammates"`
	TotalDamageShieldedOnTeammates int `json:"totalDamageShieldedOnTeammates"`
	
	// Economy
	GoldEarned  int `json:"goldEarned"`
	GoldSpent   int `json:"goldSpent"`
	
	// Items
	Item0 int `json:"item0"`
	Item1 int `json:"item1"`
	Item2 int `json:"item2"`
	Item3 int `json:"item3"`
	Item4 int `json:"item4"`
	Item5 int `json:"item5"`
	Item6 int `json:"item6"` // Trinket
	
	// CS & Monsters
	TotalMinionsKilled        int `json:"totalMinionsKilled"`
	NeutralMinionsKilled      int `json:"neutralMinionsKilled"`
	TotalTimeCCDealt          int `json:"totalTimeCCDealt"`
	
	// Vision
	VisionScore                   int  `json:"visionScore"`
	WardsPlaced                   int  `json:"wardsPlaced"`
	WardsKilled                   int  `json:"wardsKilled"`
	DetectorWardsPlaced           int  `json:"detectorWardsPlaced"`
	
	// Objectives
	DragonKills   int  `json:"dragonKills"`
	BaronKills    int  `json:"baronKills"`
	TurretKills   int  `json:"turretKills"`
	InhibitorKills int `json:"inhibitorKills"`
	FirstTowerKill bool `json:"firstTowerKill"`
	FirstTowerAssist bool `json:"firstTowerAssist"`
	
	// Game Result
	Win       bool `json:"win"`
	GameEnded bool `json:"gameEndedInEarlySurrender"`
	TeamEarlySurrendered bool `json:"teamEarlySurrendered"`
	
	// Perks (Runes)
	Perks PerksDTO `json:"perks"`
	
	// Time-based Stats
	TimePlayed                    int `json:"timePlayed"`
	TotalTimeSpentDead            int `json:"totalTimeSpentDead"`
	LongestTimeSpentLiving        int `json:"longestTimeSpentLiving"`
	
	// Additional Stats
	NexusKills               int `json:"nexusKills"`
	ObjectivesStolen         int `json:"objectivesStolen"`
	ObjectivesStolenAssists  int `json:"objectivesStolenAssists"`
}

// PerksDTO contains rune information
type PerksDTO struct {
	StatPerks PerkStatsDTO   `json:"statPerks"`
	Styles    []PerkStyleDTO `json:"styles"`
}

type PerkStatsDTO struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type PerkStyleDTO struct {
	Description string           `json:"description"`
	Selections  []PerkSelectionDTO `json:"selections"`
	Style       int              `json:"style"`
}

type PerkSelectionDTO struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

// TeamDTO contains team-level information
type TeamDTO struct {
	TeamID    int          `json:"teamId"`    // 100 (Blue) or 200 (Red)
	Win       bool         `json:"win"`
	Bans      []BanDTO     `json:"bans"`
	Objectives ObjectivesDTO `json:"objectives"`
}

type BanDTO struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type ObjectivesDTO struct {
	Baron      ObjectiveDTO `json:"baron"`
	Champion   ObjectiveDTO `json:"champion"`
	Dragon     ObjectiveDTO `json:"dragon"`
	Inhibitor  ObjectiveDTO `json:"inhibitor"`
	RiftHerald ObjectiveDTO `json:"riftHerald"`
	Tower      ObjectiveDTO `json:"tower"`
}

type ObjectiveDTO struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

// MatchTimelineDTO represents the timeline data (separate endpoint)
type MatchTimelineDTO struct {
	Metadata MatchMetadataDTO `json:"metadata"`
	Info     TimelineInfoDTO  `json:"info"`
}

type TimelineInfoDTO struct {
	FrameInterval int              `json:"frameInterval"` // Usually 60000 (1 minute)
	Frames        []MatchFrameDTO  `json:"frames"`
	GameID        int64            `json:"gameId"`
	Participants  []ParticipantInfoDTO `json:"participants"`
}

type MatchFrameDTO struct {
	Events           []MatchEventDTO                    `json:"events"`
	ParticipantFrames map[string]MatchParticipantFrameDTO `json:"participantFrames"` // Key is participant ID as string
	Timestamp        int                                `json:"timestamp"`
}

type MatchEventDTO struct {
	Type               string             `json:"type"`               // "CHAMPION_KILL", "ITEM_PURCHASED", etc.
	Timestamp          int                `json:"timestamp"`
	ParticipantID      int                `json:"participantId,omitempty"`
	ItemID             int                `json:"itemId,omitempty"`
	SkillSlot          int                `json:"skillSlot,omitempty"`
	LevelUpType        string             `json:"levelUpType,omitempty"`
	CreatorID          int                `json:"creatorId,omitempty"`
	VictimID           int                `json:"victimId,omitempty"`
	KillerID           int                `json:"killerId,omitempty"`
	AssistingParticipantIds []int         `json:"assistingParticipantIds,omitempty"`
	Position           PositionDTO        `json:"position,omitempty"`
	AfterID            int                `json:"afterId,omitempty"`
	BeforeID           int                `json:"beforeId,omitempty"`
	GoldGain           int                `json:"goldGain,omitempty"`
	TeamID             int                `json:"teamId,omitempty"`
	MonsterType        string             `json:"monsterType,omitempty"`
	MonsterSubType     string             `json:"monsterSubType,omitempty"`
	BuildingType       string             `json:"buildingType,omitempty"`
	LaneType           string             `json:"laneType,omitempty"`
	TowerType          string             `json:"towerType,omitempty"`
	WardType           string             `json:"wardType,omitempty"`
}

type MatchParticipantFrameDTO struct {
	ChampionStats        ChampionStatsDTO `json:"championStats"`
	CurrentGold          int              `json:"currentGold"`
	DamageStats          DamageStatsDTO   `json:"damageStats"`
	GoldPerSecond        int              `json:"goldPerSecond"`
	JungleMinionsKilled  int              `json:"jungleMinionsKilled"`
	Level                int              `json:"level"`
	MinionsKilled        int              `json:"minionsKilled"`
	ParticipantID        int              `json:"participantId"`
	Position             PositionDTO      `json:"position"`
	TimeEnemySpentControlled int          `json:"timeEnemySpentControlled"`
	TotalGold            int              `json:"totalGold"`
	XP                   int              `json:"xp"`
}

type ChampionStatsDTO struct {
	AbilityHaste       int `json:"abilityHaste"`
	AbilityPower       int `json:"abilityPower"`
	Armor              int `json:"armor"`
	ArmorPen           int `json:"armorPen"`
	ArmorPenPercent    int `json:"armorPenPercent"`
	AttackDamage       int `json:"attackDamage"`
	AttackSpeed        int `json:"attackSpeed"`
	BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
	BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
	CCReduction        int `json:"ccReduction"`
	CooldownReduction  int `json:"cooldownReduction"`
	Health             int `json:"health"`
	HealthMax          int `json:"healthMax"`
	HealthRegen        int `json:"healthRegen"`
	Lifesteal          int `json:"lifesteal"`
	MagicPen           int `json:"magicPen"`
	MagicPenPercent    int `json:"magicPenPercent"`
	MagicResist        int `json:"magicResist"`
	MovementSpeed      int `json:"movementSpeed"`
	Omnivamp           int `json:"omnivamp"`
	PhysicalVamp       int `json:"physicalVamp"`
	Power              int `json:"power"`
	PowerMax           int `json:"powerMax"`
	PowerRegen         int `json:"powerRegen"`
	SpellVamp          int `json:"spellVamp"`
}

type DamageStatsDTO struct {
	MagicDamageDone             int `json:"magicDamageDone"`
	MagicDamageDoneToChampions  int `json:"magicDamageDoneToChampions"`
	MagicDamageTaken            int `json:"magicDamageTaken"`
	PhysicalDamageDone          int `json:"physicalDamageDone"`
	PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
	PhysicalDamageTaken         int `json:"physicalDamageTaken"`
	TotalDamageDone             int `json:"totalDamageDone"`
	TotalDamageDoneToChampions  int `json:"totalDamageDoneToChampions"`
	TotalDamageTaken            int `json:"totalDamageTaken"`
	TrueDamageDone              int `json:"trueDamageDone"`
	TrueDamageDoneToChampions   int `json:"trueDamageDoneToChampions"`
	TrueDamageTaken             int `json:"trueDamageTaken"`
}

type PositionDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ParticipantInfoDTO struct {
	ParticipantID int    `json:"participantId"`
	PUUID         string `json:"puuid"`
}

// Helper functions for your application
func (m *MatchDTO) GetPlayerByPUUID(puuid string) *ParticipantDTO {
	for i, participant := range m.Info.Participants {
		if participant.PUUID == puuid {
			return &m.Info.Participants[i]
		}
	}
	return nil
}

func (m *MatchDTO) GetTeamByID(teamID int) *TeamDTO {
	for i, team := range m.Info.Teams {
		if team.TeamID == teamID {
			return &m.Info.Teams[i]
		}
	}
	return nil
}

func (m *MatchDTO) IsRanked() bool {
	return m.Info.QueueID == 420 || m.Info.QueueID == 440 // Solo/Duo or Flex
}

func (m *MatchDTO) GetGameDurationMinutes() float64 {
	return float64(m.Info.GameDuration) / 60.0
}

func (m *MatchDTO) GetGameStartTime() time.Time {
	return time.Unix(m.Info.GameStartTimestamp/1000, 0)
}

func (p *ParticipantDTO) GetKDA() float64 {
	if p.Deaths == 0 {
		return float64(p.Kills + p.Assists)
	}
	return float64(p.Kills+p.Assists) / float64(p.Deaths)
}

func (p *ParticipantDTO) GetCSPerMinute(gameDurationSeconds int) float64 {
	if gameDurationSeconds == 0 {
		return 0
	}
	totalCS := p.TotalMinionsKilled + p.NeutralMinionsKilled
	return float64(totalCS) / (float64(gameDurationSeconds) / 60.0)
}
```
