package vainglory

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

type PlayerStats struct {
	LossStreak   int     `json:"lossStreak"`
	WinStreak    int     `json:"winStreak"`
	LifetimeGold float64 `json:"lifetimeGold"`
}
