## League of Legends Match Data Structure

### **Main Match Data** (`/lol/match/v5/matches/{matchId}`)


**Path Parameters:**
- `puuid` (required, String): Player's unique identifier

**Query Parameters:**
- `startTime` (optional, long): Epoch timestamp in seconds. The matchlist started storing timestamps on June 16th, 2021. Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.
- `endTime` (optional, long): Epoch timestamp in seconds.
- `queue` (optional, int): Filter the list of match ids by a specific queue id. This filter is mutually inclusive of the type filter meaning any match ids returned must match both the queue and type filters.
- `type` (optional, string): Filter the list of match ids by the type of match. This filter is mutually inclusive of the queue filter meaning any match ids returned must match both the queue and type filters.
- `start` (optional, int): Defaults to 0. Start index.
- `count` (optional, int): Defaults to 20. Valid values: 0 to 100. Number of match ids to return.

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
- Team ID (100 for blue side, 200 for red side)
- Game duration, creation timestamps, queueID (420 = Ranked Solo)
- Map ID, game version, platform ID

The structs I've provided cover all the essential data you'll need for your gaming analytics app. You can access individual player data using `match_data["info"]["participants"][player_index]` and get specific stats like KDA, champion name, and win status.

For your daily report card feature, you'll primarily use the main match endpoint, while the timeline data would be useful for more advanced analytics later.

Here is an example of how the data could be represented in Elixir:

```elixir
defmodule GameReport.Lol.Match do
  defstruct metadata: %{}, info: %{}

  def new(data) do
    %__MODULE__{
      metadata: Map.get(data, "metadata"),
      info: GameReport.Lol.MatchInfo.new(Map.get(data, "info", %{}))
    }
  end
end

defmodule GameReport.Lol.MatchInfo do
  defstruct game_creation: nil,
            game_duration: nil,
            participants: [],
            teams: []
            # ... other fields

  def new(data) do
    %__MODULE__{
      game_creation: Map.get(data, "gameCreation"),
      game_duration: Map.get(data, "gameDuration"),
      participants: Enum.map(Map.get(data, "participants", []), &GameReport.Lol.Participant.new/1),
      teams: Enum.map(Map.get(data, "teams", []), &GameReport.Lol.Team.new/1)
      # ... other fields
    }
  end
end

defmodule GameReport.Lol.Participant do
  defstruct puuid: nil,
            champion_name: nil,
            kills: 0,
            deaths: 0,
            assists: 0,
            win: false
            # ... other fields

  def new(data) do
    %__MODULE__{
      puuid: Map.get(data, "puuid"),
      champion_name: Map.get(data, "championName"),
      kills: Map.get(data, "kills"),
      deaths: Map.get(data, "deaths"),
      assists: Map.get(data, "assists"),
      win: Map.get(data, "win")
      # ... other fields
    }
  end

  def kda(%__MODULE__{kills: k, deaths: d, assists: a}) do
    if d == 0, do: k + a, else: (k + a) / d
  end
end

defmodule GameReport.Lol.Team do
  defstruct team_id: nil,
            win: false,
            bans: []
            # ... other fields
  
  def new(data) do
    %__MODULE__{
      team_id: Map.get(data, "teamId"),
      win: Map.get(data, "win"),
      bans: Map.get(data, "bans", [])
      # ... other fields
    }
  end
end
```
