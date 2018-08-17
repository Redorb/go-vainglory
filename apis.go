package vainglory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	base    = "https://api.dc01.gamelockerapp.com" // base URL for making API calls
	shards  = "shards"                             // shards path segment
	matches = "matches"                            // matches end point
	players = "players"                            // players end point
	status  = "status"                             // status end point
	seasons = "seasons"                            // seasons end point

	// China Mobile China region
	China = "cn"
	// NorthAmerica Mobile North America region
	NorthAmerica = "na"
	// Europe Mobile Europe region
	Europe = "eu"
	// SouthAmerica Mobile South America region
	SouthAmerica = "sa"
	// EastAsia Mobile East Asia region
	EastAsia = "ea"
	// SoutheastAsia Mobile SEA region
	SoutheastAsia = "sg"
	// NATournament North America Tournament region
	NATournament = "tournament-na"
	// EUTournament Europe Tournament region
	EUTournament = "tournament-eu"
	// EATournament East Asia Tournament region
	EATournament = "tournament-ea"
	// SEATournament North America Tournament region
	SEATournament = "tournament-sg"
)

// GetQueueSize returns the current size of the poller queue.
// This is useful if implementing additional request limiting.
func (s *Session) GetQueueSize() (size int) {
	size = len(s.poller.queue)
	return
}

// GetStatus retrieves status data from the PUBG servers and passes the StatusResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
func (s *Session) GetStatus(clbk func(StatusResponse, error)) (size int) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", base, status), nil)
	s.poller.Request(req, func(res *http.Response, err error) {
		var sr StatusResponse
		if err != nil {
			clbk(sr, err)
			return
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &sr)
		clbk(sr, err)
	})
	return s.GetQueueSize()
}

// GetPlayer retrieves the player data for a specified player id and passes the PlayerResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
func (s *Session) GetPlayer(id string, shard string, clbk func(PlayerResponse, error)) (size int) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s/%s", base, shards, shard, players, id), nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	s.poller.Request(req, func(res *http.Response, err error) {
		var pr PlayerResponse
		if err != nil {
			clbk(pr, err)
			return
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &pr)
		if err != nil {
			clbk(pr, err)
		}
		// TODO clean up included
		clbk(pr, nil)
	})
	return s.GetQueueSize()
}

// GetMatch retrieves the match data for a specified match id and passes the MatchResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
func (s *Session) GetMatch(id string, shard string, clbk func(MatchResponse, error)) (size int) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s/%s", base, shards, shard, matches, id), nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	s.poller.Request(req, func(res *http.Response, err error) {
		var mr MatchResponse
		if err != nil {
			clbk(mr, err)
			return
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &mr)
		if err != nil {
			clbk(mr, err)
		}
		// TODO Clean up included struct
		clbk(mr, nil)
	})
	return s.GetQueueSize()
}

// GetMatches does stuff
func (s *Session) GetMatches(options GetMatchesRequestOptions, shard string, clbk func(MatchesResponse, error)) (size int) {
	v, _ := query.Values(options)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s?%s", base, shards, shard, matches, v.Encode()), nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	s.poller.Request(req, func(res *http.Response, err error) {
		var mr MatchesResponse
		if err != nil {
			clbk(mr, err)
			return
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &mr)
		if err != nil {
			clbk(mr, err)
		}
		// TODO Clean up included struct
		clbk(mr, nil)
	})
	return s.GetQueueSize()
}
