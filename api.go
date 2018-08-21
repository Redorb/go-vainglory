package vainglory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	shards  = "shards"  // shards path segment
	matches = "matches" // matches end point
	players = "players" // players end point
	status  = "status"  // status end point
	seasons = "seasons" // seasons end point

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

func (c *Client) newRequest(method, path string, body interface{}, options interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	url := u.String()
	if options != nil {
		v, err := query.Values(options)
		if err != nil {
			return nil, err
		}
		url = fmt.Sprintf("%s?%s", url, v.Encode())
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// GetStatus retrieves status data from the PUBG servers and passes back a StatusResponse.
func (c *Client) GetStatus() (*StatusResponse, error) {
	req, err := c.newRequest("GET", status, nil, nil)
	if err != nil {
		return nil, err
	}
	var response StatusResponse
	_, err = c.do(req, &response)
	return &response, err
}

// GetPlayer retrieves player data for a specified player id and passes back a PlayerResponse.
func (c *Client) GetPlayer(id string, shard string) (*PlayerResponse, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("%s/%s/%s/%s", shards, shard, players, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var response PlayerResponse
	_, err = c.do(req, &response)
	return &response, err
}

// GetMatch retrieves the match data for a specified match id and passes back a MatchResponse.
func (c *Client) GetMatch(id string, shard string) (*MatchResponse, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("%s/%s/%s/%s", shards, shard, matches, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var response MatchResponse
	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	for _, inc := range response.Included {
		var check map[string]string
		json.Unmarshal(inc, &check)
		switch check["type"] {
		case "player":
			var p MatchPlayer
			err = json.Unmarshal(inc, &p)
			response.Players = append(response.Players, p)
		case "participant":
			var p MatchParticipant
			err = json.Unmarshal(inc, &p)
			response.Participants = append(response.Participants, p)
		case "asset":
			var a MatchAsset
			err = json.Unmarshal(inc, &a)
			response.Assets = append(response.Assets, a)
		case "roster":
			var r MatchRoster
			err = json.Unmarshal(inc, &r)
			response.Rosters = append(response.Rosters, r)
		}

		if err != nil {
			return nil, err
		}
	}

	return &response, err
}

// GetMatches retrieves a list of match data and passes back a MatchesResponse.
func (c *Client) GetMatches(options GetMatchesRequestOptions, shard string) (*MatchesResponse, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("%s/%s/%s", shards, shard, matches), nil, options)
	if err != nil {
		return nil, err
	}

	var response MatchesResponse
	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	for _, inc := range response.Included {
		var check map[string]string
		json.Unmarshal(inc, &check)
		switch check["type"] {
		case "player":
			var p MatchPlayer
			err = json.Unmarshal(inc, &p)
			response.Players = append(response.Players, p)
		case "participant":
			var p MatchParticipant
			err = json.Unmarshal(inc, &p)
			response.Participants = append(response.Participants, p)
		case "asset":
			var a MatchAsset
			err = json.Unmarshal(inc, &a)
			response.Assets = append(response.Assets, a)
		case "roster":
			var r MatchRoster
			err = json.Unmarshal(inc, &r)
			response.Rosters = append(response.Rosters, r)
		}

		if err != nil {
			return nil, err
		}
	}

	return &response, err
}

// GetPlayers retrieves a list of player data and passes back a PlayersResponse.
func (c *Client) GetPlayers(options GetPlayersRequestOptions, shard string) (*PlayersResponse, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("%s/%s/%s", shards, shard, players), nil, options)
	if err != nil {
		return nil, err
	}

	var response PlayersResponse
	_, err = c.do(req, &response)
	return &response, err
}

// GetTelemetry retrieves the telemetry data at a specified url and passes back a TelemetryResponse.
func (c *Client) GetTelemetry(url string) (*TelemetryResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(resp.Body)
	return parseTelemetry(buffer.Bytes())
}

// ReadTelemetryFromFile parses json telemetry data from a given file
// and returns a TelemetryResponse struct. It is more performant to cache
// telemetry data for future use.
func ReadTelemetryFromFile(path string) (tr *TelemetryResponse, err error) {
	var b []byte
	b, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	return parseTelemetry(b)
}

// parseTelemetry reads the telemetry event type from the json
// and passes it to the unmarshaller
func parseTelemetry(b []byte) (*TelemetryResponse, error) {
	var v []json.RawMessage
	var tr TelemetryResponse
	json.Unmarshal(b, &v)
	for _, bts := range v {
		var eval map[string]interface{}
		err := json.Unmarshal(bts, &eval)
		if err != nil {
			return nil, err
		}
		tr.unmarshalEvent(bts, eval["_T"].(string))
	}
	return &tr, nil
}
