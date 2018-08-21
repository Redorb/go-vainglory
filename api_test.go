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
	Key string `json:"key"`
}

const (
	expectedAPI = "v7.10.2"
)

var client *Client
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
	sr, err := client.GetStatus()
	if err != nil {
		t.Errorf("Error getting status from API: %s", err.Error())
	}

	if sr.Data.Attributes.Version != expectedAPI {
		t.Errorf("expected version %s but received %s", expectedAPI, sr.Data.Attributes.Version)
	}
}

func TestGetPlayer(t *testing.T) {
	pr, err := client.GetPlayer(testPlayerIDs[0], SoutheastAsia)
	if err != nil {
		t.Errorf("Error getting player from API: %s", err.Error())
	}

	if pr.Data.ID != testPlayerIDs[0] {
		t.Errorf("expected a player id of %s but received %s", testPlayerIDs[0], pr.Data.ID)
	}
}

func TestGetPlayers(t *testing.T) {
	options := GetPlayersRequestOptions{
		PlayerNamesFilter: testPlayers,
	}

	pr, err := client.GetPlayers(options, NorthAmerica)

	if err != nil {
		t.Errorf("Error getting players from API: %s", err.Error())
	}
	if !strings.Contains(pr.Links.Self, testPlayers[0]) {
		t.Errorf(
			"Expected request url %s to contain player name %s, but not present",
			pr.Links.Self,
			testPlayers[0],
		)
	}
}

func TestGetMatch(t *testing.T) {
	mr, err := client.GetMatch(testMatchIDs[0], SoutheastAsia)
	if err != nil {
		t.Errorf("Error getting match from API: %s", err.Error())
	}

	if mr.GetMatchID() != testMatchIDs[0] {
		t.Errorf("expected a match id of %s but received %s", testMatchIDs[0], mr.GetMatchID())
	}
}

func TestGetMatchNonexistant(t *testing.T) {
	_, err := client.GetMatch("1213141", SoutheastAsia)
	if err == nil {
		t.Error("Expected error from api")
	}
}

func TestGetMatches(t *testing.T) {
	currentTime := time.Now()
	options := GetMatchesRequestOptions{
		PlayerNamesFilter: testPlayers,
		EndCreationFilter: currentTime,
	}

	mr, err := client.GetMatches(options, NorthAmerica)

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
}

func setup() {
	b, _ := ioutil.ReadFile("conf.json")
	var conf Conf
	json.Unmarshal(b, &conf)
	var err error
	client, err = New(conf.Key, nil)
	if err != nil {
		panic("Error instantiating client")
	}
}
func teardown() {}
