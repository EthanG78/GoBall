package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type PlayerStats struct {
	Name                 string `json:"name"`
	FieldGoalPercentage  string `json:"field_goal_percentage"`
	FreeThrowPercentage  string `json:"free_throw_percentage"`
	ThreePointPercentage string `json:"three_point_percentage"`
	PointsPerGame        string `json:"points_per_game"`
	ReboundsPerGame      string `json:"rebounds_per_game"`
	AssistsPerGame       string `json:"assists_per_game"`
	StealsPerGame        string `json:"steals_per_game"`
	BlocksPerGame        string `json:"blocks_per_game"`
	HeadShot             string
}

func FetchPlayerStats(first, last string) PlayerStats {
	firstname := url.QueryEscape(first)
	lastname := url.QueryEscape(last)

	url := fmt.Sprintf("https://nba-players.herokuapp.com/players-stats/%s/%s", lastname, firstname)
	headshot := fmt.Sprintf("https://nba-players.herokuapp.com/players/%s/%s", lastname, firstname)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Request Error: %v", err)
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Do Error: %v", err)
	}

	defer resp.Body.Close()

	var player PlayerStats

	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
		log.Printf("Error decoding: %v\n", err)
	}

	stats := PlayerStats{
		Name:                 player.Name,
		FieldGoalPercentage:  player.FieldGoalPercentage,
		FreeThrowPercentage:  player.FreeThrowPercentage,
		ThreePointPercentage: player.ThreePointPercentage,
		PointsPerGame:        player.PointsPerGame,
		ReboundsPerGame:      player.ReboundsPerGame,
		AssistsPerGame:       player.AssistsPerGame,
		StealsPerGame:        player.StealsPerGame,
		BlocksPerGame:        player.BlocksPerGame,
		HeadShot:             headshot,
	}

	return stats

}
