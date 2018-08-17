package vainglory

func (mr *MatchResponse) GetMatchID() (id string) {
	return mr.Data.ID
}
