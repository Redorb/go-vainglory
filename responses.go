package vainglory

import (
	"encoding/json"
	"time"
)

// StatusResponse is the response payload for the status end point
type StatusResponse struct {
	Data StatusResponseData `json:"data"`
}

// StatusResponseData contains all of the data returned in the StatusResponse
type StatusResponseData struct {
	typeIDPair
	Attributes StatusAttributes `json:"attributes"`
}

// StatusAttributes contains all of the attributes returned in the StatusResponse
type StatusAttributes struct {
	Released string `json:"releasedAt"`
	Version  string `json:"version"`
}

type MatchesResponse struct {
	Data         []MatchData       `json:"data"`
	Included     []json.RawMessage `json:"included"`
	Participants []MatchParticipant
	Rosters      []MatchRoster
	Assets       []MatchAsset
	Players      []MatchPlayer
	Links        struct {
		Schema string `json:"schema"`
		Self   string `json:"self"`
	} `json:"links"`
	Meta struct{} `json:"meta"`
}

type MatchResponse struct {
	Data         MatchData         `json:"data"`
	Included     []json.RawMessage `json:"included"`
	Participants []MatchParticipant
	Rosters      []MatchRoster
	Assets       []MatchAsset
	Players      []MatchPlayer
	Links        struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct{} `json:"meta"`
}

type MatchParticipant struct {
	typeIDPair
	Attributes struct {
		Actor   string           `json:"actor"`
		ShardID string           `json:"shardId"`
		Stats   ParticipantStats `json:"stats"`
	} `json:"attributes"`
	Relationships struct {
		Player MatchDataRelationships `json:"player"`
	} `json:"relationships"`
}

type ParticipantStats struct {
	Assists              int            `json:"assists"`
	CrystalMineCaptures  int            `json:"crystalMineCaptures"`
	Deaths               int            `json:"deaths"`
	Farm                 int            `json:"farm"`
	FirstAfkTime         int            `json:"firstAfkTime"`
	Gold                 float64        `json:"gold"`
	GoldMineCaptures     int            `json:"goldMineCaptures"`
	ItemGrants           map[string]int `json:"itemGrants"`
	ItemSells            map[string]int `json:"itemSells"`
	ItemUses             map[string]int `json:"itemUses"`
	Items                []string       `json:"items"`
	JungleKills          int            `json:"jungleKills"`
	KarmaLevel           int            `json:"karmaLevel"`
	Kills                int            `json:"kills"`
	KrakenCaptures       int            `json:"krakenCaptures"`
	Level                int            `json:"level"`
	MinionKills          int            `json:"minionKills"`
	NonJungleMinionKills int            `json:"nonJungleMinionKills"`
	SkillTier            int            `json:"skillTier"`
	SkinKey              string         `json:"skinKey"`
	TurretCaptures       int            `json:"turretCaptures"`
	WentAfk              bool           `json:"wentAfk"`
	Winner               bool           `json:"winner"`
}

type MatchRoster struct {
	typeIDPair
	Attributes struct {
		ShardID string      `json:"shardId"`
		Stats   RosterStats `json:"stats"`
		Won     string      `json:"won"`
	} `json:"attributes"`
	Relationships struct {
		Participants MatchDataRelationships `json:"participants"`
	} `json:"relationships"`
}

type RosterStats struct {
	AcesEarned       int    `json:"acesEarned"`
	Gold             int    `json:"gold"`
	HeroKills        int    `json:"heroKills"`
	KrakenCaptures   int    `json:"krakenCaptures"`
	Side             string `json:"side"`
	TurretKills      int    `json:"turretKills"`
	TurretsRemaining int    `json:"turretsRemaining"`
}

type MatchAsset struct {
	typeIDPair
	Attributes struct {
		URL         string    `json:"URL"`
		CreatedAt   time.Time `json:"createdAt"`
		Description string    `json:"description"`
		Name        string    `json:"name"`
	} `json:"attributes"`
}

type MatchPlayer struct {
	typeIDPair
	Attributes struct {
		Name         string      `json:"name"`
		PatchVersion string      `json:"patchVersion"`
		ShardID      string      `json:"shardId"`
		Stats        PlayerStats `json:"stats"`
		TitleID      string      `json:"titleId"`
	} `json:"attributes"`
	Relationships struct {
		Assets MatchDataRelationships `json:"assets"`
	} `json:"relationships"`
	Links struct {
		Schema string `json:"schema"`
		Self   string `json:"self"`
	} `json:"links"`
}

type PlayerStats struct {
	EloEarnedSeason4 int `json:"elo_earned_season_4"`
	EloEarnedSeason5 int `json:"elo_earned_season_5"`
	EloEarnedSeason6 int `json:"elo_earned_season_6"`
	EloEarnedSeason7 int `json:"elo_earned_season_7"`
	EloEarnedSeason8 int `json:"elo_earned_season_8"`
	EloEarnedSeason9 int `json:"elo_earned_season_9"`
	GamesPlayed      struct {
		Aral        int `json:"aral"`
		Blitz       int `json:"blitz"`
		BlitzRounds int `json:"blitz_rounds"`
		Casual      int `json:"casual"`
		Casual5V5   int `json:"casual_5v5"`
		Ranked      int `json:"ranked"`
		Ranked5V5   int `json:"ranked_5v5"`
	} `json:"gamesPlayed"`
	GuildTag     string `json:"guildTag"`
	KarmaLevel   int    `json:"karmaLevel"`
	Level        int    `json:"level"`
	LifetimeGold int    `json:"lifetimeGold"`
	LossStreak   int    `json:"lossStreak"`
	Played       int    `json:"played"`
	PlayedAral   int    `json:"played_aral"`
	PlayedBlitz  int    `json:"played_blitz"`
	PlayedCasual int    `json:"played_casual"`
	PlayedRanked int    `json:"played_ranked"`
	RankPoints   struct {
		Blitz     float64 `json:"blitz"`
		Ranked    float64 `json:"ranked"`
		Ranked5V5 float64 `json:"ranked_5v5"`
	} `json:"rankPoints"`
	SkillTier int `json:"skillTier"`
	WinStreak int `json:"winStreak"`
	Wins      int `json:"wins"`
	Xp        int `json:"xp"`
}

type MatchData struct {
	typeIDPair
	Attributes    MatchAttributes    `json:"attributes"`
	Relationships MatchRelationships `json:"relationships"`
	Links         struct {
		Schema string `json:"schema"`
		Self   string `json:"self"`
	} `json:"links"`
}

type MatchAttributes struct {
	CreatedAt    time.Time  `json:"createdAt"`
	Duration     int        `json:"duration"`
	GameMode     string     `json:"gameMode"`
	PatchVersion string     `json:"patchVersion"`
	ShardID      string     `json:"shardId"`
	Stats        MatchStats `json:"stats"`
	TitleID      string     `json:"titleId"`
}

type MatchStats struct {
	EndGameReason string `json:"endGameReason"`
	Queue         string `json:"queue"`
}

type MatchRelationships struct {
	Assets     MatchDataRelationships `json:"assets"`
	Rosters    MatchDataRelationships `json:"rosters"`
	Rounds     MatchDataRelationships `json:"rounds"`
	Spectators MatchDataRelationships `json:"spectators"`
}

type MatchDataRelationships struct {
	Data []typeIDPair `json:"data"`
}

type PlayerResponse struct {
	Data PlayerResponseData `json:"data"`
}

type PlayerResponseData struct {
	typeIDPair
	Attributes PlayerAttributes `json:"attributes"`
}

type PlayerAttributes struct {
	Name         string      `json:"name"`
	PatchVersion string      `json:"patchVersion"`
	ShardID      string      `json:"shardId"`
	Stats        PlayerStats `json:"stats"`
	TitleID      string      `json:"titleID"`
}

// TelemetryResponse is the response payload for the telemetry end point
type TelemetryResponse struct {
	Events                  []TelemetryEvent
	HeroBanEvents           []*HeroBanEvent
	HeroSelectEvents        []*HeroSelectEvent
	HeroSkinSelectEvents    []*HeroSkinSelectEvent
	PlayerFirstSpawnEvents  []*PlayerFirstSpawnEvent
	LevelUpEvents           []*LevelUpEvent
	BuyItemEvents           []*BuyItemEvent
	LearnAbilityEvents      []*LearnAbilityEvent
	UseAbilityEvents        []*UseAbilityEvent
	UseItemAbilityEvents    []*UseItemAbilityEvent
	DealDamageEvents        []*DealDamageEvent
	HealTargetEvents        []*HealTargetEvent
	EarnXPEvents            []*EarnXPEvent
	KillActorEvents         []*KillActorEvent
	VampirismEvents         []*VampirismEvent
	GoldFromTowerKillEvents []*GoldFromTowerKillEvent
	SellItemEvents          []*SellItemEvent
	NPCkillNPCEvents        []*NPCkillNPCEvent
}

// typeIDPair is a common pattern used throughout all responses
type typeIDPair struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
