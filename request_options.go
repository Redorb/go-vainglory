package vainglory

import "time"

// GetPlayersRequestOptions Filter options for the get players endpoint
type GetPlayersRequestOptions struct {
	// PlayerNamesFilter Filters by player name. Usage: filter[playerNames]=player1,player2,…
	PlayerNamesFilter []string `url:"filter[playerNames],comma,omitempty"`
	// PlayerIDsFilter Filters by player Id. Usage:filter[playerIds]=playerId,playerId,…
	PlayerIDsFilter []string `url:"filter[playerIds],comma,omitempty"`
}

// GetMatchesRequestOptions Various paging, sorting, and filter options for the get matches endpoint
type GetMatchesRequestOptions struct {
	// PageOffset Allows paging over results
	PageOffset int `url:"page[offset],omitempty"`
	// PageLimit The default (and maximum) is 5. Values less than 5 and greater than 1 are supported.
	PageLimit int `url:"page[limit],omitempty"`
	// Sort By default, Matches are sorted by creation time ascending.
	Sort string `url:"sort,omitempty"`
	// StartCreationFilter Must occur before end time. Format is iso8601 Usage: filter[createdAt-start]=2017-01-01T08:25:30Z
	StartCreationFilter time.Time `url:"filter[createdAt-start],omitempty"`
	// EndCreationFilter Queries search the last 3 hrs. Format is iso8601 i.e.filter[createdAt-end]=2017-01-01T13:25:30Z
	EndCreationFilter time.Time `url:"filter[createdAt-end],omitempty"`
	// PlayerNamesFilter Filters by player name. Usage: filter[playerNames]=player1,player2,…
	PlayerNamesFilter []string `url:"filter[playerNames],comma,omitempty"`
	// PlayerIDsFilter Filters by player Id. Usage:filter[playerIds]=playerId,playerId,…
	PlayerIDsFilter []string `url:"filter[playerIds],comma,omitempty"`
	// TeamNamesFilter Filters by team names. Team names are the same as the in game team tags. Usage: filter[teamNames]=TSM,team2,…
	TeamNamesFilter []string `url:"filter[teamNames],comma,omitempty"`
	// GameModeFilter Filter by gameMode Usage: filter[gameMode]=casual,ranked,…
	GameModeFilter []string `url:"filter[gameMode],comma,omitempty"`
	// PatchVersionFilter Filter by Vainglory patch version. Usage: filter[patchVersion]=2.10,2.11,…
	PatchVersionFilter []string `url:"filter[patchVersion],comma,omitempty"`
}
