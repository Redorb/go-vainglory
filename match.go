package vainglory

// GetStatsByName is a helper function to retrieve player MatchStats
// from a MatchResponse. A map is more performant than a slice in
// larger data sets so it is recommended to use this map instead of
// iterating through the MatchResponse data looking for players
func (mr *MatchResponse) GetStatsByName() (s map[string]*PlayerStats) {
	s = make(map[string]*PlayerStats)
	for _, p := range mr.Players {
		s[p.Attributes.Name] = &p.Attributes.Stats
	}
	return
}

// GetMatchID Pull match id from a given match response
func (mr *MatchResponse) GetMatchID() (id string) {
	return mr.Data.ID
}
