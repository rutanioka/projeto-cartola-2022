package db

import (
	"context"
	"database/sql"
)

const addPlayerToMyTeam = `-- name: AddPlayerToMyTeam :exec
INSERT INTO my_team_players (my_team_id, player_id) VALUES (?, ?)
`

type AddPlayerToMyTeamParams struct {
	MyTeamID string `json:"my_team_id"`
	PlayerID string `json:"player_id"`
}

func (q *Queries) AddPlayerToMyTeam(ctx context.Context, arg AddPlayerToMyTeamParams) error {
	_, err := q.db.ExecContext(ctx, addPlayerToMyTeam, arg.MyTeamID, arg.PlayerID)
	return err
}

const addScoreToTeam = `-- name: AddScoreToTeam :exec
UPDATE my_team SET score = score + ? WHERE id = ?
`

type AddScoreToTeamParams struct {
	Score float64 `json:"score"`
	ID    string  `json:"id"`
}

func (q *Queries) AddScoreToTeam(ctx context.Context, arg AddScoreToTeamParams) error {
	_, err := q.db.ExecContext(ctx, addScoreToTeam, arg.Score, arg.ID)
	return err
}

const createAction = `-- name: CreateAction :exec
INSERT INTO actions (id, match_id, team_id, player_id, action, score, minute) VALUES (?, ?, ?, ?, ?, ?,?)
`

type CreateActionParams struct {
	ID       string  `json:"id"`
	MatchID  string  `json:"match_id"`
	TeamID   string  `json:"team_id"`
	PlayerID string  `json:"player_id"`
	Action   string  `json:"action"`
	Score    float64 `json:"score"`
	Minute   int32   `json:"minute"`
}

func (q *Queries) CreateAction(ctx context.Context, arg CreateActionParams) error {
	_, err := q.db.ExecContext(ctx, createAction,
		arg.ID,
		arg.MatchID,
		arg.TeamID,
		arg.PlayerID,
		arg.Action,
		arg.Score,
		arg.Minute,
	)
	return err
}

const createMatch = `-- name: CreateMatch :exec
INSERT INTO matches (id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateMatchParams struct {
	ID        string         `json:"id"`
	MatchDate sql.NullTime   `json:"match_date"`
	TeamAID   sql.NullString `json:"team_a_id"`
	TeamAName sql.NullString `json:"team_a_name"`
	TeamBID   sql.NullString `json:"team_b_id"`
	TeamBName sql.NullString `json:"team_b_name"`
	Result    sql.NullString `json:"result"`
}

func (q *Queries) CreateMatch(ctx context.Context, arg CreateMatchParams) error {
	_, err := q.db.ExecContext(ctx, createMatch,
		arg.ID,
		arg.MatchDate,
		arg.TeamAID,
		arg.TeamAName,
		arg.TeamBID,
		arg.TeamBName,
		arg.Result,
	)
	return err
}

const createMyTeam = `-- name: CreateMyTeam :exec
INSERT INTO my_team (id, name, score) VALUES (?, ?, ?)
`

type CreateMyTeamParams struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

func (q *Queries) CreateMyTeam(ctx context.Context, arg CreateMyTeamParams) error {
	_, err := q.db.ExecContext(ctx, createMyTeam, arg.ID, arg.Name, arg.Score)
	return err
}

const createPlayer = `-- name: CreatePlayer :exec
INSERT INTO players (id, name, price) VALUES (?, ?, ?)
`

type CreatePlayerParams struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) error {
	_, err := q.db.ExecContext(ctx, createPlayer, arg.ID, arg.Name, arg.Price)
	return err
}

const deleteAllPlayersFromMyTeam = `-- name: DeleteAllPlayersFromMyTeam :exec
DELETE FROM my_team_players WHERE my_team_id = ?
`

func (q *Queries) DeleteAllPlayersFromMyTeam(ctx context.Context, myTeamID string) error {
	_, err := q.db.ExecContext(ctx, deleteAllPlayersFromMyTeam, myTeamID)
	return err
}

const findAllMatches = `-- name: FindAllMatches :many
SELECT id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result FROM matches
`

func (q *Queries) FindAllMatches(ctx context.Context) ([]Match, error) {
	rows, err := q.db.QueryContext(ctx, findAllMatches)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Match
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.MatchDate,
			&i.TeamAID,
			&i.TeamAName,
			&i.TeamBID,
			&i.TeamBName,
			&i.Result,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllPlayers = `-- name: FindAllPlayers :many
SELECT id, name, price FROM players
`

func (q *Queries) FindAllPlayers(ctx context.Context) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, findAllPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllPlayersByIDs = `-- name: FindAllPlayersByIDs :many
SELECT id, name, price FROM players WHERE id IN (?)
`

func (q *Queries) FindAllPlayersByIDs(ctx context.Context, id string) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, findAllPlayersByIDs, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findMatchById = `-- name: FindMatchById :one
SELECT id, match_date, team_a_id, team_a_name, team_b_id, team_b_name, result FROM matches WHERE id = ?
`

func (q *Queries) FindMatchById(ctx context.Context, id string) (Match, error) {
	row := q.db.QueryRowContext(ctx, findMatchById, id)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.MatchDate,
		&i.TeamAID,
		&i.TeamAName,
		&i.TeamBID,
		&i.TeamBName,
		&i.Result,
	)
	return i, err
}

const findMyTeamById = `-- name: FindMyTeamById :one
SELECT id, name, score FROM my_team WHERE id = ?
`

func (q *Queries) FindMyTeamById(ctx context.Context, id string) (MyTeam, error) {
	row := q.db.QueryRowContext(ctx, findMyTeamById, id)
	var i MyTeam
	err := row.Scan(&i.ID, &i.Name, &i.Score)
	return i, err
}

const findMyTeamByIdForUpdate = `-- name: FindMyTeamByIdForUpdate :one
SELECT id, name, score FROM my_team WHERE id = ? FOR UPDATE
`

func (q *Queries) FindMyTeamByIdForUpdate(ctx context.Context, id string) (MyTeam, error) {
	row := q.db.QueryRowContext(ctx, findMyTeamByIdForUpdate, id)
	var i MyTeam
	err := row.Scan(&i.ID, &i.Name, &i.Score)
	return i, err
}

const findPlayerById = `-- name: FindPlayerById :one
SELECT id, name, price FROM players WHERE id = ?
`

func (q *Queries) FindPlayerById(ctx context.Context, id string) (Player, error) {
	row := q.db.QueryRowContext(ctx, findPlayerById, id)
	var i Player
	err := row.Scan(&i.ID, &i.Name, &i.Price)
	return i, err
}

const findPlayerByIdForUpdate = `-- name: FindPlayerByIdForUpdate :one
SELECT id, name, price FROM players WHERE id = ? FOR UPDATE
`

func (q *Queries) FindPlayerByIdForUpdate(ctx context.Context, id string) (Player, error) {
	row := q.db.QueryRowContext(ctx, findPlayerByIdForUpdate, id)
	var i Player
	err := row.Scan(&i.ID, &i.Name, &i.Price)
	return i, err
}

const findTeamById = `-- name: FindTeamById :one
SELECT id, name FROM teams WHERE id = ?
`

func (q *Queries) FindTeamById(ctx context.Context, id string) (Team, error) {
	row := q.db.QueryRowContext(ctx, findTeamById, id)
	var i Team
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getMatchActions = `-- name: GetMatchActions :many
SELECT id, match_id, team_id, player_id, action, minute, score FROM actions WHERE match_id = ?
`

func (q *Queries) GetMatchActions(ctx context.Context, matchID string) ([]Action, error) {
	rows, err := q.db.QueryContext(ctx, getMatchActions, matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.MatchID,
			&i.TeamID,
			&i.PlayerID,
			&i.Action,
			&i.Minute,
			&i.Score,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMatchActionsForUpdate = `-- name: GetMatchActionsForUpdate :many
SELECT id, match_id, team_id, player_id, action, minute, score FROM actions WHERE match_id = ? FOR UPDATE
`

func (q *Queries) GetMatchActionsForUpdate(ctx context.Context, matchID string) ([]Action, error) {
	rows, err := q.db.QueryContext(ctx, getMatchActionsForUpdate, matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.MatchID,
			&i.TeamID,
			&i.PlayerID,
			&i.Action,
			&i.Minute,
			&i.Score,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMyTeamBalance = `-- name: GetMyTeamBalance :one
SELECT score as balance FROM my_team WHERE id = ?
`

func (q *Queries) GetMyTeamBalance(ctx context.Context, id string) (float64, error) {
	row := q.db.QueryRowContext(ctx, getMyTeamBalance, id)
	var balance float64
	err := row.Scan(&balance)
	return balance, err
}

const getPlayersByMyTeamID = `-- name: GetPlayersByMyTeamID :many
SELECT p.id, p.name, p.price FROM players p INNER JOIN my_team_players mtp ON p.id = mtp.player_id WHERE mtp.my_team_id = ?
`

func (q *Queries) GetPlayersByMyTeamID(ctx context.Context, myTeamID string) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getPlayersByMyTeamID, myTeamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeActionFromMatch = `-- name: RemoveActionFromMatch :exec
DELETE FROM actions WHERE match_id = ?
`

func (q *Queries) RemoveActionFromMatch(ctx context.Context, matchID string) error {
	_, err := q.db.ExecContext(ctx, removeActionFromMatch, matchID)
	return err
}

const updateMatch = `-- name: UpdateMatch :exec
UPDATE matches SET match_date = ?, team_a_id = ?, team_a_name = ?, team_b_id = ?, team_b_name = ?, result = ? WHERE id = ?
`

type UpdateMatchParams struct {
	MatchDate sql.NullTime   `json:"match_date"`
	TeamAID   sql.NullString `json:"team_a_id"`
	TeamAName sql.NullString `json:"team_a_name"`
	TeamBID   sql.NullString `json:"team_b_id"`
	TeamBName sql.NullString `json:"team_b_name"`
	Result    sql.NullString `json:"result"`
	ID        string         `json:"id"`
}

func (q *Queries) UpdateMatch(ctx context.Context, arg UpdateMatchParams) error {
	_, err := q.db.ExecContext(ctx, updateMatch,
		arg.MatchDate,
		arg.TeamAID,
		arg.TeamAName,
		arg.TeamBID,
		arg.TeamBName,
		arg.Result,
		arg.ID,
	)
	return err
}

const updateMyTeamScore = `-- name: UpdateMyTeamScore :exec
UPDATE my_team SET score = ? WHERE id = ?
`

type UpdateMyTeamScoreParams struct {
	Score float64 `json:"score"`
	ID    string  `json:"id"`
}

func (q *Queries) UpdateMyTeamScore(ctx context.Context, arg UpdateMyTeamScoreParams) error {
	_, err := q.db.ExecContext(ctx, updateMyTeamScore, arg.Score, arg.ID)
	return err
}

const updateMyTeamsScore = `-- name: UpdateMyTeamsScore :exec
UPDATE my_team SET score = ? WHERE id IN (?)
`

type UpdateMyTeamsScoreParams struct {
	Score float64 `json:"score"`
	ID    string  `json:"id"`
}

func (q *Queries) UpdateMyTeamsScore(ctx context.Context, arg UpdateMyTeamsScoreParams) error {
	_, err := q.db.ExecContext(ctx, updateMyTeamsScore, arg.Score, arg.ID)
	return err
}

const updatePlayer = `-- name: UpdatePlayer :exec
UPDATE players SET name = ?, price = ? WHERE id = ?
`

type UpdatePlayerParams struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	ID    string  `json:"id"`
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) error {
	_, err := q.db.ExecContext(ctx, updatePlayer, arg.Name, arg.Price, arg.ID)
	return err
}