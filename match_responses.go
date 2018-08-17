package vainglory

import "time"

type MatchesResponse struct {
	Data     []MatchData             `json:"data"`
	Included []MatchResponseIncluded `json:"included"`
	Links    DataLinks               `json:"links"`
	Meta     struct {                //TODO
	} `json:"meta"`
}

type MatchResponse struct {
	Data     MatchData               `json:"data"`
	Included []MatchResponseIncluded `json:"included"`
	Links    DataLinks               `json:"links"`
	Meta     struct {                // TODO
	} `json:"meta"`
}

type MatchResponseIncluded struct {
	typeIDPair
	Attributes    MatchIncludedAttributes    `json:"attributes"`
	Relationships MatchIncludedRelationships `json:"relationships,omitempty"`
	Links         DataLinks                  `json:"links,omitempty"`
}

type MatchIncludedRelationships struct {
	Assets MatchDataRelationships `json:"assets"`
}

type MatchIncludedAttributes struct {
	Name         string                      `json:"name"`
	PatchVersion string                      `json:"patchVersion"`
	ShardID      string                      `json:"shardId"`
	Stats        MatchIncludedAttributeStats `json:"stats"`
	TitleID      string                      `json:"titleId"`
}

type MatchIncludedAttributeStats struct {
	EloEarnedSeason4 int         `json:"elo_earned_season_4"`
	EloEarnedSeason5 int         `json:"elo_earned_season_5"`
	EloEarnedSeason6 int         `json:"elo_earned_season_6"`
	EloEarnedSeason7 int         `json:"elo_earned_season_7"`
	EloEarnedSeason8 int         `json:"elo_earned_season_8"`
	EloEarnedSeason9 int         `json:"elo_earned_season_9"`
	GamesPlayed      GamesPlayed `json:"gamesPlayed"`
	GuildTag         string      `json:"guildTag"`
	KarmaLevel       int         `json:"karmaLevel"`
	Level            int         `json:"level"`
	LifetimeGold     int         `json:"lifetimeGold"`
	LossStreak       int         `json:"lossStreak"`
	Played           int         `json:"played"`
	PlayedAral       int         `json:"played_aral"`
	PlayedBlitz      int         `json:"played_blitz"`
	PlayedCasual     int         `json:"played_casual"`
	PlayedRanked     int         `json:"played_ranked"`
	RankPoints       RankPoints  `json:"rankPoints"`
	SkillTier        int         `json:"skillTier"`
	WinStreak        int         `json:"winStreak"`
	Wins             int         `json:"wins"`
	Xp               int         `json:"xp"`
}

type RankPoints struct {
	Blitz     float64 `json:"blitz"`
	Ranked    float64 `json:"ranked"`
	Ranked5V5 float64 `json:"ranked_5v5"`
}

type GamesPlayed struct {
	Aral        int `json:"aral"`
	Blitz       int `json:"blitz"`
	BlitzRounds int `json:"blitz_rounds"`
	Casual      int `json:"casual"`
	Casual5V5   int `json:"casual_5v5"`
	Ranked      int `json:"ranked"`
	Ranked5V5   int `json:"ranked_5v5"`
}

type MatchData struct {
	typeIDPair
	Attributes    MatchAttributes    `json:"attributes"`
	Relationships MatchRelationships `json:"relationships"`
	Links         DataLinks          `json:"links"`
}

type MatchAttributes struct {
	CreatedAt    time.Time            `json:"createdAt"`
	Duration     int                  `json:"duration"`
	GameMode     string               `json:"gameMode"`
	PatchVersion string               `json:"patchVersion"`
	ShardID      string               `json:"shardId"`
	Stats        MatchStatsAttributes `json:"stats"`
	Tags         interface{}          `json:"tags"` // TODO what does this look like
	TitleID      string               `json:"titleId"`
}

type MatchStatsAttributes struct {
	EndGameReason string `json:"endGameReason"`
	Queue         string `json:"queue"`
}

type MatchRelationships struct {
	Assets     MatchDataRelationships `json:"assets"`
	Rosters    MatchDataRelationships `json:"rosters"`
	Rounds     MatchDataRelationships `json:"rounds"`
	Spectators MatchDataRelationships `json:"spectators"`
}

type DataLinks struct {
	Schema string `json:"schema"`
	Self   string `json:"self"`
}

type MatchDataRelationships struct {
	Data []typeIDPair `json:"data"`
}
