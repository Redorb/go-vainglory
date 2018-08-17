package vainglory

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

type Conf struct {
	Key       string `json:"key"`
	RateLimit int    `json:"rateLimit"`
}

const (
	expectedAPI = "v7.10.2"
)

var session *Session
var testPlayers = []string{"LuluXiu"}
var testPlayerIDs = []string{"52e62770-f295-11e6-912f-0223dbc2587a"}
var testMatchIDs = []string{"ae524352-8b85-11e8-8e53-0235a36ff992"}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestGetStatus(t *testing.T) {
	size := session.GetStatus(func(sr StatusResponse, err error) {
		if sr.Data.Attributes.Version != expectedAPI {
			t.Errorf("expected version %s but received %s", expectedAPI, sr.Data.Attributes.Version)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}

func TestGetPlayer(t *testing.T) {
	size := session.GetPlayer(testPlayerIDs[0], SoutheastAsia, func(pr PlayerResponse, err error) {
		if pr.Data.ID != testPlayerIDs[0] {
			t.Errorf("expected a player id of %s but received %s", testPlayerIDs[0], pr.Data.ID)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}

func TestGetMatch(t *testing.T) {
	size := session.GetMatch(testMatchIDs[0], SoutheastAsia, func(mr MatchResponse, err error) {
		if mr.GetMatchID() != testMatchIDs[0] {
			t.Errorf("expected a match id of %s but received %s", testMatchIDs[0], mr.GetMatchID())
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}

func TestGetMatches(t *testing.T) {
	currentTime := time.Now()
	options := GetMatchesRequestOptions{
		PlayerNamesFilter: testPlayers,
		EndCreationFilter: currentTime,
	}

	size := session.GetMatches(options, NorthAmerica, func(mr MatchesResponse, err error) {
		if err != nil {
			t.Errorf("Error getting matches from API: %s", err.Error())
		}
		if !strings.Contains(mr.Links.Self, testPlayers[0]) {
			t.Errorf(
				"Expected request url %s to contain player name %s, but not present",
				mr.Links.Self,
				testPlayers[0],
			)
		}

		if !strings.Contains(mr.Links.Self, testPlayers[0]) {
			t.Errorf(
				"Expected request url %s to contain time %s, but not present",
				mr.Links.Self,
				currentTime.String(),
			)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}

func setup() {
	b, _ := ioutil.ReadFile("conf.json")
	var conf Conf
	json.Unmarshal(b, &conf)
	var err error
	session, err = New(conf.Key, conf.RateLimit)
	if err != nil {
		panic("Error creating API Client")
	}
}
func teardown() {}
