package vainglory

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	base    = "ttps://api.dc01.gamelockerapp.com" // base URL for making API calls
	shards  = "/shards/"                          // shards path segment
	matches = "/matches/"                         // matches end point
	players = "/players"                          //players end point
	status  = "/status"                           // status end point
	seasons = "/seasons"                          // seasons end point

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
	req, _ := http.NewRequest("GET", base+status, nil)
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

// GetPlayer retrieves data for the specified player and passes the PlayerResponseData into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
func (s *Session) GetPlayer(id, shard string, clbk func(PlayerResponseData, error)) (size int) {
	s.GetPlayers([]string{id}, shard, func(pr PlayerResponse, err error) {
		clbk(pr.Data, err)
	})
	return s.GetQueueSize()
}

// GetPlayers retrieves data for the passed names and passes the PlayerResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
func (s *Session) GetPlayers(ids []string, shard string, clbk func(PlayerResponse, error)) (size int) {
	return s.getPlayersByFilter(ids, shard, "playerIds", clbk)
}

func (s *Session) GetPlayersByName(names []string, shard string, clbk func(PlayerResponse, error)) (size int) {
	return s.getPlayersByFilter(names, shard, "playerNames", clbk)
}

func (s *Session) getPlayersByFilter(keys []string, shard, filter string, clbk func(PlayerResponse, error)) (size int) {
	query := strings.Replace(strings.Join(keys, ","), " ", "%20", -1)
	u, _ := url.ParseRequestURI(base + shards + shard + players + "?filter[" + filter + "]=" + query)
	req, _ := http.NewRequest("GET", u.String(), nil)
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
		clbk(pr, err)
	})
	return s.GetQueueSize()
}
