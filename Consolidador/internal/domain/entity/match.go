package entity

import "time"

type MatchResult struct{
    teamAScore int
    teamBScore int
}

func NewMatchResult(teamAScore, teamBScore int) *MatchResult{
    return &MatchResult{
        teamAScore: teamAScore,
        teamBScore: teamBScore,
    }
}


func (m *MatchResult) GetResult() string{
    return strconv.Itoa(m.teamAScore)+"-"+strconv.Itoa(m.teamBScore)
}

type Match struct {
    ID string
    TeamA *Team
    TeamB *Team
    TeamAID string
    TeamBID string
    Date time.time
    Status string
    Result MatchResult
}

func NewMatch(id string, teamA *Team, teamB *Team, date time.Time) *Match{
    return &Match{
        ID: id,
        TeamA: teamA,
        TeamB: teamB,
        TeamAID: teamA.ID,
        TeamBID: teamB.ID,
        Date:date,
    }
}