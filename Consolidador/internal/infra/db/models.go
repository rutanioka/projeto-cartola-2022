
package db

import (
	"database/sql"
)

type Action struct {
	ID       string  `json:"id"`
	MatchID  string  `json:"match_id"`
	TeamID   string  `json:"team_id"`
	PlayerID string  `json:"player_id"`
	Action   string  `json:"action"`
	Minute   int32   `json:"minute"`
	Score    float64 `json:"score"`
}

type Match struct {
	ID        string         `json:"id"`
	MatchDate sql.NullTime   `json:"match_date"`
	TeamAID   sql.NullString `json:"team_a_id"`
	TeamAName sql.NullString `json:"team_a_name"`
	TeamBID   sql.NullString `json:"team_b_id"`
	TeamBName sql.NullString `json:"team_b_name"`
	Result    sql.NullString `json:"result"`
}

type MyTeam struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type MyTeamPlayer struct {
	MyTeamID string `json:"my_team_id"`
	PlayerID string `json:"player_id"`
}

type Player struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeamPlayer struct {
	TeamID   string `json:"team_id"`
	PlayerID string `json:"player_id"`
}